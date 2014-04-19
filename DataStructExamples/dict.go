package main

import (
	"Trie"
	// "flag"
	"fmt"
	"io/ioutil"
	"os"
	// "path/filepath"
	// "strconv"
	"strings"
)

// test data : ./dict ac.txt dict.txt
func main() {
	if len(os.Args) == 3 {
		printNewWord(os.Args[1], os.Args[2])

	} else {
		fmt.Println("Please input the article and dictionary file")
	}
}

func printNewWord(acfile, dictfile string) {
	word := parseActicle(acfile)
	ow := parseDict(dictfile)
	var newWord []string
flag:
	for _, v := range word {
		prefix := ow.Search(v) //Search from dict
		// fmt.Println(v, " prefix:  ", prefix, len(prefix))
		if len(prefix) > 0 && prefix[0] != " " {
			continue flag
		}
		for _, d := range newWord {
			if d == v {
				continue flag
			}
		}
		newWord = append(newWord, v)
	}
	fmt.Println("The new words are :")
	fmt.Println(newWord)
}

func parseDict(dictfile string) *Trie.Trie { //the dict file shoud have this format: word-单词\n
	dict, err := ioutil.ReadFile(dictfile)
	// fmt.Println(string(dict))
	t := Trie.NewTrie()
	handleError(err)
	word := strings.Split(string(dict), "\n")
	word = word[:len(word)-1]
	// fmt.Println(len(word), word)
	fmt.Println("The dictionary is :")
	for _, v := range word {
		fmt.Println(v)
		k := strings.Split(v, "-")
		t.Add(k[0], k[1])
	}

	return t
}

func parseActicle(acfile string) []string {
	acticle, err := ioutil.ReadFile(acfile)
	// fmt.Println(string(acticle))
	handleError(err)
	acticle = []byte(strings.TrimSuffix(string(acticle), "\n"))
	word := strings.Split(string(acticle), " ")
	fmt.Println("The article is :")
	fmt.Println(word)
	return word

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
