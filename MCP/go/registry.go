package main

import (
	"github.com/amazon-augmented-ai-runtime/mcp-server/config"
	"github.com/amazon-augmented-ai-runtime/mcp-server/models"
	tools_human_loops "github.com/amazon-augmented-ai-runtime/mcp-server/tools/human_loops"
	tools_human_loops_flowdefinitionarn "github.com/amazon-augmented-ai-runtime/mcp-server/tools/human_loops_flowdefinitionarn"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_human_loops.CreateStarthumanloopTool(cfg),
		tools_human_loops_flowdefinitionarn.CreateListhumanloopsTool(cfg),
		tools_human_loops.CreateStophumanloopTool(cfg),
		tools_human_loops.CreateDeletehumanloopTool(cfg),
		tools_human_loops.CreateDescribehumanloopTool(cfg),
	}
}
