package enums

type ResCodeCommon int
type ResCodeOfUser int
type ResCodeOfInterface int

const (
	// ResCodeCommon
	Success            ResCodeCommon = iota // 0
	ParameterError                          // 1
	AuthenticationFail                      // 2
	Unauthorized
	NotAdminRole
	GenerateRandomKeyFailed
	GenerateTokenFailed
)

const (
	// ResCodeOfUser
	UserNotExist     ResCodeOfUser = 2001
	UserExist                      // 2002
	CreateUserFailed               // 2003
	UserPasswordError
)

const (
	// ResCodeOfInterface
	InterfaceNotExist ResCodeOfInterface = 3000
	ListInterfaceFailed
	CreateInterfaceFailed
	UpdateInterfaceFailed
	DeleteInterfaceFailed
	OnlineInterfaceFailed
	OfflineInterfaceFailed
)
