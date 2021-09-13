package main

import "fmt"

func deretRekursif(i, n float64) float64 {
	divide := i / n
	if n == 1 {
		fmt.Printf("%v\n", i)
		return i
	} else {
		fmt.Printf("%v + ", divide)
		return i/n + deretRekursif(i, n-1)
	}
}

func deretRekursifD(i, n float64) float64 {
	divide := n / (n * n)
	if n == 1 {
		fmt.Printf("%v\n", i)
		return i
	} else {
		fmt.Printf("%v + ", divide)
		return n/(n*n) + deretRekursif(i, n-1)
	}
}

func deretRekursifE(i, n float64) float64 {
	divide := (n * n) / n
	if n == 1 {
		fmt.Printf("%v\n", i)
		return i
	} else {
		fmt.Printf("%v + ", divide)
		return (n*n)/n + deretRekursif(i, n-1)
	}
}

func main() {
	fmt.Printf("Hasil deret A: %v\n\n", deretRekursif(20, 4))
	fmt.Printf("Hasil deret B: %v\n\n", deretRekursif(100, 4))
	fmt.Printf("Hasil deret C: %v\n\n", deretRekursif(1, 4))
	fmt.Printf("Hasil deret D: %v\n\n", deretRekursifD(2, 4))
	fmt.Printf("Hasil deret E: %v\n\n", deretRekursifE(2, 4))

}
