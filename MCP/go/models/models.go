package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path"` // Path glob of artifact files. Note that updateProject does not require path, but omitting path results in post-build error "Collecting artifacts... Value cannot be null. Parameter name: input"
	TypeField string `json:"type,omitempty"` // Possible values from `Push-AppveyorArtifact` cmdlet `-Type` parameter. The web UI only allows selection of `Auto`, `WebDeployPackage`, and unspecified (which it sends as the empty string but functions like omitting the property) for `updateProject`. Any string can be sent and will be saved/returned, but only these values have a function (as far as I am aware).
}

// ProjectHistory represents the ProjectHistory schema from the OpenAPI specification
type ProjectHistory struct {
	Builds []Build `json:"builds,omitempty"`
	Project Project `json:"project,omitempty"`
}

// CollaboratorUpdate represents the CollaboratorUpdate schema from the OpenAPI specification
type CollaboratorUpdate struct {
	Roleid int `json:"roleId"`
	Userid int `json:"userId"`
}

// InviteUserRequest represents the InviteUserRequest schema from the OpenAPI specification
type InviteUserRequest struct {
	Roleid int `json:"roleId"`
	Email string `json:"email"`
}

// NuGetFeed represents the NuGetFeed schema from the OpenAPI specification
type NuGetFeed struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Projectid int `json:"projectId,omitempty"`
	Publishingenabled bool `json:"publishingEnabled,omitempty"`
	Accountid int `json:"accountId,omitempty"`
	Id string `json:"id,omitempty"`
	Isprivateproject bool `json:"isPrivateProject,omitempty"`
	Name string `json:"name,omitempty"`
	Nugetfeedid int `json:"nuGetFeedId,omitempty"`
}

// SessionUserModel represents the SessionUserModel schema from the OpenAPI specification
type SessionUserModel struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Twofactorauthenabled bool `json:"twoFactorAuthEnabled,omitempty"`
	Vsousername string `json:"vsoUsername,omitempty"`
	Bitbucketusername string `json:"bitBucketUsername,omitempty"`
	Fullname string `json:"fullName,omitempty"`
	Gitlabuserid string `json:"gitLabUserId,omitempty"`
	Gravatarhash string `json:"gravatarHash,omitempty"`
	Email string `json:"email"`
	Userid int `json:"userId"`
	Githubusername string `json:"gitHubUsername,omitempty"`
	Pagesize int `json:"pageSize,omitempty"`
}

// BuildWorkerImage represents the BuildWorkerImage schema from the OpenAPI specification
type BuildWorkerImage struct {
	Buildworkerimageid int `json:"buildWorkerImageId"`
	Name string `json:"name"` // Defines the available build worker image templates used to provision a virtual machine for a build. Images are updated regularly. "Previous" selects the previous version of an image, for use as a temporary workaround for regressions. `Ubuntu` is the same as `Ubuntu1604`. `Previous Ubuntu` is the same as `Previous Ubuntu1604`. See https://www.appveyor.com/docs/build-environment/#build-worker-images for details.
	Ostype string `json:"osType,omitempty"`
	Buildcloudname string `json:"buildCloudName,omitempty"`
}

// BuildStartRequest represents the BuildStartRequest schema from the OpenAPI specification
type BuildStartRequest struct {
	Environmentvariables map[string]interface{} `json:"environmentVariables,omitempty"`
	Projectslug string `json:"projectSlug"`
	Pullrequestid int `json:"pullRequestId,omitempty"` // Can not be used with `branch` or `commitId`
	Accountname string `json:"accountName"`
	Branch string `json:"branch,omitempty"`
	Commitid string `json:"commitId,omitempty"`
}

// ProjectSettingsResults represents the ProjectSettingsResults schema from the OpenAPI specification
type ProjectSettingsResults struct {
	Buildclouds []StringValueObject `json:"buildClouds,omitempty"`
	Defaultimagename string `json:"defaultImageName,omitempty"`
	Images []BuildWorkerImage `json:"images,omitempty"`
	Project Project `json:"project,omitempty"`
	Settings ProjectWithConfiguration `json:"settings,omitempty"`
}

// JoinAccountRequest represents the JoinAccountRequest schema from the OpenAPI specification
type JoinAccountRequest struct {
	Invitationid string `json:"invitationId"`
}

// UserInvitationModel represents the UserInvitationModel schema from the OpenAPI specification
type UserInvitationModel struct {
	Accountid int `json:"accountId"`
	Accountname string `json:"accountName,omitempty"`
	Created string `json:"created"`
	Email string `json:"email"`
	Roleid int `json:"roleId"`
	Rolename string `json:"roleName,omitempty"`
	Userinvitationid string `json:"userInvitationId"`
}

// StoredNameValue represents the StoredNameValue schema from the OpenAPI specification
type StoredNameValue struct {
	Name string `json:"name"`
	Value StoredValue `json:"value"`
}

