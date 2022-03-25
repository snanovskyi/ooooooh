package main

import (
	"fmt"
	"net/http"

	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/handler"
	"github.com/snanovskyi/ooooooh/protobuf"
	"github.com/snanovskyi/ooooooh/protocol"
	"github.com/snanovskyi/ooooooh/session"
	"github.com/snanovskyi/ooooooh/ticker"
	"github.com/snanovskyi/ooooooh/websocket"
)

var addr = fmt.Sprintf(":3000")

func main() {
	t := &ticker.Ticker{}
	r := session.NewRegistry()
	w := game.NewWorld(handler.NewGameHandler(r))
	c := protocol.NewCodec(protobuf.NewDecoder(), protobuf.NewEncoder())
	s := handler.NewSocketHandler(r, t, w, c)
	h := websocket.NewHandler(s)
	t.EveryTick(w.Update)
	go t.Run()
	if err := http.ListenAndServe(addr, h); err != nil {
		return
	}
}
