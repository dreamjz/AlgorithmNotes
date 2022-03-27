package hashtbale

import (
	"fmt"
	"testing"
)

func Test_Problem705(t *testing.T) {
	hs := Constructor705()
	fmt.Printf("HashSet: %v \n", hs)
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	fmt.Printf("HashSet: %v \n", hs)
	hs.Add(2)
	fmt.Printf("Add same key: 2, HashSet: %v \n", hs)
	fmt.Printf(" Contains 2: %v\n", hs.Contains(2))
	fmt.Printf(" Contains 5: %v\n", hs.Contains(5))
	hs.Remove(2)
	fmt.Printf(" After Remove 2: %v\n", hs)
	fmt.Println(hs.String())
}
