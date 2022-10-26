package context

import (
	"context"
	"log"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store was cancelled")
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result // it goes empty string or response
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
}
