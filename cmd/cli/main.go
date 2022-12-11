package main

import (
	"context"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	targetHost = "http://localhost:8080"
	targetURL  = "/ping"
)

// NOTE: Backoff algorithm experimentation
func exponentialBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// attemptNum always starts at zero but we want to start at 1 for multiplication
	attemptNum++

	if max <= min {
		// Unclear what to do here, or they are the same, so return min * attemptNum
		return min * time.Duration(attemptNum)
	}

	return time.Duration(math.Pow10(attemptNum-1)) * min
}

func getClient(retryable bool) *http.Client {

	timeoutDuration := 60 * time.Second
	transportOpts := &http.Transport{
		MaxConnsPerHost:     0, //NOTE: can act as rate limitter
		MaxIdleConnsPerHost: 5, // NOTE: can help to cache / poll available connection, because Each closed connection will be in TIME_WAIT state for two minutes, tying up that connection.
	}

	if !retryable {
		return &http.Client{
			Transport: transportOpts,
			Timeout:   timeoutDuration,
		}
	}

	logger := log.New(os.Stdout, "", log.Lmsgprefix)
	retryableClient := retryablehttp.NewClient()
	retryableClient.RetryMax = 5
	retryableClient.Backoff = retryablehttp.LinearJitterBackoff
	retryableClient.RetryWaitMin = 1 * time.Second
	retryableClient.RetryWaitMax = 20 * time.Second
	retryableClient.RequestLogHook = func(l retryablehttp.Logger, r *http.Request, i int) {
		l.Printf("[ATTEMPT]: %d | [TARGET URL]: %s", i, r.URL)
	}
	retryableClient.Logger = logger // NOTE: it use LevelledLogger interface (Debug, Info, Error, Warning)

	httpClient := retryableClient.StandardClient()
	httpClient.Timeout = timeoutDuration

	return httpClient
}

func main() {
	// NOTE: 1. Setup http client
	httpClient := getClient(true)

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	// NOTE: 3. Call with retryable approach
	ping(ctxWithCancel, httpClient, 10, quit)
	<-quit
}

func ping(ctx context.Context, httpClient *http.Client, numberOfCall int, done chan<- os.Signal) {

	var wg sync.WaitGroup

	for i := 0; i < numberOfCall; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetHost+targetURL, nil)
			if err != nil {
				log.Println(err)
				return
			}
			resp, err := httpClient.Do(req)
			if err != nil {
				log.Printf("http request error\n")
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Response status is not OK\n")
				return
			}

			log.Printf("Request successfull")
		}()
	}

	wg.Wait()
	close(done)
}
