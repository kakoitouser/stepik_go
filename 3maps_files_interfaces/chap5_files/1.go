package main //уже объявлен.
import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	sum := 0
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		sum += num
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
