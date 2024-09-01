package math

import "fmt"

type mytype struct {
	size string
	hash uint32
}

type myinterface interface {
	do() string
}

func newMytype(size string, hash uint32) myinterface {
	return &mytype{size, hash}
}

func (r *mytype) doPrivate() string {
	return r.size
}

func (r *mytype) do() string {
	return r.doPrivate()
}

func main() {
	//wrong - DO NOT instantiate struct
	a := &mytype{"1",1}
	a.doPrivate()
	// with constructor - Use this
	t := newMytype("100", 100)
	//t.do()
	fmt.Println(t.do())
	// t.doPrivate() // t.doPrivate undefined (type myinterface has no field or method doPrivate)

	// without constructor - You may use this too 
	t2 := myinterface(&mytype{"200", 200})
	fmt.Println(t2.do())
	//t2.do()
	// t.doPrivate() // t.doPrivate undefined (type myinterface has no field or method doPrivate)
}
