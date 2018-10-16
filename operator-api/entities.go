package operator_api

type Operator struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Inputs [] Input `json:"inputs"`
	Outputs [] Output `json:"outputs"`
}

type Input struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Output struct {
	Name string `json:"name"`
	Type string `json:"type"`
}