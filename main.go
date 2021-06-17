package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/m3hm3t/go-retry-library/pkg/retry"
	"time"
)

var count int

func EmulateTransientError(ctx context.Context) (string, error) {
	count++

	if count <= 3 {
		return "intentional fail", errors.New("error")
	} else {
		return "success", nil
	}
}

func main() {
	r := retry.Retry(EmulateTransientError, 5, 2*time.Second)

	res, err := r(context.Background())

	fmt.Println(res, err)
}