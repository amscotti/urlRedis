# urlRedis
A simple example of using Redis to make a URL shortener

## Dependencies
* [Redigo](http://github.com/gomodule/redigo)
* [Fiber](https://github.com/gofiber/fiber)

## Build
Run `go build` to build

## Environment variable
* `PORT` - Port number to run the servers on, defaults to port `8080`
* `REDIS_URL` - URL to Redis servers, defaults to port `:6379` (Redis on the localhost on port 6379)

## Docker

### Build
`docker build -t url_redis .`

### Docker Compose
Docker compose will setup a fully working system with a Redis database,
 
Build image with `docker compose build`
Run with `docker compose up`

## Usage
`curl -X POST http://localhost:8080/v1/keys -d "url=http://128bit.io"` - Will add the URL to the database and return JSON data with the Key and URL

`curl http://localhost:8080/v1/keys/<key>` - Will return JSON data with Key and URL from the database

`curl http://localhost:8080/<key>` - Will redirect to the URL for the give Key
