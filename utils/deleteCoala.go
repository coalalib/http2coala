package utils

import (
	"fmt"

	"github.com/coalalib/coalago"
)

func DeleteCoala(gumAddr, addr string, data []byte) (rsp *coalago.Response, err error) {
	cl := coalago.NewClient()
	return cl.DELETE(data, fmt.Sprintf("coap://%s%s", gumAddr, addr))
}
