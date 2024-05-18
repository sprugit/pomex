package utils

import (
	"strings"
)

type dependency struct {
	groupId    string
	artifactId string
	version    string
}

func (d dependency) getNameSpace() string {
	var sb strings.Builder
	sb.WriteString(d.groupId)
	sb.WriteString(".")
	sb.WriteString(d.artifactId)
	return sb.String()
}
