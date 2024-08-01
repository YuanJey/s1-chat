package server

type Server interface {
	StartServer()
	Work(msg []byte)
}
