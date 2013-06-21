package sysinfo

import (
	"fmt"
	t "testing"
)

func TestGet(t *t.T) {
	fmt.Println(Get())
	fmt.Println()
	fmt.Println(Get().ToString())
	fmt.Println()
}

func ExampleGet() {
	si := Get()
	fmt.Printf("%v\n", si.Loads)
	fmt.Println(si.ToString())
}


/*
amd64
BenchmarkGet		 1000000	      2429 ns/op
BenchmarkString		  200000	     11258 ns/op
BenchmarkToString	  200000	     13702 ns/op

arm
BenchmarkGet		  200000	      8560 ns/op
BenchmarkString		   10000	    203131 ns/op
BenchmarkToString	   10000	    242835 ns/op
*/
func BenchmarkGet(b *t.B) {
	for i := 0; i < b.N; i++ {
		Get()
	}
}

func BenchmarkString(b *t.B) {
	for i := 0; i < b.N; i++ {
		Get().String()
	}
}

func BenchmarkToString(b *t.B) {
	for i := 0; i < b.N; i++ {
		Get().ToString()
	}
}
