package main

import (
	"fmt"
)

func main() {
	startPriceOld := 7500
	startPriceNew := 32000
	savingperMonth := 300
	percentLossByMonth := 1.55

	fmt.Println(NbMonths(startPriceOld, startPriceNew, savingperMonth, percentLossByMonth))
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	var result = [2]int{}

	startPriceOld1 := float64(startPriceOld)
	startPriceNew1 := float64(startPriceNew)

	var balanceofmoney int
	var moneysafe int

	moneysafe = savingperMonth

	var counter = [2]int{1, -1}

	switch {
	case startPriceOld1 == startPriceNew1:
		counter = [2]int{0, 0}
		result = counter
		break
	case startPriceOld1 > startPriceNew1:
		counter = [2]int{0, int(startPriceOld1 - startPriceNew1)}
		result = counter
		break
	default:
		var counter = [2]int{1, -1}
		for month := 0; counter[1] <= 0; month++ { //reverse (month = 6) - counter[1] <= 0 //TEST!
			var a, b float64

			if counter[0]%2 == 0 {
				percentLossByMonth += 0.5
			}

			a = (float64(startPriceOld1) - (float64(startPriceOld1)*percentLossByMonth)/100) //price for new car
			startPriceOld1 = a

			b = (float64(startPriceNew1) - (float64(startPriceNew1)*percentLossByMonth)/100) //price for old car
			startPriceNew1 = b

			counter[0] = month
			counter[1] = balanceofmoney

			counter123 := (a + float64(savingperMonth) - b)

			balanceofmoney = int(counter123)

			var fuckmyass float64
			fuckmyass = counter123 - float64(balanceofmoney)
			var abc123 float64
			abc123 = 0.5

			if fuckmyass >= abc123 {
				balanceofmoney = balanceofmoney + 1
			}

			savingperMonth += moneysafe

			if counter[1] >= 0 && counter[0] != 0 {
				result = counter
				break
			}
		}
	}

	return result
}
