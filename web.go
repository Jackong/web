/**
 * User: Jackong
 * Date: 13-7-14
 * Time: 下午6:21
 */
package web
import (
	"github.com/Jackong/web/config"
	"net/http"
	"github.com/Jackong/web/mapper"
	"github.com/Jackong/web/common/log"
	_ "github.com/Jackong/web/static"
	_ "github.com/Jackong/web/io/o/json"
	_ "github.com/Jackong/web/io/o/tpl"
	"github.com/Jackong/web/io/i"
	"github.com/Jackong/web/io/o"
	"github.com/Jackong/web/io"
	"github.com/Jackong/web/method"
)

func Go() {
	defer log.Close()
	log.Info("collecting controller")
	log.Info("setting handler for home")
	http.HandleFunc("/", HomeHandler)
	log.Info("listening")
	http.ListenAndServe(config.Project.Server.Addr, nil)
}

func HomeHandler(writer http.ResponseWriter, req * http.Request) {
	log.Info(req.RemoteAddr, req.Method, req.URL.Path)
	acceptor := o.Acceptor(req.Header.Get("Accept"))
	if acceptor == nil {
		http.Error(writer, "406 Not Acceptable", http.StatusNotAcceptable)
		return
	}
	ctrl := mapper.Get(req.URL.Path)
	if ctrl == nil {
		log.Warning("page not found")
		http.NotFound(writer, req);
		return;
	}
	input := i.New(req)
	if ok, tips := ctrl.Check(input); !ok {
		log.Error(tips)
		http.Error(writer, tips, http.StatusNotFound)
		return
	}
	ctx := io.New(input, &o.Output{Writer: writer})
	switch req.Method {
	case method.GET:
		ctrl.Read(ctx)
	case method.POST:
		ctrl.Create(ctx)
	case method.PUT:
		ctrl.Update(ctx)
	case method.DEL:
		ctrl.Delete(ctx)
	default:
		http.Error(ctx.Writer, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if _, output := ctx.Output.Get(); output == nil {
		http.Error(ctx.Writer, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := acceptor.Present(ctx.Output); err != nil {
		log.Error(ctx.Output, "present error:", err)
		http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
