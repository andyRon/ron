package main

import (
	"flag"
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

//func main() {
//	roncache.NewGroup("scores", 2<<10, roncache.GetterFunc(
//		func(key string) ([]byte, error) {
//			log.Println("[SlowDB] search key", key)
//			if v, ok := db[key]; ok {
//				return []byte(v), nil
//			}
//			return nil, fmt.Errorf("%s not exist", key)
//		}))
//	addr := "localhost:9999"
//	peers := roncache.NewHTTPPool(addr)
//	log.Println("roncache is running at", addr)
//	log.Fatal(http.ListenAndServe(addr, peers))
//}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []string(addrs), gee)
}

func createGroup() *roncache.Group {
	return roncache.NewGroup("scores", 2<<10, roncache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, gee *roncache.Group) {
	peers := roncache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("roncache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, gee *roncache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}
