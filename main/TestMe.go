package main

import (
	"fmt"
	"io"
)
import "net/http"

func main() {
	fmt.Println("ja vi elskar kjøtt og fårepølsa")
	http.HandleFunc("/kny", handleKny)
	http.ListenAndServe("localhost:8080", nil)
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
