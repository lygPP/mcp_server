package weather

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

type MkWeather struct {
	ToolMap        map[string]mcp.Tool
	ToolHandlerMap map[string]func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

func NewMkWeather(ctx context.Context) *MkWeather {
	c := &MkWeather{
		ToolMap:        make(map[string]mcp.Tool),
		ToolHandlerMap: make(map[string]func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)),
	}
	c.ToolMap["GetWeather"] = mcp.NewTool("GetWeather",
		mcp.WithDescription("获取天气信息"),
		mcp.WithString("city",
			mcp.Required(),
			mcp.Description("要查询的城市名，eg：广州、深圳等"),
		),
		mcp.WithNumber("date",
			mcp.Required(),
			mcp.Description("要查询的日期，格式必须为：2025/01/01"),
		),
	)
	c.ToolHandlerMap["GetWeather"] = c.GetWeather
	return c
}

func (c *MkWeather) GetWeather(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	city := request.Params.Arguments["city"].(string)
	date := request.Params.Arguments["date"].(string)
	result := map[string]any{
		"weather":     "晴",
		"temperature": "28",
	}
	fmt.Printf("city: %s, date: %s, result: %+v", city, date, result)
	resBytes, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(resBytes)), nil
}
