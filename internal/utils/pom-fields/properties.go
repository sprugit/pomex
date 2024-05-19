package pomfields

import (
	"encoding/xml"
)

// ------------------------------------
/* Deprecated, mapped properly now.
type Property struct {
	Key   string
	Value string
}

func (p *Property) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.Key = start.Name.Local
	return d.DecodeElement(&p.Value, &start)
}

type PropertyWrapper struct {
	XMLName    xml.Name   `xml:"properties"`
	Properties []Property `xml:",any"`
}
*/
// ------------------------------------

type Properties map[string]string

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var Key, Value string = start.Name.Local, ""
	if err := d.DecodeElement(&Value, &start); err != nil {
		return err
	}

	//fmt.Println("inside unmarshal values (key-val):", Key, Value)

	(*p)[Key] = Value

	return nil
}

type PropertiesWrapper struct {
	XMLName xml.Name   `xml:"properties"`
	List    Properties `xml:",any"`
}

/*
type PProject struct {
	XMLName    xml.Name          `xml:"project"`
	Properties PropertiesWrapper `xml:"properties"`
}


func main() {

	const target string = "E:/codingProjects/pomex-go/test/test_data/dummy_project1/project-parent/pom.xml"

	data, err := os.ReadFile(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	proj := PProject{Properties: PropertiesWrapper{List: make(Properties)}}
	if err := xml.Unmarshal(data, &proj); err != nil {
		panic(err)
	}
	fmt.Println(proj)
	for key, value := range proj.Properties.List {
		fmt.Println("KEY:", key, "VALUE:", value)
	}
}*/
