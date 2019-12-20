package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//Home function to serve a random number to the browser
func Home(w http.ResponseWriter, r *http.Request) {
	//rando := rand.Seed(time.Now().Unix())
	welcome := "Golang, starting fresh. \nHave a random number:"
	choiceNumber := rand.Intn(666)
	numToStr := strconv.Itoa(choiceNumber)

	w.Write([]byte(fmt.Sprintf(welcome)))
	w.Write([]byte(fmt.Sprintf(numToStr)))
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	http.ListenAndServe(":8080", mux)
}
