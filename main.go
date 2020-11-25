package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://localhost:12345/")

	connectDB()

	route := mux.NewRouter()

	API(route)

	log.Fatal(http.ListenAndServe(":12345", route))
}
