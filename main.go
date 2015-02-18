package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

type student struct {
	classes []class
	gpa     float32
	name    string
}

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
func (this student) String() string {
	s := this.name
	s += "\n"
	for i, v := range this.classes {
		i = i //TODO:ugh
		s = s + v.String()
		s = s + "\n"
	}
	return s

}
func getclass(in *bufio.Reader) *class {
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
		Class.part = 0
	} else {
		Class.part = int(i)
	}

	Class.grade = rune(feilds[3][0])
	if len(feilds) > 4 {
		Class.book = feilds[4]
	}
	return Class

}

func (Student *student) tsvimport(in *bufio.Reader) {
	for c := getclass(in); c != nil; c = getclass(in) {
		Student.classes = append(Student.classes, *c)
	}
}

func main() {
	me := new(student)
	me.name = "Stephen Wiley"
	reader := bufio.NewReader(os.Stdin)
	me.tsvimport(reader)
	fmt.Print(me)
}
