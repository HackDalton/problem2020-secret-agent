package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/auyer/steganography"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inFile, err := os.Open("nhb.in.png") // opening file
	check(err)
	reader := bufio.NewReader(inFile) // buffer reader
	img, err := png.Decode(reader)    // decoding to golang's image.Image
	check(err)

	w := new(bytes.Buffer)                                                                  // buffer that will recieve the results
	err = steganography.Encode(w, img, []byte("hackDalton{w3lc0m3_t0_l4ngl3y_kgLQq-1mHU}")) // Encode the message into the image
	check(err)
	outFile, err := os.Create("nhb.out.png") // create file
	check(err)
	_, err = w.WriteTo(outFile) // write buffer to it
	check(err)
	err = outFile.Close()
	check(err)

	fmt.Println(decodeFile("nhb.out.png"))
}

func decodeFile(file string) string {
	inFile, err := os.Open(file) // opening file
	check(err)
	defer inFile.Close()

	reader := bufio.NewReader(inFile) // buffer reader
	img, err := png.Decode(reader)    // decoding to golang's image.Image
	check(err)

	sizeOfMessage := steganography.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

	msg := steganography.Decode(sizeOfMessage, img) // decoding the message from the file
	return string(msg)
}
