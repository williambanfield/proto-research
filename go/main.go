package main

import (
	"bytes"
	"io"
	"os"

	"github.com/williambanfield/prototest/protos"
)

func main() {
	f, err := os.Open("./output")
	msg := protos.AwesomeMessage{AwesomeField: "", AwesomeBool: false}
	buf, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	io.Copy(f, bytes.NewBuffer(buf))
}
