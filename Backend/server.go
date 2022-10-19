package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const Port = 3000
const rootPath = "../Frontend/"

func kitchenHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path != "/kitchen" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	path = rootPath + path
	http.ServeFile(w, r, path)
	fmt.Println("Access to kitchen")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(rootPath)))
	http.HandleFunc("/kitchen", kitchenHandler)
	//http.HandleFunc("/kitchen", kitchenHandler)

	fmt.Printf("Starting server at port %v\n", Port)
	portAddress := ":" + strconv.Itoa(Port)
	if err := http.ListenAndServe(portAddress, nil); err != nil {
		log.Fatal(err)
	}
}
