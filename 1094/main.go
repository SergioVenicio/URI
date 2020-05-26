/*
Maria acabou de iniciar seu curso de graduação na faculdade de medicina e precisa de sua ajuda para organizar os experimentos de um laboratório o qual ela é responsável.
Ela quer saber no final do ano, quantas cobaias foram utilizadas no laboratório e o percentual de cada tipo de cobaia utilizada.
Este laboratório em especial utiliza três tipos de cobaias:
	sapos, ratos e coelhos.
Para obter estas informações, ela sabe exatamente o número de experimentos que foram realizados, o tipo de cobaia utilizada e a quantidade de cobaias utilizadas em cada experimento.
Entrada
A primeira linha de entrada contém um valor inteiro N que indica os vários casos de teste que vem a seguir.
Cada caso de teste contém um inteiro Quantia (1 ≤ Quantia ≤ 15) que representa a quantidade de cobaias utilizadas e um caractere
Tipo ('C', 'R' ou 'S'), indicando o tipo de cobaia (R:Rato S:Sapo C:Coelho).
Saída
Apresente o total de cobaias utilizadas,
o total de cada tipo de cobaia utilizada e o percentual de cada uma em relação ao total de cobaias utilizadas,
sendo que o percentual deve ser apresentado com dois dígitos após o ponto.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(input string) (string, int) {
	parsedInput := strings.Split(input, " ")
	value, _ := strconv.Atoi(parsedInput[0])
	specie := parsedInput[1]

	return specie, value
}

func CalcPorcent(specieValue int, total int) float32 {
	return (float32(specieValue) / float32(total)) * 100
}

func main() {
	var cases int
	inputs := make(map[string]int)
	scanner := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &cases)

	for i := 0; i < cases; i++ {

		input, _, _ := scanner.ReadLine()
		specie, value := ParseInput(string(input))

		if _, hasSpecie := inputs[specie]; hasSpecie {
			inputs[specie] += value
		} else {
			inputs[specie] = value
		}
	}

	total := inputs["C"] + inputs["R"] + inputs["S"]

	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Printf("Total de coelhos: %d\n", inputs["C"])
	fmt.Printf("Total de ratos: %d\n", inputs["R"])
	fmt.Printf("Total de sapos: %d\n", inputs["S"])
	fmt.Printf("Percentual de coelhos: %.2f %%\n", CalcPorcent(inputs["C"], total))
	fmt.Printf("Percentual de ratos: %.2f %%\n", CalcPorcent(inputs["R"], total))
	fmt.Printf("Percentual de sapos: %.2f %%\n", CalcPorcent(inputs["S"], total))
}
