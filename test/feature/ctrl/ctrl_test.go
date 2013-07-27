/**
 * User: Jackong
 * Date: 13-7-19
 * Time: 上午8:38
 */
package ctrl

import (
	"os"
	"path/filepath"
	"testing"
	"strings"
)

func init() {
	filepath.Walk("ok", visit)
}
var (
	packages string
)

func visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		packages += "\n\t" + `_ "test/feature/ctrl/` + strings.TrimSuffix(path, "\\" + info.Name()) + `"`
		packages = strings.Replace(packages, "\\", "/", -1)
	}
	return err
}

func TestPackages(t *testing.T) {
	if packages == "" {
		t.Fail()
	}
	content := "package dc\nimport (" + packages + "\n)"
	t.Log(content)
	os.Remove("dc/dc.go")
	logfile, err := os.OpenFile("dc/dc.go", os.O_RDWR | os.O_CREATE, 0)
	if err != nil {
		t.Errorf("warning: %v\n", err)
	}
	defer logfile.Close()
	logfile.WriteString(content)
}
