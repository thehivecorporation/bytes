package bytes_test

import (
	"fmt"

	"github.com/thehivecorporation/bytes"
)

func ExampleIntToBytes() {
	var n int = 300
	bytAr := bytes.IntToBytes(n)

	m := bytes.IntFromBytes(bytAr)

	fmt.Println(bytAr, m)

	// Output: [1 44] 300
}


func ExampleIntFromBytes() {
	var n int = 300
	bytAr := bytes.IntToBytes(n)

	m := bytes.IntFromBytes(bytAr)

	fmt.Println(bytAr, m)

	// Output: [1 44] 300
}

func ExampleEqual(){
	ar1 := []byte{1,2,3}
	ar2 := []byte{1,2,3}

	fmt.Println(bytes.Equal(ar1,ar2))

	// Output: true
}

func ExampleToBytes(){
	var n int = 300
	bytAr := bytes.ToBytes(n)

	m := bytes.IntFromBytes(bytAr)

	fmt.Println(bytAr, m)

	// Output: [1 44] 300
}

func ExampleIntFromBytes_incorrect(){
	var n int = 300
	bytAr := bytes.ToBytes(n)

	m := bytes.Float64FromBytes(bytAr)

	fmt.Println(m)

	// Output: 0
}