package enums

// UserRole ...
type UserRole int

const (
	User_Unknown UserRole = iota
	SuperAdmin
	Admin
	Customer
	User
)

// Sex ...
type UserSex int

const (
	Sex_Unknown UserSex = iota
	Female
	Male
)
