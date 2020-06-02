// K-Util - General Purpose Routines
// KEC - 06/2020

package kutil

import (
	"os"
	"os/exec"
	"time"
)

// ClearScreen - Utility to Clear the Screen by executing "clear" command.
func ClearScreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

// Pause - Pause execution for <psec> seconds.
func Pause(psec int) {
	time.Sleep(time.Duration(psec) * time.Second)
}
