package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nasa9084/go-switchbot"
)

func main() {
	flag.Parse()
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

func execute() error {
	openToken := os.Getenv("OPEN_TOKEN")
	secretToken := os.Getenv("SECRET_TOKEN")

	c := switchbot.New(openToken, secretToken)
	devices, infrared, err := c.Device().List(context.Background())
	if err != nil {
		return fmt.Errorf("gettng device list")
	}

	fmt.Println("Switchbot Devices")

	for _, device := range devices {
		fmt.Printf("%s %s %s\n", device.Type, device.Name, device.ID)
	}

	fmt.Println("-----------------")

	fmt.Println("Infrared Devices")

	for _, device := range infrared {
		fmt.Printf("%s %s %s\n", device.Type, device.Name, device.ID)
	}

	return nil
}
