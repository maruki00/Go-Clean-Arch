package contracts




type CommandBus interface {
	Apply(data any) error 
}
