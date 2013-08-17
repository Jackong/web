/**
 * User: Jackong
 * Date: 13-7-18
 * Time: 下午9:58
 */
package config

import (
	"encoding/xml"
	"os"
	"fmt"
)

type project struct {
	Dir    dir `xml:"dir"`
	Server server `xml:"server"`
}

type server struct {
	Name string `xml:"name"`
	Addr string `xml:"addr"`
}
type dir struct {
	Log    string `xml:"log"`
	Tpl    tpl `xml:"tpl"`
	Static static `xml:"static"`
}

type tpl struct {
	Path   string `xml:"path"`
	Suffix string `xml:"suffix"`
}
type static struct {
	Root  string `xml:"root,attr"`
	Paths []string `xml:"path"`
}

var (
	Project project
)

func init() {
	data := read("config/project.xml")
	Project = project{}
	if err := xml.Unmarshal(data, &Project); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return
	}
}
