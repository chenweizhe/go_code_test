package wordcount

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	Key string
	Value int
}

// PairList 实现了sort接口， 可以使用sort.Sort对其排序
 type PairList []Pair

 func (p PairList) Swap(i,j int)  { p[i], p[j] = p[j], p[i]}
 func (p PairList) Len() int { return len(p)}
 func (p PairList) Less(i, j int) bool {return p[j].Value < p[j].Value} //逆序

//  提取单词
func  SplitOnNonLetter(s string) []string  {
	notALetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FiledsFunc(s, notALetter)
}

/*
   基于map实现了类型WordCount, 并对期实现了Merge(), Report(), SortReport(), UpdateFreq(),
   WordFreqCounter() 方法
*/

type WordCount map[string]int

// 用于实现合并两个wordcount
func (source WordCount) Merge(wordcount WordCount) WordCount {
	for k, v := range wordcount {
		source[k] += v
	}
	return source
}

// 打印词频统计情况
func (wordcount WordCount) Report() {
	words := make([]string, 0, len(wordcount))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range wordcount {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}

	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %s%s\n", gap, " ", "Frequency")
	
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, wordcount[word])
	}

}

// 从多到少打印词频
func (wordcount WordCount) SortReport() {

	p := make(PairList, len(wordcount))
	i := 0
	for k, v := range wordcount {
		// 将wordcount map转换成PairList
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	wordWidth, frequencyWidth := 0, 0
	for _, pair := pair.Key, pair.Value
	



}