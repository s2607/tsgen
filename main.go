package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

type class struct {
	year    int
	subject string
	part    int
	grade   rune
	book    string
}

func (this class) String() string {
	return fmt.Sprintf("%d,\t%s,\t%d,\t%c,\t%s", this.year, this.subject, this.part, this.grade, this.book)
}
func getClass(in *bufio.Reader) *class {
	Class := new(class)
	//TODO: do this correctly
	line, e := in.ReadString('\n')
	if e != nil {
		return nil
	}
	feilds := strings.Split(line, "\t")
	i, e := strconv.ParseInt(feilds[0], 0, 32)
	if e != nil {
		return nil
	}
	Class.year = int(i)
	Class.subject = feilds[1]
	i, e = strconv.ParseInt(feilds[2], 0, 32)
	if e != nil {
		return nil
	}
	Class.part = int(i)

	Class.grade = rune(feilds[3][0])
	Class.book = feilds[4]
	return Class

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	firstclass := getClass(reader)
	fmt.Println(firstclass)
}
