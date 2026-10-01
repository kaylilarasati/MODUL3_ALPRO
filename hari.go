package main
import "fmt"

func main() {
	var n, hari int
	fmt.Scan(&n)
	
	hari = (4 + n) % 7
	if hari == 0 {s
		hari = 7
	}
	fmt.Printf("%d", hari)
}
