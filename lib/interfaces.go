package lib

import (
	"parsing-service/flows-api"
	"parsing-service/operator-api"
)

type FlowApiService interface {
	GetFlowData(id string, userId string) (flows_api.Flow, error)
}

type OperatorApiService interface {
	GetOperator(id string) (op operator_api.Operator, err error)
}