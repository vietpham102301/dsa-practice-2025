package main

import "fmt"

func towerOfHanoiSolve(count *int, disk int, src, des, sub string) {
	*count++
	if disk == 1 {
		printStep(disk, src, des)
		return
	}
	towerOfHanoiSolve(count, disk-1, src, sub, des)
	printStep(disk, src, des)
	towerOfHanoiSolve(count, disk-1, sub, des, src)
}

func printStep(disk int, from, to string) {
	fmt.Printf("move disk %d from %s to %s \n", disk, from, to)
}

func main() {
	n := 3
	rod1 := "A"
	rod2 := "B"
	rod3 := "C"

	count := 0
	towerOfHanoiSolve(&count, n, rod1, rod3, rod2) // || || || 2^n -1
	fmt.Printf("\nnumbef of steps: %d", count)
}
