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

const tickRate = 20

var addr = fmt.Sprintf(":3000")

func main() {
	t := &ticker.Ticker{}
	r := session.NewRegistry()
	w := game.NewWorld(handler.NewGameHandler(r))
	c := protocol.NewCodec(&protobuf.Decoder{}, &protobuf.Encoder{})
	s := handler.NewSocketHandler(r, t, w, c)
	h := websocket.NewHandler(s)
	t.EveryTick(w.Update)
	go t.Run(tickRate)
	if err := http.ListenAndServe(addr, h); err != nil {
		return
	}
}
