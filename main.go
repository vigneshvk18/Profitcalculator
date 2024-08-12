package main

import "fmt"

func main() {
	
	revenue := getUserInput("revenue: ")
	expenses := getUserInput("expenses: ")
	taxRate := getUserInput("tax rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

	
func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64){
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
	

func getUserInput(infotext string) float64{
	var getinput float64
	fmt.Print(infotext)
	fmt.Scan(&getinput)
	return getinput
}
