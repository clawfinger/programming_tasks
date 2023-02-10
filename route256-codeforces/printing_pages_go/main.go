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
	zeroFound := false
	zeroStart := -1
	currentZeros := 0
	builderDirty := false
	for i := 1; i < len(printedPages); i++ {
		if printedPages[i] == 0 {
			if !zeroFound {
				zeroFound = true
				zeroStart = i
			}
			currentZeros++
		} else {
			if zeroFound {
				if currentZeros == 1 {
					if builderDirty {
						sb.WriteRune(',')
					}

					sb.WriteString(strconv.Itoa(zeroStart))
					builderDirty = true
				} else {
					end := zeroStart + currentZeros - 1
					if builderDirty {
						sb.WriteRune(',')
					}
					separator := '-'
					sb.WriteString(strconv.Itoa(zeroStart))
					sb.WriteRune(separator)
					sb.WriteString(strconv.Itoa(end))
					builderDirty = true
				}
			}
			zeroFound = false
			currentZeros = 0
		}

		if i == pageCount {
			if zeroFound {
				if currentZeros == 1 {
					if builderDirty {
						sb.WriteRune(',')
					}
					sb.WriteString(strconv.Itoa(zeroStart))
					builderDirty = true

				} else {
					if builderDirty {
						sb.WriteRune(',')
					}
					end := zeroStart + currentZeros - 1
					separator := '-'
					sb.WriteString(strconv.Itoa(zeroStart))
					sb.WriteRune(separator)
					sb.WriteString(strconv.Itoa(end))
					builderDirty = true

				}
			}
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
