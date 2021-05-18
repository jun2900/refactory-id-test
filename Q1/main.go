package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	var restoName, dateOfPrint, cashierName, item string
	var price, total int = 0, 0
	itemPrice := make(map[string]int)

	reader := bufio.NewReader(os.Stdin)
	p := message.NewPrinter(language.Indonesian)

	// Taking Input
	fmt.Print("Input resto name: ")
	restoName, _ = reader.ReadString('\n')
	fmt.Print("Input date of print: ")
	dateOfPrint, _ = reader.ReadString('\n')
	fmt.Print("Input cashier name: ")
	cashierName, _ = reader.ReadString('\n')

	for {
		fmt.Print("Input item (type exit to quit): ")
		fmt.Scanf("%s", &item)
		if item == "exit" {
			break
		}
		fmt.Print("Input price: ")
		fmt.Scanf("%d", &price)
		total += price
		itemPrice[item] = price
	}

	// Display output
	fmt.Printf("%15v\n", restoName)
	fmt.Printf("Tanggal :%24s\n", dateOfPrint)
	fmt.Printf("Nama Kasir :%18s\n", cashierName)
	fmt.Println("================================")
	for key, value := range itemPrice {
		p.Print(key, "...................Rp", value, "\n")
	}

	p.Print("\nTotal...................Rp", total, "\n")
}
