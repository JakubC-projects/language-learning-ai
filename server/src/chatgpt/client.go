package chatgpt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	apiKey string
}

func (c Client) CreateChatCompletion(request ChatCompletionRequest) (ChatCompletionResponse, error) {
	var result ChatCompletionResponse
	request.Stream = false

	resp, err := c.makeCompletionsRequest(request)
	if err != nil {
		return result, fmt.Errorf("cannot get response: %w", err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("cannot read response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("received invalid response")
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return result, fmt.Errorf("cannot unmarshal response: %w", err)
	}

	return result, nil
}

func (c Client) CreateChatCompletionChan(request ChatCompletionRequest, response chan ChatCompletionChunkResponse) error {
	request.Stream = true

	resp, err := c.makeCompletionsRequest(request)
	if err != nil {
		return fmt.Errorf("cannot get response: %w", err)
	}

	chunkChan := make(chan []byte)
	go readResponseChunks(resp.Body, chunkChan)

	for {
		chunk, more := <-chunkChan
		if more {
			content, _ := strings.CutPrefix(string(chunk), "data: ")
			if content == "[DONE]" {
				close(response)
				continue
			}
			var chunk ChatCompletionChunkResponse
			json.Unmarshal([]byte(content), &chunk)
			response <- chunk
		} else {
			break
		}
	}
	return nil
}

func readResponseChunks(body io.ReadCloser, chunkChan chan []byte) {
	r := bufio.NewReader(body)
	for {
		line, err := readChunkedResponseLine(r)
		if err != nil {
			if err == io.EOF {
				close(chunkChan)
				return
			}
			log.Fatal(err.Error())
		}
		if len(line) == 0 {
			continue
		}
		chunkChan <- line
	}
}

func readChunkedResponseLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	if err != nil {
		return nil, err
	}

	if isPrefix {
		rest, err := readChunkedResponseLine(r)
		if err != nil {
			return nil, err
		}
		line = append(line, rest...)
	}

	return line, nil
}

func (c Client) makeCompletionsRequest(request ChatCompletionRequest) (*http.Response, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal json: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("cannot create request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot get response: %w", err)
	}
	return resp, nil
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
