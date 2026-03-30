package main

import (
	"bufio"
	"calculator-rpc/calculator-rpc/calculatorpb"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("erro ao conectar: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite a expressão (ou 'exit'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		result, err := calcularExpressao(c, ctx, input)
		cancel()

		if err != nil {
			fmt.Println("erro:", err)
			continue
		}

		fmt.Printf("Resultado: %f\n\n", result)
	}
}

func calcularExpressao(c calculatorpb.CalculatorClient, ctx context.Context, expr string) (float64, error) {
	var nums []float64
	var ops []string

	current := ""

	for _, ch := range expr {
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			n, _ := strconv.ParseFloat(current, 64)
			nums = append(nums, n)
			ops = append(ops, string(ch))
			current = ""
		} else {
			current += string(ch)
		}
	}
	n, _ := strconv.ParseFloat(current, 64)
	nums = append(nums, n)

	result := nums[0]

	for i, op := range ops {
		next := nums[i+1]

		req := &calculatorpb.CalcRequest{A: result, B: next}

		switch op {
		case "+":
			res, _ := c.Add(ctx, req)
			result = res.Result
		case "-":
			res, _ := c.Subtract(ctx, req)
			result = res.Result
		case "*":
			res, _ := c.Multiply(ctx, req)
			result = res.Result
		case "/":
			res, _ := c.Divide(ctx, req)
			result = res.Result
		}
	}

	return result, nil
}
