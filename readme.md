# urlRedis
A simple example of using Redis to make a URL shortener

## Dependencies
* [Redigo](http://github.com/gomodule/redigo)
* [Gorilla Mux](https://github.com/gorilla/mux)

## Build
Run `go build` to build

## Environment variable
* `PORT` - Port number to run the servers on, defaults to port `8081`
* `REDIS_URL` - URL to Redis servers, defaults to port `:6379` (Redis on the localhost on port 6379)

## Usage
`curl http://localhost:8080/get/<key>` - Will return JSON data with Key and URL from the database

`curl http://localhost:8080/<key>` - Will redirect to the URL for the give Key

`curl -X POST http://localhost:8080/ -d "url=http://128bit.io"` - Will add the URL to the database and return JSON data with the Key and URL
