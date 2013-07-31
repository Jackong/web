/**
 * User: Jackong
 * Date: 13-7-27
 * Time: 上午10:46
 */
package o

import (
	"strings"
	"strconv"
)

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
	if acceptor, ok := acceptors[accept]; ok {
		return acceptor
	}

	for _, acps := range order(accept) {
		for _, acp := range acps {
			if acceptor, ok := acceptors[acp]; ok {
				return acceptor
			}
		}
	}
	return nil
}


func order(accept string) [] []string{
	order := make([] []string, 10)
	for _, acps := range strings.Split(accept, ",") {
		acp := strings.Split(acps, ";")
		q := 0
		if len(acp) != 1 {
			qf, err := strconv.ParseFloat(strings.TrimPrefix(acp[1], "q="), 64)
			if err != nil {
				return nil
			}
			q = 10 - int(qf * 10)
		}
		order[q] = append(order[q], acp[0])
	}
	return order
}
