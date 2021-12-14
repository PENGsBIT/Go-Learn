package 数据结构

import "fmt"

type Set map[interface{}]struct{}

func (s *Set) Add(k interface{}) {
	(*s)[k] = struct{}{}
}

func (s *Set) Remove(k interface{}) {
	delete(*s, k)
}

func (s *Set) Has(k interface{}) bool {
	_, ok := (*s)[k]
	return ok
}

func (s *Set) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Set) Len() int {
	return len(*s)
}

func (s *Set) All() []interface{} {
	var res []interface{}
	for i, _ := range *s {
		res = append(res, i)
	}
	return res
}

func main() {
	var s Set = Set{}
	s.Add("sss")
	s.Add("sss")
	s.Add("sss2")

	s.Add("sss4")
	s.Add(123)
	stu := struct {
		Name string
		Age  int
	}{Name: "someone", Age: 32}
	s.Add(stu)

	fmt.Println(s.All()) //[sss2 sss4 123 {someone 32} sss]
}
