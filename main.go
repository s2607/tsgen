package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

type class struct {
	year    int
	subject string
	grade   rune
	book    string
}

func (this class) String() string {
	return fmt.Sprintf("%d\t%s\t%c\t%s", this.year, this.subject, this.grade, this.book)
}
func getClass(in *bufio.Reader) *class {
	Class := new(class)
	//TODO: do this correctly
	line, e := in.ReadString('\n')
	if e == nil {
		feilds := strings.Split(line, "\t")
		i, e := strconv.ParseInt(feilds[0], 0, 32)
		Class.year = int(i)
		Class.subject = feilds[1]
		Class.grade = rune(feilds[2][0])
		Class.book = feilds[3]
		if e == nil {
			return Class
		}
	}
	return nil

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	firstclass := getClass(reader)
	fmt.Println(firstclass)
}
