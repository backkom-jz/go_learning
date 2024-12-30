package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http:127.0.0.1:7000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	//fmt.Println(string(body))
	resData := ResponseData{}
	_ = json.Unmarshal(body, &resData)
	return resData.Data
}

// rpc 远程过程调用，如果做成本地调用
func main() {

	fmt.Println(Add(1, 2))

}
