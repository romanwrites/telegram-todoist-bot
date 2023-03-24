package main

import (
	"log"
	"net/http"
	"os"
	"tg-function-go/todo"
)

var TodoistToken string
var ValidUserName string

func lookupEnv(key string) string {
	val, isFound := os.LookupEnv(key)
	if !isFound {
		log.Fatal(key + " was not found")
	}
	return val
}

func Handler(w http.ResponseWriter, r *http.Request) {
	TodoistToken = lookupEnv("TODOIST_TOKEN")
	ValidUserName = lookupEnv("VALID_USERNAME")

	todo.HandleToDoBotRequest(w, r, ValidUserName, TodoistToken)
}
