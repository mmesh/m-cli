package command

type Interface interface {
	Exec(args []string)
	// Transfer(args []string)
	PortFwd()
}
type API struct{}

func Resource() Interface {
	return &API{}
}
