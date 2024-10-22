package utils

import (
	"os/user"
	"strconv"
)

// getGroupName retrieves the group name associated with the given group ID (gid).
// If the group ID cannot be found, it returns the group ID as a string.
func getGroupName(gid int) string {
	g, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		return strconv.Itoa(gid)
	}
	return g.Name
}
