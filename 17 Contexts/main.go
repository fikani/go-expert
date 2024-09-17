package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	// Book a hotel
	select {
	case <-ctx.Done():
		println("book is cancelled")
	case <-time.After(5 * time.Second):
		println("hotel booked")
	}
}
