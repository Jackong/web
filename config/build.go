/**
 * User: Jackong
 * Date: 13-8-17
 * Time: 上午9:27
 */
package config

import (
	"encoding/xml"
	"os"
	"fmt"
)

type build struct {
	Init string `xml:"init"`
	Ctrl string `xml:"ctrl"`
}

var Build build

func init() {
	data := read("config/build.xml")
	if err := xml.Unmarshal(data, &Build); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
