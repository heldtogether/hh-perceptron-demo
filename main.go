package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/heldtogether/perceptron"
)

var (
	rounds         int
	generator      string
	numberInputs   = 2
	p              *perceptron.Perceptron
	input          []int
	expectedOutput int
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.IntVar(&rounds, "rounds", 100, "Number of rounds of training to perform.")
	flag.StringVar(&generator, "generator", "and", "Type of training data generator.")
	flag.Parse()
}

func main() {
	p = perceptron.New(numberInputs)

	fmt.Printf("Before training using the \"%s\" generator.\n", generator)
	printResults(p)

	for i := 0; i < rounds; i++ {
		switch generator {
		case "and":
			input, expectedOutput = generateAnd()
		case "nand":
			input, expectedOutput = generateNand()
		case "or":
			input, expectedOutput = generateOr()
		default:
			panic("Generator not allowed")
		}
		p.Train(input, expectedOutput)
	}

	fmt.Printf("After training with %d number of rounds using the \"%s\" generator.\n", rounds, generator)
	printResults(p)
}

func printResults(p *perceptron.Perceptron) {
	input = []int{0, 0}
	fmt.Printf("Try Input: %v. Output: %d. \n", input, p.Activates(input))

	input = []int{0, 1}
	fmt.Printf("Try Input: %v. Output: %d. \n", input, p.Activates(input))

	input = []int{1, 0}
	fmt.Printf("Try Input: %v. Output: %d. \n", input, p.Activates(input))

	input = []int{1, 1}
	fmt.Printf("Try Input: %v. Output: %d. \n", input, p.Activates(input))
}

func generateAnd() ([]int, int) {
	i1 := rand.Intn(2)
	i2 := rand.Intn(2)

	b1 := false
	if i1 == 1 {
		b1 = true
	}

	b2 := false
	if i2 == 1 {
		b2 = true
	}

	o := 0
	if b1 && b2 {
		o = 1
	}

	return []int{i1, i2}, o
}

func generateNand() ([]int, int) {
	i1 := rand.Intn(2)
	i2 := rand.Intn(2)

	b1 := false
	if i1 == 1 {
		b1 = true
	}

	b2 := false
	if i2 == 1 {
		b2 = true
	}

	o := 0
	if !(b1 && b2) {
		o = 1
	}

	return []int{i1, i2}, o
}

func generateOr() ([]int, int) {
	i1 := rand.Intn(2)
	i2 := rand.Intn(2)

	b1 := false
	if i1 == 1 {
		b1 = true
	}

	b2 := false
	if i2 == 1 {
		b2 = true
	}

	o := 0
	if b1 || b2 {
		o = 1
	}

	return []int{i1, i2}, o
}
