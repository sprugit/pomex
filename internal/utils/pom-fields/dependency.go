package pomfields

import (
	"encoding/xml"
	"strings"
)

// ------------------------------------

type Dependency struct {
	XMLName    xml.Name `xml:"dependency"`
	GroupId    string   `xml:"groupId"`
	ArtifactId string   `xml:"artifactId"`
	Version    string   `xml:"version"`
}

func (d Dependency) getNameSpace() string {
	var sb strings.Builder
	sb.WriteString(d.GroupId)
	sb.WriteString(".")
	sb.WriteString(d.ArtifactId)
	return sb.String()
}

// ------------------------------------

type Dependencies struct {
	XMLName xml.Name      `xml:"dependencies"`
	List    []*Dependency `xml:"dependency"`
}

// ------------------------------------

type DependencyManagement struct {
	Deps Dependencies `xml:"dependencies"`
}

// ------------------------------------

type Modules 

// ------------------------------------

type Project struct {
	XMLName      xml.Name          `xml:"project"`
	Dependencies Dependencies      `xml:"dependencies"`
	Properties   PropertiesWrapper `xml:"properties"`
}

/*
func main() {

	const target string = "E:/codingProjects/pomex-go/test/test_data/dummy_project1/project-parent/pom.xml"

	data, err := os.ReadFile(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var pro Project
	if err := xml.Unmarshal(data, &pro); err != nil {
		panic(err)
	}
	for _, dep := range pro.Dependencies.List {
		fmt.Println(dep.getNameSpace())
	}
	fmt.Print(pro)
}
*/
