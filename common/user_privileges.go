package common

// user/admin privileges
const (
	UserPrivilegePublic = 1 << iota
	UserPrivilegeNormal
	UserPrivilegeDonor
	AdminPrivilegeAccessRAP
	AdminPrivilegeManageUsers
	AdminPrivilegeBanUsers
	AdminPrivilegeSilenceUsers
	AdminPrivilegeWipeUsers
	AdminPrivilegeManageBeatmap
	AdminPrivilegeManageServer
	AdminPrivilegeManageSetting
	AdminPrivilegeManageBetaKey
	AdminPrivilegeManageReport
	AdminPrivilegeManageDocs
	AdminPrivilegeManageBadges
	AdminPrivilegeViewRAPLogs
	AdminPrivilegeManagePrivilege
	AdminPrivilegeSendAlerts
	AdminPrivilegeChatMod
	AdminPrivilegeKickUsers
)
