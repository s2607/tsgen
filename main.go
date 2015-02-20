package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

type student struct {
	name    string
	address string
	parents string
	bday    string //TODO: <should be time type or long int or whatever
	gender  rune
	gdate   string //TODO:see bday
	gyear   int    //gdate only contanes month/day, year is seperated out for auto gen table headers
	years   []schoolyear
}

func (this student) gpa() float32 {
	return 9000.1
}
func (this schoolyear) credits() int {
	var isum = 0
	for i, v := range this.classes {
		i = i
		isum = isum + v.credits
	}
	return isum
}

func (this student) credits() int {
	var isum = 0
	for i, v := range this.years {
		i = i //TODO: ugh
		isum = isum + v.credits()
	}
	return isum
}

type schoolyear struct {
	classes []class
}

type class struct {
	year    int
	subject string
	part    int
	grade   rune
	book    string
	credits int
}

func (this class) String() string {
	return fmt.Sprintf("%d,\t%s,\t%d,\t%c,\t%s", this.year, this.subject, this.part, this.grade, this.book)
}
func (this class) Html() string {
	if this.part > 0 {
		return fmt.Sprintf("<td>%s</td><td>%d</td><td>%c</td><td>%s</td><td>%d</td>", this.subject, this.part, this.grade, this.book, this.credits)
	} else {
		return fmt.Sprintf("<td>%s</td><td>N/A</td><td>%c</td><td>%s</td><td>%d</td>", this.subject, this.grade, this.book, this.credits)
	}
}
func (this schoolyear) String() string {
	s := "schoolyear\n"
	for i, v := range this.classes {
		i = i //TODO:ugh
		s = s + v.String()
		s = s + "\n"
	}
	return s

}

func (this student) String() string {
	s := this.name
	s += "\n"
	for i, v := range this.years {
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
	Class.credits = 1
	if len(feilds) > 5 {
		i, e = strconv.ParseInt(feilds[5], 0, 32)
		if e != nil {
			Class.credits = int(i)
		}
	}
	return Class

}

func (Student *student) tsvimport(in *bufio.Reader) {
	if len(Student.years) < 1 {
		Student.years = make([]schoolyear, 4)
	}
	for c := getclass(in); c != nil; c = getclass(in) {
		Student.years[c.year-9].classes = append(Student.years[c.year-9].classes, *c)
	}
}
func (year schoolyear) Html() string {
	//TODO: this should really use some kind of file pointer instead
	//eventually we may want to export weird formats (ex: pdf)
	s := "<table rules='all'><tr><th>Subject</th><th>Part</th><th>Grade letter</th><th>Book title</th><th>Credits</th></tr>"
	for i, v := range year.classes {
		s = s + "<tr>"
		i = i //TODO:ugh
		s = s + v.Html()
		s = s + "</tr>\n"
	}
	s = s + "</table>"
	return s

}

func (Student student) typetrans() string {
	//TODO: this should really use some kind of file pointer instead
	//eventually we may want to export weird formats (ex: pdf)
	s := "<html>"
	s = s + "<title>" + Student.name + "</title><body>\n"
	s = s + "<h1>" + Student.name + "</h1><br>\n"
	//	s = s + "<br><h2>Minfo</h2<br>"
	s = s + "Address: " + Student.address + "<br>\n"
	s = s + "Parrents: " + Student.parents + "<br>\n"
	s = s + "Birth Date: " + Student.bday + "<br>\n"
	s = s + "Graduation date: " + Student.gdate + fmt.Sprintf(" %d", Student.gyear) + "<br>\n"
	s = s + fmt.Sprintf("Total Credits earned:%d ", Student.credits())
	for i, v := range Student.years {
		s = s + "<h3> Year:" + fmt.Sprint(Student.gyear-4+i) + "-" + fmt.Sprint(Student.gyear-4+i+1) + "</h3><br>"
		s = s + v.Html()
		s = s + "<hr>\n"
	}
	s = s + "</table>"
	return s

}
func main() {
	me := new(student)
	me.name = "Stephen Wiley"
	me.gyear = 2011
	me.gender = 'M'
	me.address = "1180 Lickinghole rd. Goochland VA 23063"
	me.gdate = "July 4th"
	me.parents = "Glen and Jorie Wiley"
	me.bday = "May 24th 1994"
	reader := bufio.NewReader(os.Stdin)
	me.tsvimport(reader)
	fmt.Print(me.typetrans())
}
