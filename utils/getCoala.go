package utils

import (
	"fmt"
	"log"

	"github.com/coalalib/coalago"
)

func GetCoala(gumAddr, addr string) (rsp *coalago.Response, err error) {
	log.Println("gumAddr: ", gumAddr)
	log.Println("addr: ", addr)

	cl := coalago.NewClient()
	return cl.GET(fmt.Sprintf("coap://%s%s", gumAddr, addr))
}
