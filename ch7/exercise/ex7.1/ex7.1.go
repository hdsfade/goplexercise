//@author: hdsfade
//@date: 2020-11-06-08:32
package ex7_1

import (
	"bufio"
	"strings"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	*w += WordCounter(count)
	return count, nil
}

func (w *LineCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	for scanner.Scan() {
		count++
	}
	*w += LineCounter(count)
	return count, nil
}
