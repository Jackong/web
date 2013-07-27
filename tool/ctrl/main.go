/**
 * User: Jackong
 * Date: 13-7-19
 * Time: 上午8:30
 */
package main

import (
	"github.com/Jackong/web/config"
	"os"
	"path/filepath"
	"fmt"
	"strings"
)

func main() {
	packages = make(map[string]bool)
	filepath.Walk(config.Project.Dir.Ctrl, visit)
	if len(packages) == 0 {
		fmt.Println("waring: no ctrl packages")
	}
	fmt.Println("importing ctrl packages:")
	content := "package web\nimport ("
	for key, _ := range packages {
		fmt.Println(key)
		content += "\n\t" + `_ "` + key + `"`
	}
	content += "\n)"
	ctrls := config.Project.Dir.Ctrl + "/_ctrls.go"
	os.Remove(ctrls)
	ctrlFile, err := os.OpenFile(ctrls, os.O_RDWR | os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("warning: %v\n", err)
	}
	defer ctrlFile.Close()
	ctrlFile.WriteString(content)
	fmt.Printf("generate ctrls file: %s", ctrls)
}

var (
	packages map[string]bool
)

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		ctrl := strings.Replace(strings.TrimSuffix(strings.TrimPrefix(path, "src\\"), "\\" + info.Name()), "\\", "/", -1)
		packages[ctrl] = true
	}
	return err
}
