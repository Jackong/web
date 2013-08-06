/**
 * User: Jackong
 * Date: 13-7-14
 * Time: 下午2:57
 */
package ctrl

import (
	"github.com/Jackong/web/io"
)

type ICtrl interface {
	Create(*io.Context)
	Read(*io.Context)
	Update(*io.Context)
	Delete(*io.Context)
}

type Ctrl struct {
}

func (this *Ctrl) Create(*io.Context) {
}

func (this *Ctrl) Read(*io.Context) {
}

func (this *Ctrl) Update(*io.Context) {
}

func (this *Ctrl) Delete(*io.Context) {
}
