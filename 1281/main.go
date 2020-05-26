package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInteger(scanner *bufio.Reader) int {
	input, _, _ := scanner.ReadLine()
	integer, _ := ConsoleToInt(input)
	return integer
}

func ConsoleToInt(inputText []byte) (int, error) {
	return strconv.Atoi(string(inputText))
}

func ParseProductInput(scanner *bufio.Reader) (string, string) {
	input, _, _ := scanner.ReadLine()
	parsedIput := strings.Split(string(input), " ")
	return string(parsedIput[0]), string(parsedIput[1])
}

func main() {
	priceTable := make(map[string]float64)
	var total float64

	scanner := bufio.NewReader(os.Stdin)
	cases := ReadInteger(scanner)

	for i := 0; i < cases; i++ {
		Availableproducts := ReadInteger(scanner)
		for product := 0; product < Availableproducts; product++ {
			name, strPrice := ParseProductInput(scanner)
			strPrice = strings.TrimSpace(strPrice)
			price, _ := strconv.ParseFloat(strPrice, 64)
			priceTable[name] = price
		}

		PurshedProducts := ReadInteger(scanner)
		for product := 0; product < PurshedProducts; product++ {
			name, strQtd := ParseProductInput(scanner)
			qtd, _ := strconv.Atoi(strQtd)
			total += float64(qtd) * priceTable[name]
		}

		fmt.Printf("R$ %.2f\n", total)
		total = 0
	}
}
