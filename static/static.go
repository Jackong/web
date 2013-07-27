/**
 * User: Jackong
 * Date: 13-7-18
 * Time: 下午11:37
 */
package static

import (
	"github.com/Jackong/web/config"
	"github.com/Jackong/web/common/log"
	"net/http"
)

func init() {
	log.Info("setting static resource")
	for _, static := range config.Project.Dir.Static.Paths {
		http.Handle("/" + static + "/", http.FileServer(http.Dir(config.Project.Dir.Static.Root)))
	}
}