// DeploymentEnvironmentWithSettings represents the DeploymentEnvironmentWithSettings schema from the OpenAPI specification
type DeploymentEnvironmentWithSettings struct {
	Deploymentenvironmentid int `json:"deploymentEnvironmentId,omitempty"`
	Name string `json:"name,omitempty"`
	Provider string `json:"provider,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Projectsmode int `json:"projectsMode,omitempty"` // Project selection mode for deployment environments.
	Securitydescriptor SecurityDescriptor `json:"securityDescriptor,omitempty"`
	Tags string `json:"tags,omitempty"` // Comma-separated list of environment tags for dynamic grouping. Appears that any input is accepted. The returned value only contains case-preserving but insensitive unique values where spaces around "," are removed but otherwise preserved. Empty values and items are allowed.
	Accountid int `json:"accountId,omitempty"`
	Selectedprojects []int `json:"selectedProjects,omitempty"` // Project IDs of selected projects
	Settings DeploymentEnvironmentSettings `json:"settings,omitempty"`
	Environmentaccesskey string `json:"environmentAccessKey,omitempty"`
	Projects []DeploymentEnvironmentProject `json:"projects,omitempty"` // Projects available for selection in UI. Only present in response from getEnvironmentSettings.
}

// ProjectWithConfiguration represents the ProjectWithConfiguration schema from the OpenAPI specification
type ProjectWithConfiguration struct {
	Projectid int `json:"projectId"`
	Slug string `json:"slug,omitempty"`
	Accountname string `json:"accountName,omitempty"`
	Name string `json:"name,omitempty"`
	Updated string `json:"updated,omitempty"`
	Created string `json:"created,omitempty"`
	Enablesecurevariablesinpullrequestsfromsamerepo bool `json:"enableSecureVariablesInPullRequestsFromSameRepo,omitempty"`
	Isprivate bool `json:"isPrivate,omitempty"`
	Repositorybranch string `json:"repositoryBranch,omitempty"` // Not present in response from addProject.
	Alwaysbuildclosedpullrequests bool `json:"alwaysBuildClosedPullRequests,omitempty"`
	Skipbrancheswithoutappveyoryml bool `json:"skipBranchesWithoutAppveyorYml,omitempty"`
	Disablepushwebhooks bool `json:"disablePushWebhooks,omitempty"`
	Repositoryname string `json:"repositoryName"`
	Securitydescriptor SecurityDescriptor `json:"securityDescriptor,omitempty"`
	Rollingbuildsdonotcancelrunningbuilds bool `json:"rollingBuildsDoNotCancelRunningBuilds,omitempty"`
	Savebuildcacheinpullrequests bool `json:"saveBuildCacheInPullRequests,omitempty"`
	Enablesecurevariablesinpullrequests bool `json:"enableSecureVariablesInPullRequests,omitempty"`
	Isgithubapp bool `json:"isGitHubApp,omitempty"`
	Builds []Build `json:"builds,omitempty"` // Only non-empty for response from getProjects.
	Enabledeploymentinpullrequests bool `json:"enableDeploymentInPullRequests,omitempty"`
	Nugetfeed NuGetFeed `json:"nuGetFeed,omitempty"`
	Repositoryscm string `json:"repositoryScm,omitempty"`
	Accountid int `json:"accountId,omitempty"`
	Currentbuildid int `json:"currentBuildId,omitempty"`
	Rollingbuilds bool `json:"rollingBuilds,omitempty"`
	Rollingbuildsonlyforpullrequests bool `json:"rollingBuildsOnlyForPullRequests,omitempty"`
	Repositorytype string `json:"repositoryType,omitempty"`
	Tags string `json:"tags,omitempty"` // Comma-separated list of project tags for dynamic grouping. Appears that any input is accepted. The returned value only contains case-preserving but insensitive unique values where spaces around "," are removed but otherwise preserved. Empty values and items are allowed.
	Disablepullrequestwebhooks bool `json:"disablePullRequestWebhooks,omitempty"`
	Repositoryusername string `json:"repositoryUsername,omitempty"`
	Statusbadgeid string `json:"statusBadgeId,omitempty"`
	Ignoreappveyoryml bool `json:"ignoreAppveyorYml,omitempty"`
	Versionformat string `json:"versionFormat"`
	Customymlname string `json:"customYmlName,omitempty"`
	Sshpublickey string `json:"sshPublicKey,omitempty"`
	Webhookid string `json:"webhookId,omitempty"`
	Buildpriority int `json:"buildPriority,omitempty"`
	Schedulecrontabexpression string `json:"scheduleCrontabExpression,omitempty"` // Build schedule as an NCrontab Expression. See https://github.com/atifaziz/NCrontab/wiki/Crontab-Expression
	Webhookurl string `json:"webhookUrl,omitempty"`
	Configuration ProjectConfiguration `json:"configuration"`
	Nextbuildnumber int `json:"nextBuildNumber,omitempty"`
	Repositoryauthentication string `json:"repositoryAuthentication,omitempty"`
}

// RoleAce represents the RoleAce schema from the OpenAPI specification
type RoleAce struct {
	Accessrights []AceAccessRight `json:"accessRights,omitempty"`
	Isadmin bool `json:"isAdmin,omitempty"`
	Name string `json:"name,omitempty"`
	Roleid int `json:"roleId,omitempty"`
}

// StoredValue represents the StoredValue schema from the OpenAPI specification
type StoredValue struct {
	Isencrypted bool `json:"isEncrypted,omitempty"`
	Value string `json:"value,omitempty"` // Encrypted values can be created using the encryptValue operation. Empty environment variables are represented by missing (null) value rather than an empty string.
}

// AceAccessRightDefinition represents the AceAccessRightDefinition schema from the OpenAPI specification
type AceAccessRightDefinition struct {
	Description string `json:"description,omitempty"`
	Name string `json:"name"`
}

// SessionModel represents the SessionModel schema from the OpenAPI specification
type SessionModel struct {
	Setuprequired bool `json:"setupRequired,omitempty"`
	Twofactorauthrequired bool `json:"twoFactorAuthRequired,omitempty"`
	User SessionUserModel `json:"user,omitempty"`
	Accounts []SessionUserAccountModel `json:"accounts,omitempty"`
}

// BuildMessage represents the BuildMessage schema from the OpenAPI specification
type BuildMessage struct {
	Category string `json:"category,omitempty"`
	Created string `json:"created,omitempty"`
	Message string `json:"message,omitempty"`
}

// StoredNameValueMatrix represents the StoredNameValueMatrix schema from the OpenAPI specification
type StoredNameValueMatrix struct {
	Variables StoredNameValue `json:"variables,omitempty"`
}

// ProjectDeployment represents the ProjectDeployment schema from the OpenAPI specification
type ProjectDeployment struct {
	Project Project `json:"project"`
	Deployment Deployment `json:"deployment"`
}

// DeploymentJob represents the DeploymentJob schema from the OpenAPI specification
type DeploymentJob struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Status string `json:"status,omitempty"`
	Finished string `json:"finished,omitempty"`
	Jobid string `json:"jobId,omitempty"`
	Name string `json:"name,omitempty"`
	Started string `json:"started,omitempty"`
	Messagescount int `json:"messagesCount,omitempty"`
}

// ProjectAddition represents the ProjectAddition schema from the OpenAPI specification
type ProjectAddition struct {
	Repositoryauthentication string `json:"repositoryAuthentication,omitempty"`
	Repositoryname string `json:"repositoryName"` // URL when repositoryProvider is git, mercurial, subversion username/project when repositoryProvider is gitHub
	Repositorypassword string `json:"repositoryPassword,omitempty"` // Required if repositoryAuthentication is credentials
	Repositoryprovider string `json:"repositoryProvider"`
	Repositoryusername string `json:"repositoryUsername,omitempty"` // Required if repositoryAuthentication is credentials
}

// UserAccountSettings represents the UserAccountSettings schema from the OpenAPI specification
type UserAccountSettings struct {
	Successfuldeploymentnotification string `json:"successfulDeploymentNotification"`
	Failedbuildnotification string `json:"failedBuildNotification"`
	Faileddeploymentnotification string `json:"failedDeploymentNotification"`
	Notifywhenbuildstatuschangedonly bool `json:"notifyWhenBuildStatusChangedOnly,omitempty"` // Note that this value is `true` on user creation, but behaves as `false` when not specified on update.
	Notifywhendeploymentstatuschangedonly bool `json:"notifyWhenDeploymentStatusChangedOnly,omitempty"` // Note that this value is `true` on user creation, but behaves as `false` when not specified on update.
	Successfulbuildnotification string `json:"successfulBuildNotification"`
}

// DeploymentEnvironmentAddition represents the DeploymentEnvironmentAddition schema from the OpenAPI specification
type DeploymentEnvironmentAddition struct {
	Name string `json:"name"`
	Provider string `json:"provider"`
	Settings DeploymentEnvironmentSettings `json:"settings"`
}

// DeploymentEnvironmentProject represents the DeploymentEnvironmentProject schema from the OpenAPI specification
type DeploymentEnvironmentProject struct {
	Projectid int `json:"projectId"`
	Isselected bool `json:"isSelected"`
	Name string `json:"name"`
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Finished string `json:"finished,omitempty"`
	Jobid string `json:"jobId,omitempty"`
	Name string `json:"name,omitempty"`
	Started string `json:"started,omitempty"`
	Status string `json:"status,omitempty"`
}

// DeploymentEnvironmentDeploymentsResults represents the DeploymentEnvironmentDeploymentsResults schema from the OpenAPI specification
type DeploymentEnvironmentDeploymentsResults struct {
	Deployments []EnvironmentDeploymentModel `json:"deployments"`
	Environment DeploymentEnvironment `json:"environment"`
}

// AceAccessRight represents the AceAccessRight schema from the OpenAPI specification
type AceAccessRight struct {
	Allowed bool `json:"allowed,omitempty"` // true to allow, false to deny, undefined to inherit
	Name string `json:"name"`
}

// SecurityDescriptor represents the SecurityDescriptor schema from the OpenAPI specification
type SecurityDescriptor struct {
	Accessrightdefinitions []AceAccessRightDefinition `json:"accessRightDefinitions,omitempty"`
	Roleaces []RoleAce `json:"roleAces,omitempty"`
}

// ArtifactModel represents the ArtifactModel schema from the OpenAPI specification
type ArtifactModel struct {
	Filename string `json:"fileName,omitempty"`
	Name string `json:"name,omitempty"`
	Size int `json:"size,omitempty"`
	TypeField string `json:"type,omitempty"` // Possible values from `Push-AppveyorArtifact` cmdlet `-Type` parameter. The web UI only allows selection of `Auto`, `WebDeployPackage`, and unspecified (which it sends as the empty string but functions like omitting the property) for `updateProject`. Any string can be sent and will be saved/returned, but only these values have a function (as far as I am aware).
	Url string `json:"url,omitempty"` // This property has not been observed in JSON responses, but is present and nil in XML responses.
	Created string `json:"created,omitempty"`
}

// RoleWithGroups represents the RoleWithGroups schema from the OpenAPI specification
type RoleWithGroups struct {
	Updated string `json:"updated,omitempty"`
	Created string `json:"created,omitempty"`
	Issystem bool `json:"isSystem,omitempty"`
	Name string `json:"name"`
	Roleid int `json:"roleId"`
	Groups []GroupPermissions `json:"groups,omitempty"`
}

// EncryptRequest represents the EncryptRequest schema from the OpenAPI specification
type EncryptRequest struct {
	Plainvalue string `json:"plainValue,omitempty"`
}

// EnvironmentDeploymentModel represents the EnvironmentDeploymentModel schema from the OpenAPI specification
type EnvironmentDeploymentModel struct {
	Status string `json:"status,omitempty"`
	Build BuildLookupModel `json:"build,omitempty"`
	Deploymentid int `json:"deploymentId,omitempty"`
	Finished string `json:"finished,omitempty"`
	Started string `json:"started,omitempty"`
	Project ProjectLookupModel `json:"project,omitempty"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Build BuildLookupModel `json:"build,omitempty"`
	Deploymentid int `json:"deploymentId,omitempty"`
	Finished string `json:"finished,omitempty"`
	Started string `json:"started,omitempty"`
	Status string `json:"status,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Environment DeploymentEnvironment `json:"environment,omitempty"`
	Jobs []DeploymentJob `json:"jobs,omitempty"`
	Build Build `json:"build,omitempty"`
}

// ProjectBuildNumberUpdate represents the ProjectBuildNumberUpdate schema from the OpenAPI specification
type ProjectBuildNumberUpdate struct {
	Nextbuildnumber int `json:"nextBuildNumber"`
}

// StringValueObject represents the StringValueObject schema from the OpenAPI specification
type StringValueObject struct {
	Value string `json:"value,omitempty"`
}

// Build represents the Build schema from the OpenAPI specification
type Build struct {
	Branch string `json:"branch,omitempty"`
	Buildid int `json:"buildId,omitempty"`
	Message string `json:"message,omitempty"`
	Version string `json:"version,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Authorname string `json:"authorName,omitempty"`
	Messageextended string `json:"messageExtended,omitempty"`
	Committed string `json:"committed,omitempty"`
	Pullrequestname string `json:"pullRequestName,omitempty"`
	Projectid int `json:"projectId,omitempty"`
	Authorusername string `json:"authorUsername,omitempty"`
	Finished string `json:"finished,omitempty"`
	Jobs []BuildJob `json:"jobs,omitempty"` // Always empty in getProjectHistory and startDeployment responses.
	Buildnumber int `json:"buildNumber,omitempty"`
	Committerusername string `json:"committerUsername,omitempty"`
	Istag bool `json:"isTag,omitempty"`
	Commitid string `json:"commitId,omitempty"`
	Status string `json:"status,omitempty"`
	Pullrequestid int `json:"pullRequestId,omitempty"`
	Committername string `json:"committerName,omitempty"`
	Messages []BuildMessage `json:"messages,omitempty"`
	Started string `json:"started,omitempty"`
}

