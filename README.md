# AI Conversation Tool

## Dependencies
To run the solution you will need the following dependencies installed

1. [Node](https://nodejs.org/en)
2. [Go](https://go.dev/dl/)

## Running application locally

### Frontend
1. Go to app `cd app`
2. Install node dependencies `npm install`
3. Run start the frontend `npm run dev`

### Server
1. Go to api `cd api`
2. Install dependencies `go mod tidy`
3. Create config file `cp config.example.json config.json`
4. Fill the config file, most importantly fill, the `api_key` value
5. Run the server `go run main.go`

## Testing the application locally

After starting the frontend and the server, the application should be available on `http://localhost:8080`
