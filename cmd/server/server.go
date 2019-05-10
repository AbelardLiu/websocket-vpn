package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"lll.github.com/websocket-vpn/cmd/server/app"
	"lll.github.com/websocket-vpn/pkg/websocket/server"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewVpnServerCommand(server.SetupSignalHandler())
	if err := command.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}