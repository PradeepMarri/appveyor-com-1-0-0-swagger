package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/appveyor-rest-api/mcp-server/config"
	"github.com/appveyor-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetprojecthistoryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["recordsNumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("recordsNumber=%v", val))
		}
		if val, ok := args["startBuildId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startBuildId=%v", val))
		}
		if val, ok := args["branch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("branch=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/projects/%s/%s/history%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
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
		var result models.ProjectHistory
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

func CreateGetprojecthistoryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_projects_accountName_projectSlug_history",
		mcp.WithDescription("Get project history"),
		mcp.WithNumber("recordsNumber", mcp.Required(), mcp.Description("Number of results to include in the response. getProjectDeployments is documented to have a maximum of 20. It currently returns 500 Internal Server Error for recordsNumber <= 5. In the past it has returned 500 Internal Server Error for many different values which did not match the value used by the ci.appveyor.com web interface at the time.  As of 2018-09-08, the value used by the web interface is 10.")),
		mcp.WithNumber("startBuildId", mcp.Description("Maximum `buildId` to include in the results (exclusive).")),
		mcp.WithString("branch", mcp.Description("Repository Branch")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetprojecthistoryHandler(cfg),
	}
}
