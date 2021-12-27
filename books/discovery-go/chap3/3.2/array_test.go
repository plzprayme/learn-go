package __2

import "fmt"

func Example_array() {
	// 배열은 고정된 크기
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruits := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruits)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}
