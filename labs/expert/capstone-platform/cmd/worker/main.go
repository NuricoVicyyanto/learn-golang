package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"learn-golang/labs/expert/capstone-platform/internal/config"
	"learn-golang/labs/expert/capstone-platform/internal/delivery"
	"learn-golang/labs/expert/capstone-platform/internal/store"
)

func main() {
	cfg := config.Load()
	worker := delivery.NewWorker(store.NewMemoryStore(), cfg.WorkerCount)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	fmt.Println("worker started with", cfg.WorkerCount, "workers")
	if err := worker.Run(ctx); err != nil && err != context.Canceled {
		fmt.Fprintln(os.Stderr, "worker error:", err)
		os.Exit(1)
	}
}
