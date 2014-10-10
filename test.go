package main

import (
	"fmt"
	"currency_convert"
)

func main() {
	//	currency_convert.Convert(SOURCE_CURRENCY, TARGET_CURRENCY, AMOUNT);
	
	convertedRate := currency_convert.Convert("EUR", "INR", 50)
	fmt.Println(convertedRate)
}