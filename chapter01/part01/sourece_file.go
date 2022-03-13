package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println("This is simple go")

	//v1 := "123"
	//v2 := 123
	//v3 := "Have a nice day\n"
	//v4 := "abc"

	//fmt.Print(v1, v2, v3, v4)
	//fmt.Println()
	//fmt.Println(v1, v2, v3, v4)
	//fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	//fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	//myString := ""
	//arguments := os.Args
	//if len(arguments) == 1 {
	//	myString = "Please give me one argument!"
	//} else {
	//	myString = arguments[1]
	//}

	//io.WriteString(os.Stdout, "This is Standard output\n")
	//io.WriteString(os.Stderr, myString)
	//io.WriteString(os.Stderr, "\n")

	//f := os.Stdin
	//defer f.Close()

	//scanner := bufio.NewScanner(f)

	//// while 문 하고 비슷함
	//for scanner.Scan() {
	//	fmt.Println(">", scanner.Text())
	//}

	//if len(os.Args) == 1 {
	//	fmt.Println("Please give one or more floats.")
	//	os.Exit(1)
	//}
	//
	//arguments := os.Args
	//
	//min, _ := strconv.ParseFloat(arguments[1], 64)
	//max, _ := strconv.ParseFloat(arguments[1], 64)
	//
	//for i := 2; i < len(arguments); i++ {
	//	n, _ := strconv.ParseFloat(arguments[i], 64)
	//	if n < min {
	//		min = n
	//	}
	//	if n > max {
	//		max = n
	//	}
	//}
	//
	//fmt.Println("Min", min)
	//fmt.Println("Max", max)

	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)

	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	// err.Error 는 string 타입으로 변환해준다.
	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}

}

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}
