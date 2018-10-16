package parser

import (
	"parsing-service/lib"
	"parsing-service/flows-api"
	"fmt"
	"strconv"
)

type FlowParser struct {
	flowApi lib.FlowApiService
	operatorApi lib.OperatorApiService
}

func NewFlowParser (flowApi lib.FlowApiService, operatorApi lib.OperatorApiService) * FlowParser {

	return &FlowParser{flowApi, operatorApi}
}

func (f FlowParser) ParseFlow (id string, userId  string) Pipeline {
	var pipeline =  make(Pipeline)
	flow, _ := f.flowApi.GetFlowData(id, userId)

	// Create Nodes
	for _ , node := range flow.Model.Nodes {
		var operator = Operator{node.Id, node.Name, node.ImageId, make(map [string] InputTopic)}
		pipeline[node.Id] =  operator
	}

	// Create Input Topics
	for _, edge := range flow.Model.Edges{
		var sourceConnector = flow.Model.GetConnectorById(edge.Source)
		var destinationConnector = flow.Model.GetConnectorById(edge.Destination)

		// Make mapping
		var mapping = Mapping{}

		//get name of output topic
		var outputTopic = ""

		//TODO: Needs refactoring
		if flow.Model.GetConnectorById(edge.Source).Type != "topConnector" {
			mapping = Mapping{sourceConnector.Value.Name, destinationConnector.Value.Name}
			outputTopic = getOperatorOutputTopic(pipeline[flow.Model.GetNodeIdByConnectorId(edge.Source)].Name)
			topic := pipeline[flow.Model.GetNodeIdByConnectorId(edge.Destination)].InputTopics[outputTopic]
			topic.FilterValue = strconv.Itoa(pipeline[flow.Model.GetNodeIdByConnectorId(edge.Source)].Id)
			topic.FilterType = "OperatorId"
			topic.Mappings = append(topic.Mappings, mapping)
			pipeline[flow.Model.GetNodeIdByConnectorId(edge.Destination)].InputTopics[outputTopic] = topic
		} else {
			mapping = Mapping{destinationConnector.Value.Name, sourceConnector.Value.Name}
			outputTopic = getOperatorOutputTopic(pipeline[flow.Model.GetNodeIdByConnectorId(edge.Destination)].Name)
			topic := pipeline[flow.Model.GetNodeIdByConnectorId(edge.Source)].InputTopics[outputTopic]
			topic.Mappings = append(topic.Mappings, mapping)
			pipeline[flow.Model.GetNodeIdByConnectorId(edge.Source)].InputTopics[outputTopic] = topic
		}
	}
	return pipeline
}

func (f FlowParser) GetInputs (id string, userId string) ([] flows_api.Node) {
	flow, _ := f.flowApi.GetFlowData(id, userId)
	fmt.Println(flow)
	return flow.Model.GetEmptyNodeInputs()
}

func getOperatorOutputTopic (name string) (op_name string) {
	op_name = "analytics-" + name
	return
}
