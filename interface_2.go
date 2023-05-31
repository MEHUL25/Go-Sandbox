package main

import (
	"fmt"
)

// employee - contractor, full time
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name       string
	hourlyRate int
	noOFhours  int
}

type fulltime struct {
	name         string
	yearlySalary int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyRate * c.noOFhours
}

func (f fulltime) getName() string {
	return f.name
}

func (f fulltime) getSalary() int {
	return f.yearlySalary
}

func getDetails(e employee) {
	fmt.Println(e)
	fmt.Println(e.getName())
	fmt.Println(e.getSalary())
}

func main() {
	c := contractor{
		name:       "Mehul",
		hourlyRate: 10,
		noOFhours:  2000,
	}

	getDetails(c)
	fmt.Println()

	f := fulltime{
		name:         "Mehul2",
		yearlySalary: 2500000,
	}
	getDetails(f)
}
