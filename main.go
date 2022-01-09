package main

import (
	"bytes"
	"errors"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"github.com/robotn/gohook"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func init() {
	add()
}

func main() {

}

func AltF4() {
	//This sucks a lot cause go-ps cant read the full process name from league for some reason, probs cause of wine
	activeWindow, err := GetActiveWindow()
	if err != nil {
		fmt.Println("Error on getActiveWindow", err, "output:", activeWindow)
	}
	fmt.Println(activeWindow)
	if activeWindow != "League of Legends (TM) Client\n" {
		return
	}
	if err != nil {
		fmt.Println(err)
	}

	processList, err := ps.Processes()
	if err != nil {
		log.Println("List process Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		executable := process.Executable()
		if strings.Contains(executable, "League") {
			fmt.Printf("%v %v %v\n", process.Pid(), executable, len(process.Executable()))
			if executable == "League of Legen" {
				KillProcess(process.Pid())
			}
		}
	}
}

func GetActiveWindow() (string, error) {
	proc := exec.Command("xdotool", "getwindowfocus", "getwindowname")
	var outb, errb bytes.Buffer
	proc.Stdout = &outb
	proc.Stderr = &errb
	err := proc.Run()
	if len(errb.String()) > 4 {
		return "", errors.New(errb.String())
	}
	if err != nil {
		return "", err
	}
	return outb.String(), nil
}

func KillProcess(pid int) error {
	proc := exec.Command("kill", "-TERM", strconv.Itoa(pid))
	var outb, errb bytes.Buffer
	proc.Stdout = &outb
	proc.Stderr = &errb
	err := proc.Run()
	if len(errb.String()) > 4 {
		return errors.New(errb.String())
	}
	if len(outb.String()) > 4 {
		return errors.New(outb.String())
	}
	if err != nil {
		return err
	}

	return nil
}

func add() {
	hook.Register(hook.KeyDown, []string{"alt", "f4"}, func(e hook.Event) {
		go AltF4()
	})

	s := hook.Start()
	<-hook.Process(s)
}
