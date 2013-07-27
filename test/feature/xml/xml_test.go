/**
 * User: Jackong
 * Date: 13-7-18
 * Time: 下午8:51
 */
package xml

import (
	"testing"
	"encoding/xml"
	"os"
	"io/ioutil"
)

type Project struct {
	Self xml.Name `xml:"project"`
	Dir dir `xml:"dir"`
}

type dir struct {
	Self xml.Name `xml:"dir"`
	Log string `xml:"log"`
	Ctrl string `xml:"ctrl"`
	Tpl string `xml:"tpl"`
	Static static `xml:"static"`
}

type static struct {
	Self xml.Name `xml:"static"`
	Root string `xml:"root,attr"`
	Paths []string `xml:"path"`
}

func TestRead(t *testing.T) {
	file, err := os.Open("test.xml")
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}

	project := Project{}
	err = xml.Unmarshal(data, &project)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}

	expectNode(t, "log", project.Dir.Log)
	expectNode(t, "ctrl", project.Dir.Ctrl)
	expectNode(t, "view/tpl", project.Dir.Tpl)
	expectNode(t, "view/static", project.Dir.Static.Root)
	expectNode(t, "js", project.Dir.Static.Paths[0])
	expectNode(t, "image", project.Dir.Static.Paths[1])
	expectNode(t, "css", project.Dir.Static.Paths[2])
	expectNode(t, "html", project.Dir.Static.Paths[3])
}

func expectNode(t * testing.T, expect, actual interface {}) {
	if expect != actual {
		t.Error(expect)
	}
}