// Role represents the Role schema from the OpenAPI specification
type Role struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Issystem bool `json:"isSystem,omitempty"`
	Name string `json:"name"`
	Roleid int `json:"roleId"`
}

// SessionUserAccountModel represents the SessionUserAccountModel schema from the OpenAPI specification
type SessionUserAccountModel struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Unverified bool `json:"unverified,omitempty"`
	Timezoneid string `json:"timeZoneId,omitempty"`
	Isowner bool `json:"isOwner,omitempty"`
	Planend string `json:"planEnd,omitempty"`
	Rolename string `json:"roleName,omitempty"`
	Featureflags string `json:"featureFlags,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Planstatus string `json:"planStatus,omitempty"`
	Allowcustombuildenvironment bool `json:"allowCustomBuildEnvironment,omitempty"`
	Blocked bool `json:"blocked,omitempty"`
	Planstart string `json:"planStart,omitempty"`
	Unpaid bool `json:"unpaid,omitempty"`
	Githubplan bool `json:"gitHubPlan,omitempty"`
	Manualpayments bool `json:"manualPayments,omitempty"`
	Roleid int `json:"roleId,omitempty"`
	Githubplanorg string `json:"gitHubPlanOrg,omitempty"`
	Iscollaborator bool `json:"isCollaborator,omitempty"`
	Planid string `json:"planId,omitempty"`
	Isenterpriseplan bool `json:"isEnterprisePlan,omitempty"`
	Name string `json:"name,omitempty"`
	Accountid int `json:"accountId,omitempty"`
}

// BuildJob represents the BuildJob schema from the OpenAPI specification
type BuildJob struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Name string `json:"name,omitempty"`
	Started string `json:"started,omitempty"`
	Status string `json:"status,omitempty"`
	Finished string `json:"finished,omitempty"`
	Jobid string `json:"jobId,omitempty"`
	Compilationwarningscount int `json:"compilationWarningsCount,omitempty"`
	Passedtestscount int `json:"passedTestsCount,omitempty"`
	Allowfailure bool `json:"allowFailure,omitempty"`
	Artifactscount int `json:"artifactsCount,omitempty"`
	Compilationerrorscount int `json:"compilationErrorsCount,omitempty"`
	Testscount int `json:"testsCount,omitempty"`
	Messagescount int `json:"messagesCount,omitempty"`
	Ostype string `json:"osType,omitempty"`
	Compilationmessagescount int `json:"compilationMessagesCount,omitempty"`
	Failedtestscount int `json:"failedTestsCount,omitempty"`
}

// DeploymentEnvironmentSettings represents the DeploymentEnvironmentSettings schema from the OpenAPI specification
type DeploymentEnvironmentSettings struct {
	Environmentvariables []StoredNameValue `json:"environmentVariables,omitempty"`
	Notifications []NotificationProviderSettings `json:"notifications,omitempty"`
	Providersettings []StoredNameValue `json:"providerSettings,omitempty"`
}

// DeploymentLookupModel represents the DeploymentLookupModel schema from the OpenAPI specification
type DeploymentLookupModel struct {
	Started string `json:"started,omitempty"`
	Status string `json:"status,omitempty"`
	Build BuildLookupModel `json:"build,omitempty"`
	Deploymentid int `json:"deploymentId,omitempty"`
	Finished string `json:"finished,omitempty"`
}

// ProjectDeploymentModel represents the ProjectDeploymentModel schema from the OpenAPI specification
type ProjectDeploymentModel struct {
	Started string `json:"started,omitempty"`
	Status string `json:"status,omitempty"`
	Build BuildLookupModel `json:"build,omitempty"`
	Deploymentid int `json:"deploymentId,omitempty"`
	Finished string `json:"finished,omitempty"`
	Environment DeploymentEnvironmentLookupModel `json:"environment,omitempty"`
}

// DeploymentCancellation represents the DeploymentCancellation schema from the OpenAPI specification
type DeploymentCancellation struct {
	Deploymentid int `json:"deploymentId"`
}

// RoleAddition represents the RoleAddition schema from the OpenAPI specification
type RoleAddition struct {
	Name string `json:"name"`
}

// ProjectBuildResults represents the ProjectBuildResults schema from the OpenAPI specification
type ProjectBuildResults struct {
	Build Build `json:"build,omitempty"`
	Project Project `json:"project,omitempty"`
}

// DeploymentEnvironmentLookupModel represents the DeploymentEnvironmentLookupModel schema from the OpenAPI specification
type DeploymentEnvironmentLookupModel struct {
	Deploymentenvironmentid int `json:"deploymentEnvironmentId,omitempty"`
	Name string `json:"name,omitempty"`
	Provider string `json:"provider,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Projectid int `json:"projectId"`
	Slug string `json:"slug,omitempty"`
	Accountname string `json:"accountName,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Enablesecurevariablesinpullrequestsfromsamerepo bool `json:"enableSecureVariablesInPullRequestsFromSameRepo,omitempty"`
	Isprivate bool `json:"isPrivate,omitempty"`
	Repositorybranch string `json:"repositoryBranch,omitempty"` // Not present in response from addProject.
	Alwaysbuildclosedpullrequests bool `json:"alwaysBuildClosedPullRequests,omitempty"`
	Skipbrancheswithoutappveyoryml bool `json:"skipBranchesWithoutAppveyorYml,omitempty"`
	Disablepushwebhooks bool `json:"disablePushWebhooks,omitempty"`
	Repositoryname string `json:"repositoryName"`
	Securitydescriptor SecurityDescriptor `json:"securityDescriptor,omitempty"`
	Rollingbuildsdonotcancelrunningbuilds bool `json:"rollingBuildsDoNotCancelRunningBuilds,omitempty"`
	Savebuildcacheinpullrequests bool `json:"saveBuildCacheInPullRequests,omitempty"`
	Enablesecurevariablesinpullrequests bool `json:"enableSecureVariablesInPullRequests,omitempty"`
	Isgithubapp bool `json:"isGitHubApp,omitempty"`
	Builds []Build `json:"builds,omitempty"` // Only non-empty for response from getProjects.
	Enabledeploymentinpullrequests bool `json:"enableDeploymentInPullRequests,omitempty"`
	Nugetfeed NuGetFeed `json:"nuGetFeed,omitempty"`
	Repositoryscm string `json:"repositoryScm,omitempty"`
	Accountid int `json:"accountId,omitempty"`
	Currentbuildid int `json:"currentBuildId,omitempty"`
	Rollingbuilds bool `json:"rollingBuilds,omitempty"`
	Rollingbuildsonlyforpullrequests bool `json:"rollingBuildsOnlyForPullRequests,omitempty"`
	Repositorytype string `json:"repositoryType,omitempty"`
	Tags string `json:"tags,omitempty"` // Comma-separated list of project tags for dynamic grouping. Appears that any input is accepted. The returned value only contains case-preserving but insensitive unique values where spaces around "," are removed but otherwise preserved. Empty values and items are allowed.
	Disablepullrequestwebhooks bool `json:"disablePullRequestWebhooks,omitempty"`
}

// NotificationSettings represents the NotificationSettings schema from the OpenAPI specification
type NotificationSettings struct {
	Account string `json:"account,omitempty"`
	Headersvalue string `json:"headersValue,omitempty"`
	Onbuildfailure bool `json:"onBuildFailure,omitempty"`
	Subjecttemplate string `json:"subjectTemplate,omitempty"`
	From string `json:"from,omitempty"`
	Template string `json:"template,omitempty"`
	Vsoaccount string `json:"vsoAccount,omitempty"`
	Incomingwebhookurl string `json:"incomingWebhookUrl,omitempty"`
	Onbuildstatuschanged bool `json:"onBuildStatusChanged,omitempty"`
	Onbuildsuccess bool `json:"onBuildSuccess,omitempty"`
	Password StoredValue `json:"password,omitempty"`
	Serverurl string `json:"serverUrl,omitempty"`
	Channel string `json:"channel,omitempty"`
	Url string `json:"url,omitempty"`
	Method string `json:"method,omitempty"`
	Bodytemplate string `json:"bodyTemplate,omitempty"`
	Customrequestbodycontenttype string `json:"customRequestBodyContentType,omitempty"`
	Recipients []StringValueObject `json:"recipients,omitempty"`
	Customrequestbody string `json:"customRequestBody,omitempty"`
	Recipientsvalue string `json:"recipientsValue,omitempty"`
	TypeField string `json:"$type,omitempty"`
	Addcustomrequestbody bool `json:"addCustomRequestBody,omitempty"`
	Authtoken StoredValue `json:"authToken,omitempty"`
	Room string `json:"room,omitempty"`
	Username string `json:"username,omitempty"`
	Headers []StoredNameValue `json:"headers,omitempty"`
}

// Script represents the Script schema from the OpenAPI specification
type Script struct {
	Language string `json:"language,omitempty"`
	Script string `json:"script"`
}

// PermissionState represents the PermissionState schema from the OpenAPI specification
type PermissionState struct {
	Allowed bool `json:"allowed,omitempty"` // State of the named permission. `true` to allow, `false` to deny, missing to inherit.
	Description string `json:"description,omitempty"`
	Name string `json:"name"` // Available permission names. The names correspond to the following groups: #### AccountPermission - ManageApplicationAuthorizations - UpdateAccountDetails - UpdateBillingDetails #### BuildEnvironmentPermission - ConfigureBuildEnvironment #### DenyPermission - DenyAllProjectsEnvironments #### EnvironmentsPermission - DeployToEnvironment - ManageEnvironments - UpdateEnvironmentSettings #### ProjectsPermission - ManageProjects - RunProjectBuild - UpdateProjectSettings #### RolesPermission - AddRole - DeleteRole - UpdateRoleDetails #### UserPermission - ConfigureApiKeys #### UsersPermission - AddUser - DeleteUser - UpdateUserDetails
}

// DeploymentEnvironment represents the DeploymentEnvironment schema from the OpenAPI specification
type DeploymentEnvironment struct {
	Deploymentenvironmentid int `json:"deploymentEnvironmentId,omitempty"`
	Name string `json:"name,omitempty"`
	Provider string `json:"provider,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Tags string `json:"tags,omitempty"` // Comma-separated list of environment tags for dynamic grouping. Appears that any input is accepted. The returned value only contains case-preserving but insensitive unique values where spaces around "," are removed but otherwise preserved. Empty values and items are allowed.
	Accountid int `json:"accountId,omitempty"`
	Projectsmode int `json:"projectsMode,omitempty"` // Project selection mode for deployment environments.
	Securitydescriptor SecurityDescriptor `json:"securityDescriptor,omitempty"`
}

