/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午9:03
 */
package tpl

import (
	"os"
	"path/filepath"
	"html/template"

	"github.com/Jackong/web/io/o"
	"github.com/Jackong/web/common/log"
)

var (
	tplFiles []string
	tpls *template.Template
)

func Init(tplPath, suffix string) {
	log.Info("loading templates")
	filepath.Walk(tplPath, visit)
	tpls = template.Must(template.ParseFiles(tplFiles ...))
	o.Register("text/html", &Tpl{suffix: suffix})
}

func Close() {
	tpls.Clone()
}

func visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		tplFiles = append(tplFiles, path)
	}
	return err
}

type Tpl struct {
	suffix string
}

func (this *Tpl) Present(output *o.Output) error {
	name, out := output.Get()
	if err := tpls.ExecuteTemplate(output.Writer, name + this.suffix, out); err != nil {
		return err
	}
	return nil
}


