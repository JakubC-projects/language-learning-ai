import { Chat } from "../types/chat";
import { Message } from "../types/message";

export async function getChats(): Promise<Chat[]> {
    return []
}

export type CreateChatResponse = {
    type: 'done' | 'message' | 'chat'
    content: string
}

type SSEResponse<T = any> = {
    event: string
    data: T
}

export async function createChat(prompt: string, callback: (resp: CreateChatResponse) => any) {
    const res = await fetch("http://localhost:8081/chats", {
        method: "POST",
        headers: {
          'Content-Type': 'text/event-stream'
        },
        body: JSON.stringify({
            prompt
        })
    })
    if (!res.body) {
        throw Error("invalid response")
    }

    const reader = res.body.pipeThrough(new TextDecoderStream()).getReader()
    while (true) {
      const {value, done} = await reader.read();
      if (done) {
        callback({type: "done", content: ""})
        break;
      }
      const event = parseSSE<Message>(value)

      callback({type:"message", content: event.data.content})
    }
}

function parseSSE<T = any>(value: string):SSEResponse<T> {
    const regex = /event:([a-zA-Z0-9]*)\ndata:"([a-zA-Z0-9=]*)"/
    const res = regex.exec(value)
    if(!res || res.length < 3) {
        throw Error("invalid response")
    }
    const event = res[1]
    const data = JSON.parse(atob(res[2]))

    return {event, data}
}