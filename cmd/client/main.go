package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/TsuchiyaYugo/grpc-example-go/service"
	"os"
	"strconv"
)

func main() {
	ctx := context.Background()
	fmt.Println("start Rock-paper-scissors game.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1: play game")
		fmt.Println("2: show match results")
		fmt.Println("3: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		mode := scanner.Text()

		switch mode {
		case "1":
			fmt.Println("Please enter Rock, Paper, or Scissors.")
			fmt.Println("1: Rock")
			fmt.Println("2: Paper")
			fmt.Println("3: Scissors")
			fmt.Print("please enter >")

			scanner.Scan()
			inputHand := scanner.Text()
			switch inputHand {
			case "1", "2", "3":
				handShapes, _ := strconv.Atoi(inputHand)
				service.PlayGame(ctx, int32(handShapes))
			default:
				fmt.Println("Invalid command")
				continue
			}
		case "2":
			fmt.Println("Here are your match results.")
			service.ReportMatchResults(ctx)
			continue
		case "3":
			fmt.Println("bye.")
			// breakはだめ。switchを抜けるだけでループを抜けない
			goto M
		default:
			fmt.Println("Invalid command.")
			continue
		}
	}
M:
}
