package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-augmented-ai-runtime/mcp-server/config"
	"github.com/amazon-augmented-ai-runtime/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListhumanloopsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["CreationTimeAfter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("CreationTimeAfter=%v", val))
		}
		if val, ok := args["CreationTimeBefore"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("CreationTimeBefore=%v", val))
		}
		if val, ok := args["FlowDefinitionArn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("FlowDefinitionArn=%v", val))
		}
		if val, ok := args["SortOrder"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("SortOrder=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/human-loops#FlowDefinitionArn%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ListHumanLoopsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateListhumanloopsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_human-loops#FlowDefinitionArn",
		mcp.WithDescription("Returns information about human loops, given the specified parameters. If a human loop was deleted, it will not be included."),
		mcp.WithString("CreationTimeAfter", mcp.Description("(Optional) The timestamp of the date when you want the human loops to begin in ISO 8601 format. For example, <code>2020-02-24</code>.")),
		mcp.WithString("CreationTimeBefore", mcp.Description("(Optional) The timestamp of the date before which you want the human loops to begin in ISO 8601 format. For example, <code>2020-02-24</code>.")),
		mcp.WithString("FlowDefinitionArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of a flow definition.")),
		mcp.WithString("SortOrder", mcp.Description("Optional. The order for displaying results. Valid values: <code>Ascending</code> and <code>Descending</code>.")),
		mcp.WithString("NextToken", mcp.Description("A token to display the next page of results.")),
		mcp.WithNumber("MaxResults", mcp.Description("The total number of items to return. If the total number of available items is more than the value specified in <code>MaxResults</code>, then a <code>NextToken</code> is returned in the output. You can use this token to display the next page of results. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListhumanloopsHandler(cfg),
	}
}
