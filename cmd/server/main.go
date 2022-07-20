package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/treyburn/boggle/api"
	"github.com/treyburn/boggle/repository"
	"github.com/treyburn/boggle/rpc"
	"github.com/treyburn/boggle/solver"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = 50051

	useCustomerSolver = true
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("err creating logger: ", err)
	}
	dictionary, err := filepath.Abs("./assets/3_letter_dictionary.txt")
	if err != nil {
		logger.Error("creating filepath", zap.Error(err))
	}
	repo := repository.NewInMemory()

	var sol solver.Solver
	if !useCustomerSolver {
		sol = solver.NewOffTheShelf(dictionary, repo, logger)
	} else {
		sol, err = buildCustomSolver(dictionary, repo, logger)
		if err != nil {
			logger.Error("creating solver", zap.Error(err))
		}
	}

	service := rpc.NewBoggleService(repo, sol, logger)
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		logger.Error("creating listener", zap.Error(err))
	}

	api.RegisterBoggleServiceServer(server, service)

	logger.Info(fmt.Sprintf("serving on port %v", port))

	if err := server.Serve(listener); err != nil {
		logger.Fatal("server shutting down", zap.Error(err))
	}
}
