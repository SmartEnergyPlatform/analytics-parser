package flows_api

type Flow struct {
	Id    string `json:"_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Model Model  `json:"model,omitempty"`
}

type Model struct {
	Edges []Edge `json:"edges,omitempty"`
	Nodes []Node `json:"nodes,omitempty"`
}

type Edge struct {
	Destination int `json:"destination,omitempty"`
	Source      int `json:"source,omitempty"`
}

type Node struct {
	Id         int         `json:"id, omitempty"`
	Name       string      `json:"name,omitempty"`
	ImageId    string      `json:"imageId"`
	Connectors []Connector `json:"connectors"`
}

type Connector struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Value struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
