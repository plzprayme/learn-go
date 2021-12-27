package __2

import "fmt"

func Example_slicing() {
	// 배열: [5]int{1,2,3,4,5} 슬라이드: []int{1,2,3,4,5}
	// 빈 슬라이스에는 nil 값이 들어있다.
	// 1차원 슬라이스 만들기
	// Go에서는 음수 쓸 수 없다. nums[:-1] 안됨.
	// nums[:len(nums) - 1]로 쓰자
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
}

func Example_slicing2() {
	// 1차원 슬라이스 5개 만들기

	// make 로 만든 슬라이스들에는 해당 자료형의 기본값이 들어간다.
	// string: "" , int: 0, pointer, slice, map, interface ... : nil
	fruits := make([]string, 3) // 용량이 3인 슬라이스생성
	fruits[0] = "바나나"
	fruits[1] = "사과"
	fruits[2] = "토마토"
	fmt.Println(fruits)
	// Output:
	// [바나나 사과 토마토]
}

func Example_slicing3() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
	// [사과 바나나 토마토]
	// [포도 딸기]
	// [사과 바나나 토마토 포도 딸기]
	// [사과 바나나 포도 딸기]
}

func Example_slicing4() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()

	sliced3 := nums[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
	// Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// [1 2 3]
	// len: 3
	// cap: 5
	//
	// [3 4 5]
	// len: 3
	// cap: 3
	//
	// [1 2 3 4]
	// len: 4
	// cap: 5
	//
	// [1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}
