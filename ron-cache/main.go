package main

import (
	"fmt"
	"log"
	"net/http"
	"roncache/roncache"
)

var db = map[string]string{
	"Mark": "123",
	"Andy": "456",
	"Jack": "789",
}

func main() {
	roncache.NewGroup("scores", 2<<10, roncache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := roncache.NewHTTPPool(addr)
	log.Println("roncache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
