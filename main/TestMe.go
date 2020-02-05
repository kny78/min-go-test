package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io"
)
import "net/http"

func main() {
	fmt.Println("ja vi elskar kjøtt og fårepølsa")
	buf, err := json2.Marshal(map[string]string{
		"name": "Kjetil Nygård",
		"email": "pulk.hesten@gmail.com",
		"source_code": "https://github.com/kny78/min-go-test.git",
	})
	if err != nil {
		println("Error found")
	}
	buf2 := bytes.NewBuffer( buf)
	client:=http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	client.Post("http://10.200.233.89:8080/contest", "application/json", buf2)
}

type Hest struct {
	Navn string
	Vekt int64
}

func (h Hest) data() string {
	return fmt.Sprintf("%s (%d kg)\n\n", h.Navn, h.Vekt)
}

func (h *Hest) flipName(ret chan string) {
	for i := 0; i < 10; i++ {
		ret <- fmt.Sprint("%s%d\n\n", "goo", i)
	}
}

func handleKny(writer http.ResponseWriter, request *http.Request) {
	kjell := Hest{"Rex Rodney", 432}

	r := make(chan string, 1000)
	go kjell.flipName(r)

	io.WriteString(writer, kjell.data())
	for key := range request.Header {
		value := request.Header.Get(key)
		io.WriteString(writer, key+": "+value+"\n")
	}
	for i := 0; i < 9; i++ {
		f := <-r
		io.WriteString(writer, f+"\r\n")
	}

	//io.Copy(writer, request.Body)
}
