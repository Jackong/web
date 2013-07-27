/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午9:03
 */
package tpl

import (
	"github.com/Jackong/web/io/o"
	"path/filepath"
	"html/template"
	"github.com/Jackong/web/config"
	"os"
	"github.com/Jackong/web/common/log"
)

var (
	tplFiles []string
	tpls *template.Template
)

func init() {
	log.Info("loading templates")
	filepath.Walk(config.Project.Dir.Tpl, visit)
	tpls = template.Must(template.ParseFiles(tplFiles ...))
	o.Register("text/html", new(Tpl))
}

func visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		tplFiles = append(tplFiles, path)
	}
	return err
}

type Tpl struct {}

func (this *Tpl) Present(output *o.Output) error {
	name, out := output.Get()
	if err := tpls.ExecuteTemplate(output.Writer, name + ".html", out); err != nil {
		return err
	}
	return nil
}


