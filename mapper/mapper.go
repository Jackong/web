/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午4:22
 */
package mapper

import (
	"githud.com/Jackong/web/ctrl"
)

var (
	//not use map and exported: match pattern by Get
	mapper map[string] ctrl.ICtrl
)

func init() {
	mapper = make(map[string] ctrl.ICtrl)
}

func Get(key string) ctrl.ICtrl {
	return mapper[key]
}

func Set(key string, value ctrl.ICtrl) {
	mapper[key] = value
}
