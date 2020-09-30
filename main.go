package main

import "fmt"

func main() {
	var dia int
	var mes int

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch {
	case (mes == 12 && dia >= 22) || (mes == 1 && dia <= 20):
		fmt.Println("Capricornio")
	case (mes == 1 && dia >= 21) || (mes == 2 && dia <= 18):
		fmt.Println("Acuario")
	case (mes == 2 && dia >= 19) || (mes == 3 && dia <= 20):
		fmt.Println("Piscis")
	case (mes == 3 && dia >= 21) || (mes == 4 && dia <= 20):
		fmt.Println("Aries")
	case (mes == 4 && dia >= 21) || (mes == 5 && dia <= 20):
		fmt.Println("Tauro")
	case (mes == 5 && dia >= 21) || (mes == 6 && dia <= 21):
		fmt.Println("Géminis")
	case (mes == 6 && dia >= 22) || (mes == 7 && dia <= 22):
		fmt.Println("Cáncer")
	case (mes == 7 && dia >= 23) || (mes == 8 && dia <= 22):
		fmt.Println("Leo")
	case (mes == 8 && dia >= 23) || (mes == 9 && dia <= 22):
		fmt.Println("Virgo")
	case (mes == 9 && dia >= 23) || (mes == 10 && dia <= 22):
		fmt.Println("Libra")
	case (mes == 10 && dia >= 23) || (mes == 11 && dia <= 22):
		fmt.Println("Escorpio")
	case (mes == 11 && dia >= 23) || (mes == 12 && dia <= 21):
		fmt.Println("Sagitario")
	default:
		fmt.Println("Error de fecha")
	}

}
