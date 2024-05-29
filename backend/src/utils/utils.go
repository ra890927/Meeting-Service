package utils

func CheckPermission(role, perm interface{}) bool {
	return (role.(int) & perm.(int)) > 0
}
