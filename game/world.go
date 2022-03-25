package game

type World struct {
	handler  Handler
	nextID   uint32
	entities map[uint32]Entity
}

func NewWorld(h Handler) *World {
	return &World{
		handler:  h,
		entities: make(map[uint32]Entity),
	}
}

func (w *World) Handler() Handler {
	return w.handler
}

func (w *World) Entities() []Entity {
	entities := make([]Entity, len(w.entities))
	i := 0
	for _, e := range w.entities {
		entities[i] = e
		i++
	}
	return entities
}

func (w *World) NewID() uint32 {
	id := w.nextID
	w.nextID++
	return id
}

func (w *World) Spawn(entity Entity) {
	w.entities[entity.ID()] = entity
	entity.Spawn()
}

func (w *World) Destroy(entity Entity) {
	entity.Destroy()
	delete(w.entities, entity.ID())
	w.handler.DestroyEntity(entity)
}

func (w *World) Update() {
	for _, e := range w.entities {
		e.Update()
	}
}
