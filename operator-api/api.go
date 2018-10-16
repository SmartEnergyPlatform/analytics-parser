package operator_api

import (
	"github.com/parnurzeal/gorequest"
	"encoding/json"
)

type OperatorApi struct {
	url string
}

func NewOperatorApi(url string) *OperatorApi {
	return &OperatorApi{url}
}

func (o OperatorApi) GetOperator(id string) (op Operator, err error) {
	request := gorequest.New()
	_, body, _ := request.Get(o.url + "operator/" + id + "/").End()
	err = json.Unmarshal([]byte(body), &op)
	return
}
