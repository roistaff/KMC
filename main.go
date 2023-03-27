package main

import (
	"fmt"
	"strings"
	"github.com/gvalkov/golang-evdev"
)

func main() {
	fmt.Println("\033[1m","Keyboard Mistake Counter is running...","\033[0m")
	devices, err := evdev.ListInputDevices()
	if err != nil {
		panic(err)
	}
	var keyboard *evdev.InputDevice
	for _, device := range devices {
		if strings.Contains(device.Name, "Keyboard") {
			keyboard = device
			break
		}
	}
	if keyboard == nil { 
		panic("Keyboard not found")
	}
	count := 0
	for {
		events, err := keyboard.Read()
		if err != nil {
			panic(err)
		}
		for _, event := range events {
			if event.Type == evdev.EV_KEY {
				counter := "\033[33m" + fmt.Sprint(count) + "\033[0m"
				fmt.Printf("\033[2K\rYou made %s mistakes!", counter)
				if event.Value == 1 {
					if event.Code == 14{
					count++
				}
				} else if event.Value == 2 {
					if event.Code == 14{
					count++
				}
				}
			}
		}
	}
}

