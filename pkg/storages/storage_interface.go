package storages

// stotages have actual Database

type GreetStorager interface {
	GetGreet() string
	PutGreet(name string) error
}

type ElizaStorager interface {
	GetEliza() string
	PutEliza(name string) error
}
