package main

import (
	"bufio"
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

type ByCount []Word

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func CompareTwoByteSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {

	filepath := flag.String("path", "mobydick.txt", "path to file")
	flag.Parse()

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var str []byte
	var Words []Word
	var found bool

	reader := bufio.NewReader(file)

	for {
		elem, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if 65 <= elem && elem <= 90 {
			str = append(str, elem+32)
		} else if 97 <= elem && elem <= 122 {
			str = append(str, elem)
		} else {
			if str != nil {

				found = false

				for i, v := range Words {
					if CompareTwoByteSlices(v.Name, str) {
						Words[i].Count++
						found = true
						break
					}
				}
				if !found {
					Words = append(Words, Word{str, 1})
				}

				str = nil
			}

		}

	}

	sort.Sort(ByCount(Words))

	for i, elem := range Words {
		stringfrombytes := fmt.Sprintf("%s", elem.Name)
		fmt.Println(stringfrombytes, elem.Count)
		if i == 20 {
			break
		}
	}

}
