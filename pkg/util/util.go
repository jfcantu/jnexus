package util

import "strings"

// GetServerName returns the hostname component of an FQDN
func GetServerName(hostname string) string {
	components := strings.Split(hostname, ".")
	return components[0]
}
