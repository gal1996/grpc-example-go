package service

import (
	"context"
	"fmt"
	"github.com/TsuchiyaYugo/grpc-example-go/pb"
	"github.com/TsuchiyaYugo/grpc-example-go/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RockPaperScissorsService struct {
	numberOfGames int32
	numberOfWins  int32
	matchResults  []*pb.MatchResult
}

func NewRockPaperScissorsService() *RockPaperScissorsService {
	return &RockPaperScissorsService{
		numberOfGames: 0,
		numberOfWins:  0,
		matchResults:  make([]*pb.MatchResult, 0),
	}
}

func (s *RockPaperScissorsService) PlayGame(ctx context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	log.Println("start PlayGame")
	if req.HandShapes == pb.HandShapes_HAND_SHAPES_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	opponentHandShapes := pkg.EncodeHandShapes(int32(rand.Intn(3) + 1))

	var result pb.Result
	if req.HandShapes == opponentHandShapes {
		result = pb.Result_DRAW
	} else if (req.HandShapes.Number()-opponentHandShapes.Number()+3)%3 == 1 {
		result = pb.Result_WIN
	} else {
		result = pb.Result_LOSE
	}

	now := time.Now()
	matchResult := &pb.MatchResult{
		YourHAndShapes:     req.HandShapes,
		OpponentHandShapes: opponentHandShapes,
		Result:             result,
		CreateTime: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	s.numberOfGames = s.numberOfGames + 1
	if result == pb.Result_WIN {
		s.numberOfWins = s.numberOfWins + 1
	}
	s.matchResults = append(s.matchResults, matchResult)

	return &pb.PlayResponse{
		MatchResult: matchResult,
	}, nil
}

func (s *RockPaperScissorsService) ReportMatchResults(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	fmt.Println("start ReportMatchResults")
	return &pb.ReportResponse{
		Report: &pb.Report{
			NumberOfGames: s.numberOfGames,
			NumberOfWins:  s.numberOfWins,
			MatchResults:  s.matchResults,
		},
	}, nil
}
