package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListHumanLoopsResponse represents the ListHumanLoopsResponse schema from the OpenAPI specification
type ListHumanLoopsResponse struct {
	Humanloopsummaries interface{} `json:"HumanLoopSummaries"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteHumanLoopResponse represents the DeleteHumanLoopResponse schema from the OpenAPI specification
type DeleteHumanLoopResponse struct {
}

// HumanLoopDataAttributes represents the HumanLoopDataAttributes schema from the OpenAPI specification
type HumanLoopDataAttributes struct {
	Contentclassifiers interface{} `json:"ContentClassifiers"`
}

// DeleteHumanLoopRequest represents the DeleteHumanLoopRequest schema from the OpenAPI specification
type DeleteHumanLoopRequest struct {
}

// DescribeHumanLoopResponse represents the DescribeHumanLoopResponse schema from the OpenAPI specification
type DescribeHumanLoopResponse struct {
	Humanloopname interface{} `json:"HumanLoopName"`
	Humanloopoutput interface{} `json:"HumanLoopOutput,omitempty"`
	Humanloopstatus interface{} `json:"HumanLoopStatus"`
	Creationtime interface{} `json:"CreationTime"`
	Failurecode interface{} `json:"FailureCode,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn"`
	Humanlooparn interface{} `json:"HumanLoopArn"`
}

// DescribeHumanLoopRequest represents the DescribeHumanLoopRequest schema from the OpenAPI specification
type DescribeHumanLoopRequest struct {
}

// StartHumanLoopResponse represents the StartHumanLoopResponse schema from the OpenAPI specification
type StartHumanLoopResponse struct {
	Humanlooparn interface{} `json:"HumanLoopArn,omitempty"`
}

// ListHumanLoopsRequest represents the ListHumanLoopsRequest schema from the OpenAPI specification
type ListHumanLoopsRequest struct {
}

// StopHumanLoopRequest represents the StopHumanLoopRequest schema from the OpenAPI specification
type StopHumanLoopRequest struct {
	Humanloopname interface{} `json:"HumanLoopName"`
}

// HumanLoopOutput represents the HumanLoopOutput schema from the OpenAPI specification
type HumanLoopOutput struct {
	Outputs3uri interface{} `json:"OutputS3Uri"`
}

// StartHumanLoopRequest represents the StartHumanLoopRequest schema from the OpenAPI specification
type StartHumanLoopRequest struct {
	Dataattributes interface{} `json:"DataAttributes,omitempty"`
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn"`
	Humanloopinput interface{} `json:"HumanLoopInput"`
	Humanloopname interface{} `json:"HumanLoopName"`
}

// HumanLoopInput represents the HumanLoopInput schema from the OpenAPI specification
type HumanLoopInput struct {
	Inputcontent interface{} `json:"InputContent"`
}

// HumanLoopSummary represents the HumanLoopSummary schema from the OpenAPI specification
type HumanLoopSummary struct {
	Humanloopname interface{} `json:"HumanLoopName,omitempty"`
	Humanloopstatus interface{} `json:"HumanLoopStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Flowdefinitionarn interface{} `json:"FlowDefinitionArn,omitempty"`
}

// StopHumanLoopResponse represents the StopHumanLoopResponse schema from the OpenAPI specification
type StopHumanLoopResponse struct {
}
