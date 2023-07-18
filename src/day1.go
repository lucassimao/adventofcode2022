package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Panic("Error getting current directory:", err)
	}

	filepath := filepath.Join(currentDir, "inputs", "day1.txt")

	file, err := os.Open(filepath)
	if err != nil {
		log.Panic("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	biggestSumOfCaloriesPerElf := -1
	sumOfCaloriesPerElf := 0
	allCaloriesSumPerElf := sort.IntSlice{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) > 0 {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Panic("Could not convert line to int:", err)
			}

			sumOfCaloriesPerElf += calories
		} else {
			if sumOfCaloriesPerElf > biggestSumOfCaloriesPerElf {
				biggestSumOfCaloriesPerElf = sumOfCaloriesPerElf
			}
			allCaloriesSumPerElf = append(allCaloriesSumPerElf, sumOfCaloriesPerElf)
			sumOfCaloriesPerElf = 0
		}

	}

	allCaloriesSumPerElf.Sort()
	fmt.Printf("Biggest Total Of Calories: %d\n", biggestSumOfCaloriesPerElf)

	top3CaloriesSum := allCaloriesSumPerElf[len(allCaloriesSumPerElf)-3:]
	fmt.Printf("Top 3 Total of Calories: %v\n", top3CaloriesSum)

	top3CaloriesSumTotal := 0
	for _, total := range top3CaloriesSum {
		top3CaloriesSumTotal += total
	}
	fmt.Printf("Sum of Top 3 Total of Calories: %v\n", top3CaloriesSumTotal)

}
