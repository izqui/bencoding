package main

import (
	"flag"
	"fmt"
	"github.com/izqui/bencoding"
	"io/ioutil"
	"os"
)

var filePath = flag.String("torrent", "/Users/izqui/Dropbox/HackerSchool/bencoding/example/element.torrent", "Torrent file to open")

func init() {

	flag.Parse()
}
func main() {

	file, err := os.Open(*filePath)
	panicOnError(err)

	d, err := ioutil.ReadAll(file)
	panicOnError(err)

	b, _ := bencoding.Decode(d)

	fmt.Println(b)
}

func panicOnError(err error) {

	if err != nil {

		panic(err)
	}
}
