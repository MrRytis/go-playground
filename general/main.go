package main

import "fmt"

func AddNumbers(a, b int64) int64 {
	return a + b
}

type Data struct {
	Number int64
}

type OtherData struct {
	Chars []string
}

func main() {
	data := Data{
		Number: 10,
	}

	data.Number = AddNumbers(1, 2)

	fmt.Printf("number: %d\n", data.Number)

	d := &data // pointer to data
	d.Number = AddNumbers(10, 5)

	fmt.Printf("number: %d\n", data.Number)

	fmt.Println("----")

	otherData := OtherData{
		Chars: []string{"a", "b", "c"},
	}
	PrintChars(otherData)
	fmt.Println("-- After appended --")
	otherData.Chars = append(otherData.Chars, "d")
	PrintChars(otherData)

	fmt.Println("----")

	_, err := CheckEven(10)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	_, err = CheckEven(11)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	fmt.Println("----")

}

func PrintChars(otherData OtherData) {
	for _, char := range otherData.Chars {
		fmt.Printf("char: %s\n", char)
	}
}

func CheckEven(number int64) (bool, error) {
	if number%2 == 0 {
		return true, nil
	}

	return false, fmt.Errorf("number %d is not even", number)
}
