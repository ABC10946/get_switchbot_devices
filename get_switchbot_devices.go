package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/nasa9084/go-switchbot"
)

var (
	openToken = flag.String("t", "", "")
)

func main() {
	flag.Parse()
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	if *openToken == "" {
		return errors.New("-t option is required")
	}

	c := switchbot.New(*openToken)
	devices, _, err := c.Device().List(context.Background())
	if err != nil {
		return fmt.Errorf("gettng device list")
	}

	for _, device := range devices {
		fmt.Printf("%s %s\n", device.Name, device.ID)
	}

	return nil
}
