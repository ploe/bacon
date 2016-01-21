package main

import (
	"fmt"
	"net/http"
	"io"
	"./party"
)

func httpMethod(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func init() {
	http.HandleFunc("/party/", version)
}

// bacon --version
// localhost:8080/version
func version(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "{'desc':'early demo of bacond','state':'up','version':'0'}")
}

// bacon call [party].[method]
// localhost:8080/call/method
func partyCall(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "call: todo")
}

// bacon join party@host
// localhost:8080/join/[party]
func partyJoin(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "join: todo")
}

func main() {
	fmt.Printf("jam: comin' up... comin' up... comin' up...")
	party.Doit()
	http.HandleFunc("/", version)
	http.HandleFunc("/version/", version)
	http.HandleFunc("/register/", httpMethod)

	http.ListenAndServe(":8080", nil)
}
