package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	router := NewRouter()
	log.Println("Listening... :80")
	log.Println("X-Auth-Token . : [" +os.Getenv("XTOKEN")+"]")
	log.Fatal(http.ListenAndServe(":80", router))
}
