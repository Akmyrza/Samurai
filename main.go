package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
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

	start := time.Now()

	var path = flag.String("path", "default value", "path to file")
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, _ := ioutil.ReadAll(f)

	var str []byte
	var Words []Word
	var found bool

	for _, elem := range res {

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

	elapsed := time.Since(start)
	fmt.Printf("time %s", elapsed)
}
