// Package keyboard handles the low
// level interaction with the OS to
// trigger keyboard inputs.
//
// @see github.com/micmonay/keybd_event
package keyboard

import (
	kb "github.com/micmonay/keybd_event"
	"log"
	"os"
	"runtime"
	"time"
)

// Given a key returns the int mapped
// for the OS we're running on.
// Note: only a few keys are mapped here
// as these are the ones that you will
// need the most. Happy to add on-demand ;-)
func getKeyBindings(k []string) []int {
	var m = make(map[string]int)
	m["up"] = kb.VK_UP
	m["left"] = kb.VK_LEFT
	m["right"] = kb.VK_RIGHT
	m["down"] = kb.VK_DOWN
	m["a"] = kb.VK_A
	m["b"] = kb.VK_B
	m["c"] = kb.VK_C
	m["d"] = kb.VK_D
	m["e"] = kb.VK_E
	m["f"] = kb.VK_F
	m["g"] = kb.VK_G
	m["h"] = kb.VK_H
	m["i"] = kb.VK_I
	m["l"] = kb.VK_L
	m["m"] = kb.VK_M
	m["n"] = kb.VK_N
	m["o"] = kb.VK_O
	m["p"] = kb.VK_P
	m["q"] = kb.VK_Q
	m["r"] = kb.VK_R
	m["s"] = kb.VK_S
	m["t"] = kb.VK_T
	m["u"] = kb.VK_U
	m["v"] = kb.VK_V
	m["z"] = kb.VK_Z

	v := make([]int, 0)

	for _, value := range k {
		if m[value] > 0 {
			v = append(v, m[value])
		}
	}

	return v
}

// Utility to check if a list contains
// the given element
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// Presses the given keys.
func Press(keys []string) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	kb, err := kb.NewKeyBonding()

	if err != nil {
		logger.Print("error creating keybinding", err)
	}

	// @see https://github.com/micmonay/keybd_event/blob/fa4361d5d1fe03cfa1493f538283b8e92c6e937a/README.md#an-example-
	if runtime.GOOS == "linux" {
		time.Sleep(50 * time.Millisecond)
	}

	keyBindings := getKeyBindings(keys)
	kb.SetKeys(keyBindings...)

	if contains(keys, "shift") {
		kb.HasSHIFT(true)
	}

	if contains(keys, "ctrl") {
		kb.HasCTRL(true)
	}

	logger.Print("pressing: ", keys, keyBindings)

	err = kb.Launching()

	if err != nil {
		logger.Print("error pressing", err)
	}
}
