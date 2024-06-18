package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/AnuragProg/ssh-portfolio/ui"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

const (
	HOST = "localhost"
	PORT = "8000"
)

func main() {
	done := make(chan os.Signal)

	sshServer, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(HOST, PORT)),
		wish.WithHostKeyPath("./.ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(ui.UIHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		panic(err.Error())
	}
	defer sshServer.Shutdown(context.Background())


	go func() {
		if err := sshServer.ListenAndServe(); err != nil {
			log.Println(err.Error())
			done<- nil
		}
	}()
	log.Println("Server listening on", net.JoinHostPort(HOST, PORT))

	<-done
	log.Println("Shutting down server...")
}
