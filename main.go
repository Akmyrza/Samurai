package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Word struct {
	Name  []byte
	Count int
}

func giveMeFileBro() *os.File {

	filepath := flag.String("path", "mobydick.txt", "path to file")
	flag.Parse()

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func readThatFileBro(thatfile *os.File) {

	var theword []byte
	var Words []Word
	var found bool

	reader := bufio.NewReader(thatfile)

	for {
		onebyte, err := reader.ReadByte()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if 65 <= onebyte && onebyte <= 90 { // from A to Z
			theword = append(theword, onebyte+32)
		} else if 97 <= onebyte && onebyte <= 122 { // a to z
			theword = append(theword, onebyte)
		} else {
			if theword != nil {

				found = false

				for index, word := range Words {
					if bytes.Compare(word.Name, theword) == 0 {
						Words[index].Count++
						found = true
						break
					}
				}

				if !found {
					Words = append(Words, Word{theword, 1})
				}

				theword = nil
			}
		}
	}

	sort.Slice(Words[:], func(i, j int) bool {
		return Words[i].Count > Words[j].Count
	})

	for _, word := range Words[:20] {
		str := fmt.Sprintf("%s", word.Name)
		fmt.Println(str, word.Count)
	}
}

func main() {

	thatfile := giveMeFileBro()
	readThatFileBro(thatfile)

	defer thatfile.Close()
}
