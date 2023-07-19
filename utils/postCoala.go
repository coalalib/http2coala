package utils

import (
	"fmt"
	"log"

	"github.com/coalalib/coalago"
)

func PostCoala(gumAddr, addr string, data []byte) (rsp *coalago.Response, err error) {
	log.Println("gumAddr: ", gumAddr)
	log.Println("addr: ", addr)

	cl := coalago.NewClient()
	return cl.POST(data, fmt.Sprintf("coap://%s%s", gumAddr, addr))
}
