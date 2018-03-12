package main

import "errors"

func main() {
	a := test()
	println(len(a))

}

func test()  map[int64]error{
	 m := make(map[int64]error)
	var err = errors.New("1111")

	defer func(){
			m[1] = err
	}()

	return m
}