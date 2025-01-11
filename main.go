package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	listener, err := InitializeApp()

	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	errChan := make(chan error, 2)

	go func() {
		if err := listener.Start(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		log.Printf("Erro: %v", err)
	case sig := <-sigChan:
		log.Printf("Sinal recebido: %v", sig)
	}

	if err := listener.Close(); err != nil {
		log.Printf("Erro ao fechar listener1: %v", err)
	}

}
