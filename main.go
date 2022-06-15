package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type Word struct {
	Name  string
	Count int
}

type ByCount []Word

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	f, _ := os.Open("mobydick.txt")
	defer f.Close()
	
	res, _ := ioutil.ReadAll(f)

	var str string
	var Words []Word
	var found bool = false

	for _, elem := range res {

		if 65 <= elem && elem <= 90 {
			str = str + string(elem+32)
		} else if 97 <= elem && elem <= 122 {
			str = str + string(elem)
		} else {
			//if elem == 32 || elem == 10 {
				if str != "" {

					found = false

					for i, v := range Words {
						if v.Name == str {
							Words[i].Count++
							found = true
							break
						}
					}

					if !found {
						Words = append(Words, Word{str, 1})
					}

					str = ""
				}
			//}
		}

	}

	sort.Sort(ByCount(Words))

	for i, elem := range Words {
		fmt.Println(elem)
		if i == 20 {
			break
		}
	}

}
