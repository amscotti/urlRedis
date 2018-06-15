# urlRedis
A simple example of using Redis to make a URL shortener

## Dependencies
* [redigo](http://github.com/gomodule/redigo)
* [bone](https://github.com/go-zoo/bone)

## Setup
Install Godep, https://github.com/tools/godep
run `godep get` to download and install dependencies

## Usage
``curl http://localhost:8080/get/<key>`` - Will return JSON data with Key and URL from the database

``curl http://localhost:8080/<key>`` - Will redirect to the URL for the give Key

``curl -X POST http://localhost:8080/ -d "url=http://128bit.io"`` - Will add the URL to the database and return JSON data with the Key and URL
