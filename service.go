package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kardianos/service"
)

var (
	logger service.Logger
)

type pcService struct {
	exit chan struct{}
}

func (p *pcService) run() error {

	logger.Info("punch-clock service Start !!!")

	bot()

	for {
		select {
		case <-p.exit:
			logger.Info("punch-clock service Stop ...")
			return nil
		}
	}
}

func (p *pcService) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	go p.run()
	return nil
}

func (p *pcService) Stop(s service.Service) error {
	close(p.exit)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "punch-clock-service",
		DisplayName: "punch-clock-service",
		Description: "This is remote dakoku service.",
	}

	// Create RD-Service service
	program := &pcService{}
	s, err := service.New(program, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the logger
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal()
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Failed (%s) : %s\n", os.Args[1], err)
		} else {
			fmt.Printf("Succeeded (%s)\n", os.Args[1])
		}
		return
	}

	// run in terminal
	s.Run()
}
