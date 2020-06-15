// K-Util - General Purpose Routines
// KEC - 06/2020
//
// ClearScreen Code From --
// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go

package kutil

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //Create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ClearScreen - Utility to Clear the Screen by executing "clear" command.
func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// Pause - Pause execution for <psec> seconds.
func Pause(psec int) {
	time.Sleep(time.Duration(psec) * time.Second)
}
