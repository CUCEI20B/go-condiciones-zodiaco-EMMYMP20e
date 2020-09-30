package main

import "fmt"

func main() {
	var dia int
	var mes int

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch {
	case (mes == 12 && dia >= 22) || (mes == 1 && dia <= 20):
		fmt.Println("capricornio")
	case (mes == 1 && dia >= 21) || (mes == 2 && dia <= 18):
		fmt.Println("acuario")
	case (mes == 2 && dia >= 19) || (mes == 3 && dia <= 20):
		fmt.Println("piscis")
	case (mes == 3 && dia >= 21) || (mes == 4 && dia <= 20):
		fmt.Println("aries")
	case (mes == 4 && dia >= 21) || (mes == 5 && dia <= 20):
		fmt.Println("tauro")
	case (mes == 5 && dia >= 21) || (mes == 6 && dia <= 21):
		fmt.Println("geminis")
	case (mes == 6 && dia >= 22) || (mes == 7 && dia <= 22):
		fmt.Println("cancer")
	case (mes == 7 && dia >= 23) || (mes == 8 && dia <= 22):
		fmt.Println("leo")
	case (mes == 8 && dia >= 23) || (mes == 9 && dia <= 22):
		fmt.Println("virgo")
	case (mes == 9 && dia >= 23) || (mes == 10 && dia <= 22):
		fmt.Println("libra")
	case (mes == 10 && dia >= 23) || (mes == 11 && dia <= 22):
		fmt.Println("escorpio")
	case (mes == 11 && dia >= 23) || (mes == 12 && dia <= 21):
		fmt.Println("sagitario")
	}

}
