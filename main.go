package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Operation struct {
	Operation string  `json:"operation"`
	UnitCost  float32 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

func main() {
	input, _ := readInput()
	fmt.Println(input)

	res := strings.Split(input, "][")
	fmt.Println(res)
	for i := range res {
		if len(res) > 1 {
			if i%2 == 0 {
				res[i] += "]"
			} else {

				res[i] = "[" + res[i]
			}
		}
		fmt.Println(i, res[i])
		var operations []Operation
		err := json.Unmarshal([]byte(res[i]), &operations)
		if err != nil {
			log.Fatalf("Error unmarshalling operations: %v", err)
		}
		processOperations(operations)
	}

}

type Tax struct {
	Tax float32 `json:"tax"`
}

const (
	BUY  = "buy"
	SELL = "sell"
)

func processOperations(operations []Operation) {
	var weightedAverage, loss float32
	var totalStockNumber int

	for _, operation := range operations {
		var tax float32
		if operation.Operation == BUY {
			newTotalStockNumber := totalStockNumber + operation.Quantity
			newWeightedAverage := ((float32(totalStockNumber) * weightedAverage) + (float32(operation.Quantity) * operation.UnitCost)) / float32(newTotalStockNumber)
			weightedAverage = newWeightedAverage
			totalStockNumber = newTotalStockNumber
		} else {
			totalStockNumber -= operation.Quantity
			salesPrice := float32(operation.Quantity) * operation.UnitCost
			averagePrice := float32(operation.Quantity) * weightedAverage

			if salesPrice < averagePrice {
				loss += averagePrice - salesPrice
			} else {
				rawProfit := salesPrice - averagePrice
				profit := rawProfit - loss
				if profit > 0 && salesPrice > 20000 {
					tax = profit * 0.2
				}

				if rawProfit > loss {
					loss = 0
				} else {
					loss -= rawProfit
				}
			}
		}
		fmt.Println(operation.Operation, weightedAverage, totalStockNumber, tax)
	}

}

func readInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		input = append(input, line)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(input, ""), nil
}
