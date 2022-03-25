package game

type Handler interface {
	DestroyEntity(entity Entity)
	SpawnPlayer(player *Player)
	UpdatePlayer(player *Player)
}
