package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
)

var encoder *base64.Encoding

var punc1 = flag.String("p", "?", "first potential punctuation in password, default = ?")

var punc2 = flag.String("P", "!", "second potential punctuation in password, default = !")

func init() {
	flag.Parse()
	if len(*punc1) != 1 || len(*punc2) != 1 {
		log.Fatalln("only single character for potential punctuation")
	}
	encoder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" + *punc1 + *punc2)
}

var length = flag.Int("l", 12, "length of generated password, default = 12")

func main() {
	var bytelength int = (*length * 3) / 4
	if mod := bytelength % 3; mod != 0 {
		switch mod {
		case 1:
			bytelength += 2
		case 2:
			bytelength++
		}
	}
	buffer := make([]byte, bytelength)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatalf("unable to acquire random bytes: %v\n", err)
	}
	fmt.Println(encoder.EncodeToString(buffer)[:*length])
}
