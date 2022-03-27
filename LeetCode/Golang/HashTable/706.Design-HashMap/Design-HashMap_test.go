package hashtable

import (
	"fmt"
	"testing"
)

func Test_Problem706(t *testing.T) {
	hm := Constructor706()
	fmt.Printf("HashMap: %v\n", hm)
	hm.put(1, 10)
	hm.put(2, 20)
	hm.put(3, 30)
	fmt.Printf("HashMap: %v\n", hm)
	fmt.Printf("Get 2: %v\n", hm.Get(2))
	fmt.Printf("Get 4: %v\n", hm.Get(4))
	hm.Remove(2)
	fmt.Printf("Remove 2: %v\n", hm)
}
