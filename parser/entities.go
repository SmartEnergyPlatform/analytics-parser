package parser

type Pipeline map[int] Operator

type Operator struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ImageId string
	InputTopics map [string] InputTopic
}

type InputTopic struct {
	FilterType string
	FilterValue string
	Mappings [] Mapping
}

type Mapping struct {
	Source string
	Dest string
}