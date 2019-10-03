package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net", "io", "co.uk"}

const allowedChars = "abcdefghijklmnoprstuvwxyz0123456789_-"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}

		checkEnd := newText[len(newText)-1:]
		ntts := string(newText)
		if checkEnd[0] == '-' {
			ntts = strings.TrimRight(ntts, "-")
		}
		fmt.Println(ntts + "." + tlds[rand.Intn(len(tlds))])
	}
}
