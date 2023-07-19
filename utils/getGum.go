package utils

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func GetGum(cid string) (gum string, err error) {
	_, resp, err := fasthttp.Get(nil, fmt.Sprintf("%s?cid=%s", "http://dev.keenetic.cloud:8081", cid))
	if err != nil {
		return
	}

	gum = string(resp)
	return
}
