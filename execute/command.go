package execute


type ICommand interface {
	IContext
}


type Command struct {
	Context
}
