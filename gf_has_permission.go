package helpers

// check the permissions match with the access keys
func HasPermission(permissions, accessKeys []string) bool {
	hasPermission := false
	for _, item := range accessKeys {
		if contains(permissions, item) {
			hasPermission = true
			break
		}
	}
	return hasPermission
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
