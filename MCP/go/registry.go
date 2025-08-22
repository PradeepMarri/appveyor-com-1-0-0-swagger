package main

import (
	"github.com/appveyor-rest-api/mcp-server/config"
	"github.com/appveyor-rest-api/mcp-server/models"
	tools_project "github.com/appveyor-rest-api/mcp-server/tools/project"
	tools_collaborator "github.com/appveyor-rest-api/mcp-server/tools/collaborator"
	tools_deployment "github.com/appveyor-rest-api/mcp-server/tools/deployment"
	tools_environment "github.com/appveyor-rest-api/mcp-server/tools/environment"
	tools_user "github.com/appveyor-rest-api/mcp-server/tools/user"
	tools_build "github.com/appveyor-rest-api/mcp-server/tools/build"
	tools_role "github.com/appveyor-rest-api/mcp-server/tools/role"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_project.CreateEncryptvalueTool(cfg),
		tools_collaborator.CreateGetcollaboratorsTool(cfg),
		tools_collaborator.CreateUpdatecollaboratorTool(cfg),
		tools_deployment.CreateCanceldeploymentTool(cfg),
		tools_environment.CreateGetenvironmentsTool(cfg),
		tools_environment.CreateAddenvironmentTool(cfg),
		tools_environment.CreateUpdateenvironmentTool(cfg),
		tools_user.CreateGetusersTool(cfg),
		tools_user.CreateUpdateuserTool(cfg),
		tools_environment.CreateDeleteenvironmentTool(cfg),
		tools_project.CreateGetprojectbuildbyversionTool(cfg),
		tools_user.CreateDeleteuserTool(cfg),
		tools_user.CreateGetuserTool(cfg),
		tools_build.CreateGetbuildartifactsTool(cfg),
		tools_project.CreateGetprojectsTool(cfg),
		tools_project.CreateAddprojectTool(cfg),
		tools_project.CreateUpdateprojectTool(cfg),
		tools_role.CreateDeleteroleTool(cfg),
		tools_role.CreateGetroleTool(cfg),
		tools_project.CreateGetprojectlastbuildbranchTool(cfg),
		tools_deployment.CreateStartdeploymentTool(cfg),
		tools_collaborator.CreateDeletecollaboratorTool(cfg),
		tools_collaborator.CreateGetcollaboratorTool(cfg),
		tools_project.CreateGetprojectdeploymentsTool(cfg),
		tools_deployment.CreateGetdeploymentTool(cfg),
		tools_user.CreateCanceluserinvitationTool(cfg),
		tools_build.CreateCancelbuildTool(cfg),
		tools_project.CreateDeleteprojectTool(cfg),
		tools_project.CreateGetprojectlastbuildTool(cfg),
		tools_project.CreateGetprojectsettingsyamlTool(cfg),
		tools_project.CreateUpdateprojectsettingsyamlTool(cfg),
		tools_project.CreateUpdateprojectbuildnumberTool(cfg),
		tools_user.CreateJoinaccountTool(cfg),
		tools_environment.CreateGetenvironmentdeploymentsTool(cfg),
		tools_project.CreateDeleteprojectbuildcacheTool(cfg),
		tools_role.CreateGetrolesTool(cfg),
		tools_role.CreateAddroleTool(cfg),
		tools_role.CreateUpdateroleTool(cfg),
		tools_project.CreateGetprojectenvironmentvariablesTool(cfg),
		tools_project.CreateUpdateprojectenvironmentvariablesTool(cfg),
		tools_build.CreateStartbuildTool(cfg),
		tools_build.CreateRerunbuildTool(cfg),
		tools_project.CreateGetprojectsettingsTool(cfg),
		tools_environment.CreateGetenvironmentsettingsTool(cfg),
		tools_project.CreateGetprojecthistoryTool(cfg),
		tools_user.CreateGetuserinvitationsTool(cfg),
		tools_user.CreateInviteuserTool(cfg),
	}
}
