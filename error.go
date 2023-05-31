package main

import (
	"fmt"
)

func sendToBoth(customer, spouse string) (int, error) {
	custCost, err := calculateCost(customer)
	if err != nil {
		return 0, err
	}
	spouseCost, err := calculateCost(spouse)
	if err != nil {
		return 0, err
	}
	return spouseCost + custCost, nil
}

func calculateCost(message string) (int, error) {
	if len(message) > 25 {
		return 0, fmt.Errorf("Error due to length exceeded")
	}
	return len(message), nil
}

func test(customer, spouse string) {
	cost, err := sendToBoth(customer, spouse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cost)
}

func main() {
	test(
		"customer  - sfnvdb",
		"spouse - Hi",
	)
	test(
		"customer - fgdgggggfdddddddddddddddddddddddddddddddffffffffffffffffffffffffffffffffffffff",
		"spouse - dfsgfsgs",
	)

}
