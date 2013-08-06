/**
 * User: Jackong
 * Date: 13-7-18
 * Time: 下午11:37
 */
package static

import (
	"net/http"

	"github.com/Jackong/web/common/log"
)

func Init(root string, paths []string) {
	log.Info("setting static resource")
	for _, static := range  paths {
		http.Handle("/" + static + "/", http.FileServer(http.Dir(root)))
	}
}
