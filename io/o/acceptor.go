/**
 * User: Jackong
 * Date: 13-7-27
 * Time: 上午10:46
 */
package o

type IAcceptor interface {
	Present(*Output) error
}

var (
	acceptors map[string] IAcceptor
)

func init() {
	//unexported: match pattern by Get
	acceptors = make(map[string] IAcceptor)
}

func Register(accept string, acceptor IAcceptor) {
	acceptors[accept] = acceptor
}

func Acceptor(accept string) (IAcceptor) {
	return acceptors[accept]
}
