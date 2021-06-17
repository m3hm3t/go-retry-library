package retry

import (
	"context"
	"github.com/m3hm3t/go-retry-library/pkg/retry/effector"
	"log"
	"time"
)

func Retry(effector effector.Effector, retries int, delay time.Duration) effector.Effector {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			log.Printf("Attempt %d failed; retrying in %v", r + 1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}