package history

type UserUpdateHistory struct {
	UserId   int64
	EmailOld string
	EmailNew string
	NameOld  string
	NameNew  string
	RoleOld  int
	RoleNew  int
}
