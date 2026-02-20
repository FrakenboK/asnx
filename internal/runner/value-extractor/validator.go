package extractor

import "regexp"

var (
	ipRegex     = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.){3}(25[0-5]|(2[0-4]|1\d|[1-9]|)\d)$`)
	domainRegex = regexp.MustCompile(`^[A-Za-z0-9-]{1,63}\.[A-Za-z]{2,6}$`)
)
