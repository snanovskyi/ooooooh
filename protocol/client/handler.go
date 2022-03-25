package client

type Handler interface {
	Ping(ping *Ping)
	MovePlayer(movePlayer *MovePlayer)
}
