package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/appveyor-rest-api/mcp-server/config"
	"github.com/appveyor-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdateprojectHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ProjectWithConfiguration
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/projects", cfg.BaseURL)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
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
		var result models.Error
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

func CreateUpdateprojectTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_projects",
		mcp.WithDescription("Update project"),
		mcp.WithString("accountName", mcp.Description("")),
		mcp.WithString("name", mcp.Description("")),
		mcp.WithNumber("projectId", mcp.Description("")),
		mcp.WithString("slug", mcp.Description("")),
		mcp.WithString("created", mcp.Description("")),
		mcp.WithString("updated", mcp.Description("")),
		mcp.WithBoolean("enableSecureVariablesInPullRequestsFromSameRepo", mcp.Description("")),
		mcp.WithBoolean("isPrivate", mcp.Description("")),
		mcp.WithString("repositoryBranch", mcp.Description("Input parameter: Not present in response from addProject.")),
		mcp.WithBoolean("alwaysBuildClosedPullRequests", mcp.Description("")),
		mcp.WithBoolean("skipBranchesWithoutAppveyorYml", mcp.Description("")),
		mcp.WithBoolean("disablePushWebhooks", mcp.Description("")),
		mcp.WithString("repositoryName", mcp.Description("")),
		mcp.WithObject("securityDescriptor", mcp.Description("")),
		mcp.WithBoolean("rollingBuildsDoNotCancelRunningBuilds", mcp.Description("")),
		mcp.WithBoolean("saveBuildCacheInPullRequests", mcp.Description("")),
		mcp.WithBoolean("enableSecureVariablesInPullRequests", mcp.Description("")),
		mcp.WithBoolean("isGitHubApp", mcp.Description("")),
		mcp.WithArray("builds", mcp.Description("Input parameter: Only non-empty for response from getProjects.")),
		mcp.WithBoolean("enableDeploymentInPullRequests", mcp.Description("")),
		mcp.WithObject("nuGetFeed", mcp.Description("")),
		mcp.WithString("repositoryScm", mcp.Description("")),
		mcp.WithNumber("accountId", mcp.Description("")),
		mcp.WithNumber("currentBuildId", mcp.Description("")),
		mcp.WithBoolean("rollingBuilds", mcp.Description("")),
		mcp.WithBoolean("rollingBuildsOnlyForPullRequests", mcp.Description("")),
		mcp.WithString("repositoryType", mcp.Description("")),
		mcp.WithString("tags", mcp.Description("Input parameter: Comma-separated list of project tags for dynamic grouping.\nAppears that any input is accepted.  The returned value only\ncontains case-preserving but insensitive unique values where\nspaces around \",\" are removed but otherwise preserved.  Empty\nvalues and items are allowed.")),
		mcp.WithBoolean("disablePullRequestWebhooks", mcp.Description("")),
		mcp.WithString("statusBadgeId", mcp.Description("")),
		mcp.WithBoolean("ignoreAppveyorYml", mcp.Description("")),
		mcp.WithString("versionFormat", mcp.Description("")),
		mcp.WithString("customYmlName", mcp.Description("")),
		mcp.WithString("sshPublicKey", mcp.Description("")),
		mcp.WithString("webhookId", mcp.Description("")),
		mcp.WithNumber("buildPriority", mcp.Description("")),
		mcp.WithString("scheduleCrontabExpression", mcp.Description("Input parameter: Build schedule as an NCrontab Expression.  See https://github.com/atifaziz/NCrontab/wiki/Crontab-Expression")),
		mcp.WithString("webhookUrl", mcp.Description("")),
		mcp.WithObject("configuration", mcp.Description("")),
		mcp.WithNumber("nextBuildNumber", mcp.Description("")),
		mcp.WithString("repositoryAuthentication", mcp.Description("")),
		mcp.WithString("repositoryUsername", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateprojectHandler(cfg),
	}
}
