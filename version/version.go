package version

import "fmt"

var (
	version   = "unknown"
	commitID  = "unknown"
	buildDate = "unknown"
	buildTags = "unknown"
)

func VersionInfo() string {
	return fmt.Sprintf("paopao %s (build:%s %s)", version, commitID, buildDate)
}
