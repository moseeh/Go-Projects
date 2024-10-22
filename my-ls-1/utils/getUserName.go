package utils

import (
	"os/user"
	"strconv"
)

// getUserName retrieves the username associated with the given user ID (uid).
// If the user ID cannot be found, it returns the user ID as a string.
func getUserName(uid int) string {
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		return strconv.Itoa(uid)
	}
	return u.Username
}
