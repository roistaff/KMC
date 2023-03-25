package main

import (
	"fmt"
	"github.com/gvalkov/golang-evdev"
	"strings"
)

func search() string {
	devices, err := evdev.ListInputDevices()
	if err != nil {
		panic(err)
	}

	for _, dev := range devices {
		fmt.Printf("Device name: %s, FN: %s\n", dev.Name, dev.Fn)
		if strings.Contains(dev.Name, "Keyboard") == true {
			devicepath := dev.Fn
			return devicepath
		}
	}
	return ""
}

func Count(path string) int {
	var count int
	if path == "" {
		panic("Keyboard device not found")
	}
	dev, err := evdev.Open(path)
	if err != nil {
		panic(err)
	}
	defer dev.File.Close()
	for {
		events, err := dev.Read()
		if err != nil {
			panic(err)
		}
		for i := range events {
			event := &events[i]
			if event.Type == evdev.EV_KEY {
				keyCode := event.Code
				keyValue := event.Value
				if keyValue != 0 {
					count++
				}
				if keyCode == 14 {
					fmt.Printf("\r backspace key. You made %d mistakes.", count/2)
				}
			}
		}
	}
}

func main() {
	path := search()
	Count(path)
}

