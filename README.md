# AI Conversation Tool

## Dependencies
To run the solution you will need the following dependencies installed

1. [Node](https://nodejs.org/en)
2. [Go](https://go.dev/dl/)

## Running application locally

### Frontend
1. Open a terminal window
2. Go to app `cd app`
3. Install node dependencies `npm install`
4. Run start the frontend `npm run dev`

### Server
1. Open a terminal window
2. Go to api `cd api`
3. Install dependencies `go mod tidy`
4. Create config file `cp config.example.json config.json`
5. Open and customize the created `config.json` file
   1. most importantly fill, the `api_key` value
   2. Some of the properties are explained [here](https://platform.openai.com/docs/api-reference/chat/create)
6. Run the server `go run main.go`

## Testing the application locally

After starting the frontend and the server, the application should be available on `http://localhost:8080`

**Important:** Changing the config will only take effect after you restart the server
