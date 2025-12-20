package auth

import "strings"

//stringContainsFold performs case-insensitive substring check
func stringContainsFold(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
