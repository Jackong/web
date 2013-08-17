/**
 * User: Jackong
 * Date: 13-7-19
 * Time: 上午8:30
 */
package main

import (
	"os"
	"path/filepath"
	"fmt"
	"strings"

	"github.com/Jackong/web/config"
)

func main() {
	content := `package main

import (
` + func() string {
		if len(config.Build.Init) == 0 {
			return ""
		}
		return `	_"` + config.Build.Init + `"`
	}() + `
	"github.com/Jackong/web"
` + ctrls() + `
)

func main() {
	web.Go()
}`
	os.Mkdir("src/main", os.ModeDir)
	generate("src/main/main.go", content)
}

func ctrls() (pkgs string){
	packages = make(map[string] bool)
	dir := "src/" + config.Build.Ctrl
	filepath.Walk(dir, visit)
	if len(packages) == 0 {
		fmt.Println("waring: no ctrl packages")
	}
	for pkg, _ := range packages {
		fmt.Println(pkg)
		pkgs += `	_ "` + pkg + `"` + "\n"
	}
	return pkgs
}

func generate(file, content string) {
	os.Remove(file)
	ctrlFile, err := os.OpenFile(file, os.O_RDWR | os.O_CREATE, 0)
	if err != nil {
		fmt.Printf("warning: %v\n", err)
	}
	defer ctrlFile.Close()
	ctrlFile.WriteString(content)
	fmt.Printf("generate ctrls file: %s", file)
}
var (
	packages map[string] bool
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
