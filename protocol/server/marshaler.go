package server

type Marshaler interface {
	Pong(pong *Pong) ([]byte, error)
	JoinGame(joinGame *JoinGame) ([]byte, error)
	DestroyEntity(destroyEntity *DestroyEntity) ([]byte, error)
	SpawnPlayer(spawnPlayer *SpawnPlayer) ([]byte, error)
	UpdatePlayer(updatePlayer *UpdatePlayer) ([]byte, error)
}
