package server

import (
	"net/http"

	"sniper/cmd/server/hook"
	example_v1 "sniper/rpc/example/v1"
	"sniper/server/exampleServerTest"

	"github.com/bilibili/twirp"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewCheckLogin(),
	hook.NewLog(),
)

func initMux(mux *http.ServeMux,isInternal bool) {

	server := &exampleServerTest.Server{}
	handler := example_v1.NewEchoServer(server, hooks)
	mux.Handle(example_v1.EchoPathPrefix, handler)
}

func initInternalMux(mux *http.ServeMux) {

}
