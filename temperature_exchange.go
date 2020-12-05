package main  // 指明main包

import "fmt"  // 导入fmt包

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n%", freezingF, fToC(freezingF)) 
	fmt.Println("%gF = %g\n", boilingF, fToC(boilingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}