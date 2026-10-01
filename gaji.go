package main
import "fmt"

func main(){
	var Gp, Bl int

	fmt.Println("Masukkan gaji pokok:")
	fmt.Scan(&Gp)

	fmt.Println("Masukkan jumlah jam lembur: ")
	fmt.Scan(&Bl)

	totalGaji := Gp + (Bl * 45000) 
	potongan := Gp * 55 / 1000
	totalGaji = totalGaji - potongan
	fmt.Printf("Total gaji: %d\n", totalGaji)
}