// BuildLookupModel represents the BuildLookupModel schema from the OpenAPI specification
type BuildLookupModel struct {
	Version string `json:"version,omitempty"`
	Branch string `json:"branch,omitempty"`
	Buildid int `json:"buildId,omitempty"`
	Message string `json:"message,omitempty"`
}

// GroupPermissions represents the GroupPermissions schema from the OpenAPI specification
type GroupPermissions struct {
	Name string `json:"name"`
	Permissions []PermissionState `json:"permissions"`
}

// ProjectLookupModel represents the ProjectLookupModel schema from the OpenAPI specification
type ProjectLookupModel struct {
	Projectid int `json:"projectId"`
	Slug string `json:"slug,omitempty"`
	Accountname string `json:"accountName,omitempty"`
	Name string `json:"name,omitempty"`
}

// UserAccount represents the UserAccount schema from the OpenAPI specification
type UserAccount struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Isowner bool `json:"isOwner,omitempty"`
	Password string `json:"password,omitempty"`
	Roleid int `json:"roleId,omitempty"`
	Accountid int `json:"accountId,omitempty"`
	Twofactorauthenabled bool `json:"twoFactorAuthEnabled,omitempty"`
	Userid int `json:"userId,omitempty"`
	Accountname string `json:"accountName,omitempty"`
	Fullname string `json:"fullName"`
	Iscollaborator bool `json:"isCollaborator,omitempty"`
	Pagesize int `json:"pageSize,omitempty"`
	Rolename string `json:"roleName,omitempty"`
	Email string `json:"email"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Message string `json:"message"`
	Modelstate map[string]interface{} `json:"modelState,omitempty"` // When present, this property is a map of property names in the format `request.<capitalized name>` to an `Array` of validation error message strings for the property.
}

// ProjectDeploymentsResults represents the ProjectDeploymentsResults schema from the OpenAPI specification
type ProjectDeploymentsResults struct {
	Project Project `json:"project,omitempty"`
	Deployments []ProjectDeploymentModel `json:"deployments,omitempty"`
}

// DeploymentEnvironmentSettingsResults represents the DeploymentEnvironmentSettingsResults schema from the OpenAPI specification
type DeploymentEnvironmentSettingsResults struct {
	Environment DeploymentEnvironmentWithSettings `json:"environment,omitempty"`
}

// Timestamped represents the Timestamped schema from the OpenAPI specification
type Timestamped struct {
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
}

// ReRunBuildRequest represents the ReRunBuildRequest schema from the OpenAPI specification
type ReRunBuildRequest struct {
	Buildid int `json:"buildId"`
	Rerunincomplete bool `json:"reRunIncomplete,omitempty"` // Set `reRunIncomplete` set to `false` (default value) for full build re-run. Set it set to `true` to rerun only failed or cancelled jobs in multijob build.
}

// ProjectConfiguration represents the ProjectConfiguration schema from the OpenAPI specification
type ProjectConfiguration struct {
	Branchesmode string `json:"branchesMode,omitempty"`
	Skipnontags bool `json:"skipNonTags,omitempty"`
	Dotnetcsprojfile string `json:"dotnetCsprojFile,omitempty"`
	Msbuildprojectfilename string `json:"msBuildProjectFileName,omitempty"`
	Dotnetcsprojpackageversionformat string `json:"dotnetCsprojPackageVersionFormat,omitempty"`
	Configurenugetaccountsource bool `json:"configureNuGetAccountSource,omitempty"`
	Testcategoriesmode string `json:"testCategoriesMode,omitempty"`
	Assemblyinformationalversionformat string `json:"assemblyInformationalVersionFormat,omitempty"`
	Initscripts []Script `json:"initScripts,omitempty"`
	Dotnetcsprojinformationalversionformat string `json:"dotnetCsprojInformationalVersionFormat,omitempty"`
	Disablenugetpublishforoctopuspackages bool `json:"disableNuGetPublishForOctopusPackages,omitempty"`
	Matrixonly []StoredNameValueMatrix `json:"matrixOnly,omitempty"`
	Packagewebapplicationprojects bool `json:"packageWebApplicationProjects,omitempty"`
	Aftertestscripts []Script `json:"afterTestScripts,omitempty"`
	Skipcommitsfiles []StringValueObject `json:"skipCommitsFiles,omitempty"`
	Onbuilderrorscripts []Script `json:"onBuildErrorScripts,omitempty"`
	Matrixfastfinish bool `json:"matrixFastFinish,omitempty"`
	Assemblyfileversionformat string `json:"assemblyFileVersionFormat,omitempty"`
	Includebranches []StringValueObject `json:"includeBranches,omitempty"`
	Cacheentries []StringValueObject `json:"cacheEntries,omitempty"`
	Forcehttpsclone bool `json:"forceHttpsClone,omitempty"`
	Installscripts []Script `json:"installScripts,omitempty"`
	Operatingsystem []map[string]interface{} `json:"operatingSystem,omitempty"`
	Donotincrementbuildnumberonpullrequests bool `json:"doNotIncrementBuildNumberOnPullRequests,omitempty"`
	Services []map[string]interface{} `json:"services,omitempty"`
	Deployments []DeploymentProvider `json:"deployments,omitempty"`
	Packageaspnetcoreprojects bool `json:"packageAspNetCoreProjects,omitempty"`
	Beforebuildscripts []Script `json:"beforeBuildScripts,omitempty"`
	Buildcloud []StringValueObject `json:"buildCloud,omitempty"`
	Patchassemblyinfo bool `json:"patchAssemblyInfo,omitempty"`
	Patchdotnetcsproj bool `json:"patchDotnetCsproj,omitempty"`
	Testscripts []Script `json:"testScripts,omitempty"` // Only set/used when `testMode` is `script`.
	Environmentvariables []StoredNameValue `json:"environmentVariables,omitempty"`
	Onbuildsuccessscripts []Script `json:"onBuildSuccessScripts,omitempty"`
	Afterbuildscripts []Script `json:"afterBuildScripts,omitempty"`
	Notifications []NotificationProviderSettings `json:"notifications,omitempty"`
	Clonescripts []Script `json:"cloneScripts,omitempty"`
	Testcategoriesmatrix []map[string]interface{} `json:"testCategoriesMatrix,omitempty"`
	Packagenugetprojects bool `json:"packageNuGetProjects,omitempty"`
	Matrixallowfailures []StoredNameValueMatrix `json:"matrixAllowFailures,omitempty"` // Although the names and values are not enforced, the combinations which are meaningful are documented at https://www.appveyor.com/docs/build-configuration/#allow-failing-jobs
	Excludebranches []StringValueObject `json:"excludeBranches,omitempty"`
	Maxjobs int `json:"maxJobs,omitempty"`
	Deployscripts []Script `json:"deployScripts,omitempty"`
	Onlycommitsfiles []StringValueObject `json:"onlyCommitsFiles,omitempty"`
	Packagedotnetconsoleprojects bool `json:"packageDotnetConsoleProjects,omitempty"`
	Matrixexclude []StoredNameValueMatrix `json:"matrixExclude,omitempty"`
	Packagewebapplicationprojectsoctopus bool `json:"packageWebApplicationProjectsOctopus,omitempty"`
	Skiptags bool `json:"skipTags,omitempty"`
	Afterdeployscripts []Script `json:"afterDeployScripts,omitempty"`
	Beforedeployscripts []Script `json:"beforeDeployScripts,omitempty"`
	Buildmode string `json:"buildMode,omitempty"`
	Onbuildfinishscripts []Script `json:"onBuildFinishScripts,omitempty"`
	Matrixexcept []StoredNameValueMatrix `json:"matrixExcept,omitempty"`
	Packageazurecloudserviceprojects bool `json:"packageAzureCloudServiceProjects,omitempty"`
	Assemblyinfofile string `json:"assemblyInfoFile,omitempty"`
	Xamarinregisteriosproduct bool `json:"xamarinRegisterIosProduct,omitempty"`
	Skipbranchwithpullrequests bool `json:"skipBranchWithPullRequests,omitempty"`
	Configuration []StringValueObject `json:"configuration,omitempty"`
	Dotnetcsprojversionformat string `json:"dotnetCsprojVersionFormat,omitempty"`
	Packagewebapplicationprojectsbeanstalk bool `json:"packageWebApplicationProjectsBeanstalk,omitempty"`
	Hostsentries []HostEntry `json:"hostsEntries,omitempty"`
	Packagenugetsymbols bool `json:"packageNuGetSymbols,omitempty"`
	Hotfixscripts []Script `json:"hotFixScripts,omitempty"`
	Testcategories []StringValueObject `json:"testCategories,omitempty"`
	Clonefolder string `json:"cloneFolder,omitempty"`
	Environmentvariablesmatrix []StoredNameValueMatrix `json:"environmentVariablesMatrix,omitempty"`
	Disablenugetpublishonpullrequests bool `json:"disableNuGetPublishOnPullRequests,omitempty"`
	Artifacts []Artifact `json:"artifacts,omitempty"`
	Clonedepth int `json:"cloneDepth,omitempty"`
	Xamarinregisterandroidproduct bool `json:"xamarinRegisterAndroidProduct,omitempty"`
	Testmode string `json:"testMode,omitempty"`
	Msbuildverbosity string `json:"msBuildVerbosity,omitempty"`
	Testassemblies []StringValueObject `json:"testAssemblies,omitempty"`
	Dotnetcsprojfileversionformat string `json:"dotnetCsprojFileVersionFormat,omitempty"`
	Buildscripts []Script `json:"buildScripts,omitempty"` // Only set/used when `buildMode` is `script`.
	Assemblyversionformat string `json:"assemblyVersionFormat,omitempty"`
	Beforetestscripts []Script `json:"beforeTestScripts,omitempty"`
	Packagewebapplicationprojectsxcopy bool `json:"packageWebApplicationProjectsXCopy,omitempty"`
	Includenugetreferences bool `json:"includeNuGetReferences,omitempty"`
	Deploymode string `json:"deployMode,omitempty"`
	Platform []map[string]interface{} `json:"platform,omitempty"`
	Dotnetcsprojassemblyversionformat string `json:"dotnetCsprojAssemblyVersionFormat,omitempty"`
	Msbuildinparallel bool `json:"msBuildInParallel,omitempty"`
	Shallowclone bool `json:"shallowClone,omitempty"`
	Configurenugetprojectsource bool `json:"configureNuGetProjectSource,omitempty"`
	Stacks []string `json:"stacks,omitempty"`
	Beforepackagescripts []Script `json:"beforePackageScripts,omitempty"`
}

// DeploymentStartRequest represents the DeploymentStartRequest schema from the OpenAPI specification
type DeploymentStartRequest struct {
	Accountname string `json:"accountName"`
	Buildjobid string `json:"buildJobId,omitempty"` // Optional job id with artifacts if build contains multiple jobs.
	Buildversion string `json:"buildVersion"` // Build to deploy
	Environmentname string `json:"environmentName"`
	Environmentvariables map[string]interface{} `json:"environmentVariables,omitempty"`
	Projectslug string `json:"projectSlug"`
}

// UserAccountRolesResults represents the UserAccountRolesResults schema from the OpenAPI specification
type UserAccountRolesResults struct {
	User UserAccount `json:"user,omitempty"`
	Roles []Role `json:"roles,omitempty"`
}

// NotificationProviderSettings represents the NotificationProviderSettings schema from the OpenAPI specification
type NotificationProviderSettings struct {
	Provider string `json:"provider"`
	Settings NotificationSettings `json:"settings"` // This type is the union of the settings types for each of the various notification types supported by the API. The properties correspond to the following notification types: #### All Types - onBuildSuccess - onBuildFailure - onBuildStatusChanged #### Campfire - account - authToken - room - template #### Email - subjectTemplate - bodyTemplate - recipients - recipientsValue #### GitHubPullRequest - authToken - template #### HipChat - authToken - from - room - template - serverUrl #### Slack - incomingWebhookUrl - authToken - channel - template #### Webhook - method - url - headers - headersValue - addCustomRequestBody - customRequestBodyContentType - customRequestBody #### VSOTeamRoom - vsoAccount - username - password - room - template
}

// DeploymentProvider represents the DeploymentProvider schema from the OpenAPI specification
type DeploymentProvider struct {
	Onbranch []StringValueObject `json:"onBranch,omitempty"`
	Onenvironmentvariables []StoredNameValue `json:"onEnvironmentVariables,omitempty"`
	Provider string `json:"provider"`
	Providersettings []StoredNameValue `json:"providerSettings,omitempty"`
}

// HostEntry represents the HostEntry schema from the OpenAPI specification
type HostEntry struct {
	Host string `json:"host"`
	Ip string `json:"ip"`
}
