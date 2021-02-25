package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/reujab/wallpaper"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	img := "https://filmdaily.co/wp-content/uploads/2021/02/rick-lede-1-1300x650.jpg"

	wallpaper.SetFromURL(img)

	fmt.Println("sleeping for 2s")
	time.Sleep(time.Second * 2)

	openBrowser("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
}
