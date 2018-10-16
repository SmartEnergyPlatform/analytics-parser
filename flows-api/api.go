package flows_api

import (
	"github.com/parnurzeal/gorequest"
	"encoding/json"
)

type FlowApi struct {
	url string
}

func NewFlowApi(url string) *FlowApi {
	return &FlowApi{url}
}

func (f FlowApi) GetFlowData(id string, userId string) (flow Flow, err error) {
	request := gorequest.New()
	_, body, _ := request.Get(f.url+"flow/" + id + "/").Set("X-UserId", userId).End()
	err = json.Unmarshal([]byte(body), &flow)
	return
}