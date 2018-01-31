package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	protoConvert := string("")

	err = c.Call("Course.ProtoConvert", string(""), &protoConvert)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Course.ProtoConvert:", protoConvert)
	}
}

func main() {
	client()
}

