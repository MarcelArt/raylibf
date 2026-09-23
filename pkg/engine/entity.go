package engine

type IEntity interface {
	Update()
	Draw()
}

type Entity struct {
	IsActive bool
}
