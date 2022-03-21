package main

import (
	"github.com/kardianos/service"
	"github.com/maestre3d/gtavd"
)

func main() {
	s, err := service.New(&gtavd.Daemon{}, &service.Config{
		Name:        "GrandTheftAutoVDaemon",
		DisplayName: "Grand Theft Auto V Modification Daemon",
		Description: "Runs background tasks related to modifications of the Grand Theft Auto V game",
	})
	if err != nil {
		panic(err)
	}
	if err = s.Run(); err != nil {
		panic(err)
	}
}
