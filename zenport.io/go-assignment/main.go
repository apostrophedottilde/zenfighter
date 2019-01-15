package main

import (
	"gitlab.com/zenport.io/go-assignment/adapters/http"
	"gitlab.com/zenport.io/go-assignment/engine"
	"gitlab.com/zenport.io/go-assignment/providers/database"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	provider := database.NewProvider()
	e := engine.NewEngine(provider)

	adapter := http.NewHTTPAdapter(e)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
	provider.Close()
}
