package main

type Stringer interface {
	Stringer()
}
type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	if r, ok := stringer.(Reader); ok {
		r.Read()
	}

}

func main() {

}
