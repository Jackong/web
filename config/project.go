/**
 * User: Jackong
 * Date: 13-7-18
 * Time: 下午9:58
 */
package config

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
)

type project struct {
	Dir dir `xml:"dir"`
	Server server `xml:"server"`
}

type server struct {
	Name string `xml:"name"`
	Addr string `xml:"addr"`
}
type dir struct {
	Log string `xml:"log"`
	Ctrl string `xml:"ctrl"`
	Tpl tpl `xml:"tpl"`
	Static static `xml:"static"`
}

type tpl struct {
	Path string `xml:"path"`
	Suffix string `xml:"suffix"`
}
type static struct {
	Root string `xml:"root,attr"`
	Paths []string `xml:"path"`
}

var (
	Project project
)

func init() {
	file, err := os.Open("config/project.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return
	}

	Project = project{}
	err = xml.Unmarshal(data, &Project)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return
	}
}
