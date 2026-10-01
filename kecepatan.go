package main
import "fmt"

func main() {
	var kecepatan, menit, jam, totalJarak int
	fmt.Scanf("%d", &kecepatan)
	totalJarak = 100 + 60 + 170
	menit = (totalJarak * 60) / kecepatan
	jam = menit / 60
	menit = menit % 60
	fmt.Printf("Waktu tempuh: %d jam %d menit\n", jam, menit)
}