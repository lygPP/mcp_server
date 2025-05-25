package clothing

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

type MkClothes struct {
	ToolMap        map[string]mcp.Tool
	ToolHandlerMap map[string]func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

func NewMkClothes(ctx context.Context) *MkClothes {
	c := &MkClothes{
		ToolMap:        make(map[string]mcp.Tool),
		ToolHandlerMap: make(map[string]func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}
	c.ToolMap["GetClothes"] = mcp.NewTool("GetClothes",
		mcp.WithDescription("获取当前衣橱里的衣服信息"),
	)
	c.ToolHandlerMap["GetClothes"] = c.GetClothes
	return c
}

func (c *MkClothes) GetClothes(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result := make([]map[string]any, 0)
	result = append(result, map[string]any{
		"type": "T-shirt",
		"name": "T1",
	})
	result = append(result, map[string]any{
		"type": "T-shirt",
		"name": "T1",
	})
	result = append(result, map[string]any{
		"type": "jeans",
		"name": "J1",
	})
	result = append(result, map[string]any{
		"type": "Sun hat",
		"name": "SH1",
	})
	result = append(result, map[string]any{
		"type": "Basketball shoes",
		"name": "AJ1",
	})
	result = append(result, map[string]any{
		"type": "skate shoes",
		"name": "SK1",
	})
	fmt.Printf("result: %+v", result)
	resBytes, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(resBytes)), nil
}
