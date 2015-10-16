package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/amscotti/urlRedis/Godeps/_workspace/src/github.com/garyburd/redigo/redis"
	"github.com/amscotti/urlRedis/Godeps/_workspace/src/github.com/gorilla/mux"
)

type status struct {
	Key   string
	URL   string
	Count int
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

var pool = newPool()

func getid() string {
	c := pool.Get()
	defer c.Close()

	value, _ := redis.Int(c.Do("INCR", "IDCOUNT"))

	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(value)))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", create).Methods("POST")
	router.HandleFunc("/{key}", redirect).Methods("GET")
	router.HandleFunc("/get/{key}", get).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	c := pool.Get()
	defer c.Close()

	id := getid()

	c.Do("SET", id, html.EscapeString(r.FormValue("url")))
	json, err := json.Marshal(status{id, r.FormValue("url"), 0})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func get(w http.ResponseWriter, r *http.Request) {
	c := pool.Get()
	defer c.Close()
	vars := mux.Vars(r)
	key := vars["key"]

	count, err := redis.Int(c.Do("GET", fmt.Sprintf("%q_COUNT", key)))
	if err != nil {
		count = 0
	}

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json, err := json.Marshal(status{key, value, count})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	c := pool.Get()
	defer c.Close()
	vars := mux.Vars(r)
	key := vars["key"]

	c.Do("INCR", fmt.Sprintf("%q_COUNT", key))
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		http.Redirect(w, r, html.EscapeString(value), 301)
	}
}
