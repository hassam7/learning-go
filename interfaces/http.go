package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	url := "http://facebook.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (l logWriter) Write(b []byte) (int, error) {
	return 1, nil
}
