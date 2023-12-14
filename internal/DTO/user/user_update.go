package user

type UserUpdate struct {
	Id            int64
	NameValue     string
	NameProvided  bool
	EmailValue    string
	EmailProvided bool
	RoleValue     int
	RoleProvided  bool
}
