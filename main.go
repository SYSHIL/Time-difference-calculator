package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// var input2 string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the first date and time (YYYY-MM-DD HH:MM:SS): ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSuffix(input1, "\n")

	// fmt.Println(input1)
	fmt.Println("Enter the first date and time (YYYY-MM-DD HH:MM:SS): ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSuffix(input2, "\n")

	time1, _ := time.Parse(time.DateTime, input1)
	time2, _ := time.Parse(time.DateTime, input2)
	duration := time1.Sub(time2).Abs()
	hoursDifference := duration.Hours()
	fmt.Println(hoursDifference)
}
