package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/maestre3d/gtavd/dlclist"
)

func main() {
	ctx := context.Background()
	go dlclist.Watch(ctx)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	ctx.Done()
}
