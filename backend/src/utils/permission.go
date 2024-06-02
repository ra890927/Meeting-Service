package utils

type Permission uint

const (
	Empty  Permission = 0
	Create Permission = 1 << iota
	Update
	Upload
	Delete
	Read
)

func CheckPermission(role, perm Permission) bool {
	return (role & perm) > 0
}
