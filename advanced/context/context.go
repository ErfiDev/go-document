package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// In Go servers, each incoming request is handled
	// in its own goroutine. Request handlers often start
	// additional goroutines to access backends such as
	// databases and RPC services. The set of goroutines
	// working on a request typically needs access to
	// request-specific values such as the identity of
	// the end user, authorization tokens, and the
	// request’s deadline. When a request is canceled or
	// times out, all the goroutines working on that
	// request should exit quickly so the system can
	// reclaim any resources they are using.
	// At Google, we developed a context package that
	// makes it easy to pass request-scoped values,
	// cancelation signals, and deadlines across API
	// boundaries to all the goroutines involved in
	// handling a request. The package is publicly
	// available as context. This article describes how
	// to use the package and provides a complete working
	// example.

	shortDuration := 1 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("after one milisecond")

	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
