package main

import (
	"fmt"
	"strings"
	"github.com/gvalkov/golang-evdev"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	time.Sleep(time.Second * 1)
	fmt.Println("\033[1m","Keyboard Mistake Counter is running...","\033[0m")
	sigs := make(chan os.Signal,1)
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)
	go func(){
		sig := <-sigs
		fmt.Println()
		fmt.Println("   ",sig)
		message := "   "+"\033[32m"+"○"+"\033[0m"+" Ran successfully."
		fmt.Println(message)
		os.Exit(0)
	}()
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
		errormessage := "  \033[31m"+"×"+"\033[0m"+" Keyboard not found."
		fmt.Printf(errormessage)
		os.Exit(1)
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
				fmt.Printf("\033[2K\r   You made %s mistakes!", counter)
				if event.Value == 1 {
					if event.Code == 14{
					time.Sleep(time.Millisecond * 600)
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
