package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := fileSHA1("http.log.gz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)
	//ef7dc39754fd23f7a1a8657e2ffda49edc49fff9

	sig, err = fileSHA1("sha1.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)
}

func fileSHA1(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz") {

		r, err = gzip.NewReader(file)
		if err != nil {
			return "", err
		}
	}

	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	} //copy from reader to writer

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
