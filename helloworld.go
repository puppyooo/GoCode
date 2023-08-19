package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch = make(chan string, 10)

func load(url string) {
	fmt.Println("start to load", url)
	time.Sleep(time.Second)
	ch <- url
}

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}

type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is empty")
	}
	fmt.Printf("Hello,", name)
	return nil
}

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from", r)
			ret = -1
		}
	}()
	arr := [3]int{1, 2, 3}
	return arr[index]
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func main() {
	// 添加元素，切片容量可以根据需要自动扩展
	slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3容量为5的切片
	slice2 = append(slice2, 1, 2, 3, 4)   // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12
	// 子切片 [start, end)
	sub1 := slice2[3:]  // [1 2 3 4]
	sub2 := slice2[:3]  // [0 0 0]
	sub3 := slice2[1:4] // [0 0 1]
	// 合并切片
	combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]

	fmt.Println(sub1, sub2, sub3, combined)

	// map
	m1 := make(map[string]int)
	// define map
	m2 := map[string]string{
		"puppy": "dog",
		"dog":   "animal",
	}
	m1["puppy"] = 18
	fmt.Println(m2)

	str := "Golang"
	var p *string = &str
	*p = "Hello"
	fmt.Println(str)

	age := 18
	if age < 18 {
		fmt.Printf("Kid")
	} else {
		fmt.Printf("Adult")
	}

	if age := 18; age < 18 {
		fmt.Printf("Kid")
	} else {
		fmt.Printf("Adult")
	}

	type Gender int8
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)

	gender := MALE

	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}

	// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	nums := []int{10, 20, 30}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	quo, rem := div(10, 3)
	fmt.Println(add(quo, rem))

	if err := hello(""); err != nil {
		fmt.Println(err)
	}

	// fmt.Println(get(10))
	fmt.Println("finished")

	// struct
	stu := &Student{
		name: "puppy",
		age:  18,
	}
	msg := stu.hello("dog")
	fmt.Println(msg)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")

	// chanel
	for i := 0; i < 3; i++ {
		go load("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
	fmt.Println("Done!")
}
