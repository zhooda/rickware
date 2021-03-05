package main

import (
	"log"

	"github.com/itchyny/volume-go"
)

var VOL_LIMIT int = 50

func main() {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	if vol < 50 {
		err = volume.SetVolume(50)
		if err != nil {
			log.Fatalf("set volume failed: %+v", err)
		}
	}
	for {
		vol, err := volume.GetVolume()
		if err != nil {
			log.Fatalf("get volume failed: %+v", err)
		}
		if vol < VOL_LIMIT {
			VOL_LIMIT += 10
			if VOL_LIMIT > 100 {
				VOL_LIMIT = 100
			}
			err = volume.SetVolume(VOL_LIMIT)
			if err != nil {
				log.Fatalf("set volume failed: %+v", err)
			}
		}
	}
}
