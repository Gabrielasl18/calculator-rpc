package main

import (
	"calculator-rpc/calculator-rpc/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Println("Resultado gerado com sucesso")
	return &calculatorpb.CalcResponse{Result: in.A + in.B}, nil
}

func (s *server) Subtract(ctx context.Context, in *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Println("Resultado gerado com sucesso")
	return &calculatorpb.CalcResponse{Result: in.A - in.B}, nil
}

func (s *server) Multiply(ctx context.Context, in *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Println("Resultado gerado com sucesso")
	return &calculatorpb.CalcResponse{Result: in.A * in.B}, nil
}

func (s *server) Divide(ctx context.Context, in *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	if in.B == 0 {
		return nil, fmt.Errorf("divisão por zero")
	}
	fmt.Println("Resultado gerado com sucesso")
	return &calculatorpb.CalcResponse{Result: in.A / in.B}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("erro ao ouvir: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})

	log.Println("Servidor rodando na porta 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("erro ao iniciar: %v", err)
	}
}
