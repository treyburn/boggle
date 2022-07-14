package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/treyburn/boggle/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serviceAddress = "localhost:50051"

func main() {
	client := buildClinet()

	args := os.Args

	if len(args) != 3 {
		help()
		os.Exit(1)
	}

	switch strings.ToLower(args[1]) {
	case "solve":
		resp, err := client.Solve(context.Background(), &api.SolveRequest{Board: args[2]})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Solution ID: ", resp.GetId())
		os.Exit(0)
	case "solution":
		resp, err := client.Solution(context.Background(), &api.SolutionRequest{Id: args[2]})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Solution: ", resp.GetWords())
		os.Exit(0)
	default:
		help()
		os.Exit(1)
	}
}

func buildClinet() api.BoggleServiceClient {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return api.NewBoggleServiceClient(conn)
}

func help() {
	fmt.Println("Invalid arguments")
	fmt.Println("Valid arguments are \"Solve\" and \"Solution\"")
	fmt.Println("Example usage for solve: boggle-cli solve \"a,b,c;d,a,a;d,t,t\"")
	fmt.Println("Note: the board should have lines separated by ; and characters separated by comma")
	fmt.Println("Example usage for solution: boggle-cli solution \"the-id-you-got-from-solve\"")
}
