package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	var currentYear = 2024

	if age := currentYear - 1994; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	var score = 8
	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough
	case score < 5:
		fmt.Println("it is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	for i := 0; i < 3; i++ {
		fmt.Println("angka", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("angka", i)
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	greet("Arif Muspita")

	var SgreetMore = []string{"apa", "asd"}

	var result = greetMore(SgreetMore)

	fmt.Println(result)

	var area, circumference = calculate(15)

	fmt.Println("Area", area)
	fmt.Println("Circumference", circumference)

	profile("Arif Muspita", "Dada", "Paha", "Tulang", "Atas")

}

func greet(name string) {
	fmt.Printf("hello %s from function \n", name)
}

func greetMore(name []string) string {

	total := len(name)

	var result string

	var joinStr = strings.Join(name, ",")

	if total == 1 {
		result = fmt.Sprintf("Halo %s from function", joinStr)
	} else {
		result = fmt.Sprintf("Halo guys %s from function", joinStr)
	}

	return result

}

func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)

	var circumference = math.Pi * d

	return area, circumference
}

func profile(name string, fav ...string) {
	mergeFav := strings.Join(fav, ",")

	fmt.Printf("Hello %s from Variadic function \n", name)
	fmt.Println("I really like", mergeFav)
}
