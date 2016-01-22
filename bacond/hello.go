package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

//	"./party"
)

func httpMethod(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func init() {
	http.HandleFunc("/", version)
	http.HandleFunc("/version/", version)
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

// New SSL keys are generated each time the piggies are brought up, as
// it's a teensy bit more secure and it only matters that their channels
// of communication encrypted.
func keygen() {
	os.Remove("server.key")
	os.Remove("server.pem")

	key := exec.Command("openssl","genrsa","-out","server.key","2048")
	_ = key.Run()

	cert := exec.Command("openssl", "req", "-new", "-x509", "-key", "server.key", "-out", "server.pem", "-days", "3650", "-subj", "/C=GB/ST=Not Applicable/L=Somewhere/O=bacond/OU=piggy/CN=piggy.bacond")
	_ = cert.Run()

}

func main() {
	keygen()

	fmt.Printf("bacond: comin' up... comin' up... comin' up...")
	http.ListenAndServeTLS(":8888", "server.pem", "server.key", nil)
}

//	This behaviour will be integral
//	var outbuf, errbuf bytes.Buffer
//	cmd.Stdout = &outbuf
//	cmd.Stderr = &errbuf
//	stdout := outbuf.String()
//	stderr := errbuf.String()


