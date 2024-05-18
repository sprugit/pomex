package utils

type POM struct {
	groupId      string
	artifactId   string
	version      string
	children     []*POM
	dependencies []dependency
	depMgmt      map[string]string
	props        map[string]string
}
