package main

import (
	"log"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatal(err)
	}

	// For linux, need to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	for i := 0; i < 10000; i++ {
		//set keys
		kb.SetKeys(keybd_event.VK_D, keybd_event.VK_I, keybd_event.VK_R, keybd_event.VK_ENTER)
		//launch
		err = kb.Launching()
		if err != nil {
			log.Fatal(err)
		}
	}
}
