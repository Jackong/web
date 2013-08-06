/**
 * User: Jackong
 * Date: 13-7-14
 * Time: 下午6:21
 */
package web
import (
	"net/http"

	"github.com/Jackong/web/config"
	"github.com/Jackong/web/mapper"
	"github.com/Jackong/web/common/log"
	"github.com/Jackong/web/io/i"
	"github.com/Jackong/web/io/o"
	"github.com/Jackong/web/io"
	"github.com/Jackong/web/method"
	"github.com/Jackong/web/static"
	"github.com/Jackong/web/io/o/tpl"
	"github.com/Jackong/web/io/o/json"
)

func Go() {
	log.Init(config.Project.Dir.Log)
	defer log.Close()

	tpl.Init(config.Project.Dir.Tpl.Path, config.Project.Dir.Tpl.Suffix)
	defer tpl.Close()

	json.Init()

	static.Init(config.Project.Dir.Static.Root, config.Project.Dir.Static.Paths)

	log.Info("collecting controller")
	log.Info("setting handler for home")
	http.HandleFunc("/", HomeHandler)
	log.Info("listening")
	http.ListenAndServe(config.Project.Server.Addr, nil)
}

func HomeHandler(writer http.ResponseWriter, req * http.Request) {
	defer func() {
		if e := recover(); e != nil {
			err := e.(i.InputError)
			log.Error(err)
			http.NotFound(writer, req)
			return
		}
	}()

	log.Info(req.RemoteAddr, req.Method, req.URL.Path)
	accept := req.Header.Get("Accept")
	acceptor := o.Acceptor(accept)
	if acceptor == nil {
		log.Error("not acceptable", accept)
		http.Error(writer, "406 Not Acceptable", http.StatusNotAcceptable)
		return
	}

	ctrl := mapper.Get(req.URL.Path)
	if ctrl == nil {
		log.Warning("page not found", req.URL.Path)
		http.NotFound(writer, req);
		return;
	}

	ctx := io.New(&i.Input{Req: req}, &o.Output{Writer: writer})
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
		http.Error(writer, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if _, output := ctx.Output.Get(); output == nil {
		http.Error(writer, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := acceptor.Present(ctx.Output); err != nil {
		log.Error(ctx.Output, "present error", err)
		http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
