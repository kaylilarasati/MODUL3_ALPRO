package main
import "fmt"

func main() {
	var koin, emas, perak, tembaga int
	fmt.Scanf("%d", &koin)

	emas = koin / 9
	perak = (koin % 9) / 3
	tembaga = (koin % 9) % 3
	fmt.Printf("Jumlah koin emas: %d\n", emas)
	fmt.Printf("Jumlah koin perak: %d\n", perak)
	fmt.Printf("Jumlah koin tembaga: %d\n", tembaga)
}