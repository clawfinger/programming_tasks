package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doWork(in *bufio.Reader, out *bufio.Writer) {
	var pageCount int
	fmt.Fscan(in, &pageCount)
	var printedPagesLine string
	fmt.Fscan(in, &printedPagesLine)

	printedPages := make([]int, pageCount+1)

	entries := strings.Split(printedPagesLine, ",")
	for _, entry := range entries {
		splittedEntry := strings.Split(entry, "-")
		if len(splittedEntry) > 1 {
			from, _ := strconv.Atoi(splittedEntry[0])
			to, _ := strconv.Atoi(splittedEntry[1])
			for i := from; i <= to; i++ {
				printedPages[i] = i
			}
		} else {
			page, _ := strconv.Atoi(splittedEntry[0])
			printedPages[page] = page
		}
	}
	var sb strings.Builder

	firstWrite := true
	startIndex := 1
	isOnZero := printedPages[1] == 0
	for i := 2; i < len(printedPages); i++ {
		if isOnZero {
			if printedPages[i] != 0 {
				if i-startIndex == 1 {
					if !firstWrite {
						sb.WriteRune(',')
					}
					sb.WriteString(strconv.Itoa(startIndex))
					firstWrite = false
				} else {
					if !firstWrite {
						sb.WriteRune(',')
					}
					sb.WriteString(strconv.Itoa(startIndex))
					sb.WriteRune('-')
					sb.WriteString(strconv.Itoa(i - 1))
					firstWrite = false
				}
			}
		} else {
			if printedPages[i] == 0 {
				startIndex = i
			}
		}
		isOnZero = printedPages[i] == 0
	}

	if isOnZero {
		if startIndex == len(printedPages)-1 {
			if !firstWrite {
				sb.WriteRune(',')
			}
			sb.WriteString(strconv.Itoa(startIndex))
			firstWrite = false
		} else {
			if !firstWrite {
				sb.WriteRune(',')
			}
			sb.WriteString(strconv.Itoa(startIndex))
			sb.WriteRune('-')
			sb.WriteString(strconv.Itoa(len(printedPages)))
			firstWrite = false
		}
	}
	res := sb.String()
	sb.Reset()
	fmt.Fprintln(out, res)
	out.Flush()
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var checksCount int
	fmt.Fscan(in, &checksCount)
	for i := 0; i < checksCount; i++ {
		doWork(in, out)
	}
	defer out.Flush()
}
