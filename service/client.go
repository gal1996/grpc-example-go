package service

import (
	"context"
	"fmt"
	"github.com/TsuchiyaYugo/grpc-example-go/pb"
	"github.com/TsuchiyaYugo/grpc-example-go/pkg"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8080"
)

func PlayGame(ctx context.Context, handShapes int32) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctxWC, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	client := pb.NewRockPaperScissorsServiceClient(conn)
	playRequest := pb.PlayRequest{HandShapes: pkg.EncodeHandShapes(handShapes)}

	reply, err := client.PlayGame(ctxWC, &playRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	matchResult := reply.GetMatchResult()
	fmt.Println("***********************************")
	fmt.Printf("Your hand shapes: %s \n", matchResult.YourHAndShapes.String())
	fmt.Printf("Opponent hand shapes: %s \n", matchResult.OpponentHandShapes.String())
	fmt.Printf("Result: %s \n", matchResult.Result.String())
	fmt.Println("***********************************")
	fmt.Println()
}

func ReportMatchResults(ctx context.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctxWC, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	client := pb.NewRockPaperScissorsServiceClient(conn)

	reportRequest := pb.ReportRequest{}

	reply, err := client.ReportMatchResults(ctxWC, &reportRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	report := reply.GetReport()
	if len(report.MatchResults) == 0 {
		fmt.Println("***********************************")
		fmt.Println("There are no match results.")
		fmt.Println("***********************************")
		fmt.Println()
		return
	}

	fmt.Println("***********************************")
	for k, v := range report.MatchResults {
		fmt.Println(k + 1)
		fmt.Printf("Your hand shapes: %s \n", v.YourHAndShapes.String())
		fmt.Printf("Opponent hand shapes: %s \n", v.OpponentHandShapes.String())
		fmt.Printf("Result: %s \n", v.Result.String())
		fmt.Printf("Datetime of match: %s \n", v.CreateTime.AsTime().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format(time.ANSIC))
		fmt.Println()
	}

	fmt.Printf("Number of games: %d \n", reply.GetReport().NumberOfGames)
	fmt.Printf("Number of wins: %d \n", reply.GetReport().NumberOfWins)
	fmt.Println("***********************************")
	fmt.Println()
}
