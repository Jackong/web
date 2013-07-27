/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午8:32
 */
package json

import (
	"github.com/Jackong/web/io/o"
	"encoding/json"
	"fmt"
)

func init() {
	o.Register("application/json", new(Json))
}

type Json struct {}

func (this *Json) Present(output *o.Output) error {
	name, out := output.Get()
	b, err := json.Marshal(out)
	if err != nil {
		return err
	}
	str := string(b)
	if name != "" {
		str = `{"` + name + `":` + str + `}`
	}
	fmt.Fprint(output.Writer, str)
	return nil
}
