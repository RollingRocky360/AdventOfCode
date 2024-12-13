package day13

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Button struct {
	dx float64
	dy float64
}

type Prize struct {
	x float64
	y float64
}

type Machine struct {
	buttonA Button
	buttonB Button
	prize   Prize
}

func matrixInverse(m [2][2]float64) ([2][2]float64, bool) {
	a, b, c, d := m[0][0], m[0][1], m[1][0], m[1][1]
	i := [2][2]float64{}

	det := a*d - b*c
	if det == 0 {
		return i, false
	}

	i[0][0] = d / det
	i[0][1] = -b / det
	i[1][0] = -c / det
	i[1][1] = a / det

	return i, true
}

func matrixMul(a [2][2]float64, b [2]float64) [2]float64 {
	prod := [2]float64{}

	prod[0] = a[0][0]*b[0] + a[0][1]*b[1]
	prod[1] = a[1][0]*b[0] + a[1][1]*b[1]

	return prod
}

func computeMinCost(add bool) float64 {
	machines := parseInput(add)

	cost := 0.0
	for _, machine := range machines {
		A := machine.buttonA
		B := machine.buttonB
		prize := machine.prize

		m := [2][2]float64{
			{A.dx, B.dx},
			{A.dy, B.dy},
		}

		inv, ok := matrixInverse(m)
		if !ok {
			continue
		}

		sol := matrixMul(inv, [2]float64{prize.x, prize.y})
		// fmt.Println(sol)
		if math.Abs(sol[0]-math.Round(sol[0])) < 0.0001 && math.Abs(sol[1]-math.Round(sol[1])) < 0.0001 {
			cost += 3*sol[0] + sol[1]
		}
	}

	return cost
}

func Part1() {
	cost := computeMinCost(false)
	log.Println("part 1:", cost)
}

func Part2() {
	cost := computeMinCost(true)
	log.Println("part 2:", cost)
}

func Run() {
	Part1()
	Part2()
}

func parseInput(add bool) []Machine {
	input, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	machines := []Machine{}

	for {
		scanner.Scan()
		buttonA := Button{}
		// fmt.Println(scanner.Text())
		fmt.Sscanf(scanner.Text(), "Button A: X+%f, Y+%f", &buttonA.dx, &buttonA.dy)

		scanner.Scan()
		buttonB := Button{}
		// fmt.Println(scanner.Text())
		fmt.Sscanf(scanner.Text(), "Button B: X+%f, Y+%f", &buttonB.dx, &buttonB.dy)

		scanner.Scan()
		prize := Prize{}
		// fmt.Println(scanner.Text())
		fmt.Sscanf(scanner.Text(), "Prize: X=%f, Y=%f", &prize.x, &prize.y)
		if add {
			prize.x += 10000000000000.0
			prize.y += 10000000000000.0
		}

		machines = append(machines, Machine{
			buttonA: buttonA,
			buttonB: buttonB,
			prize:   prize,
		})

		if !scanner.Scan() {
			break
		}
	}

	return machines
}
