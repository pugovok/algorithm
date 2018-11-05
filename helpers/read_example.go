package helpers

import (
	"bufio"
	"os"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func ReadInt32() int {
	scanner.Scan()
	ans, _ := strconv.Atoi(scanner.Text())
	return ans
}

func PrintInts(ints ...int) {
	for _, value := range ints {
		writer.WriteString(strconv.Itoa(value))
		writer.WriteByte(' ')
	}
}

func main() {
	defer writer.Flush() // очищает буфер при выходе из функции main
	scanner.Split(bufio.ScanWords) // заставляет scanner считать отдельными токенами слова, разделенные пробелами
	// также имеются функции bufio.ScanLines и bufio.ScanBytes
	var n, m int = ReadInt32(), ReadInt32()
	PrintInts(n, m)
	writer.WriteByte('\n')
	slice := []int{1, 2, 3, 4, 5}
	PrintInts(slice...)
}
