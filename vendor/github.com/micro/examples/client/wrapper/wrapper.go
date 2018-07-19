package main

import (
	"fmt"
	"time"

	example "github.com/micro/examples/server/proto/example"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"context"
)

// wrapper example code

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// trace wrapper attaches a unique trace ID - timestamp
type traceWrapper struct {
	client.Client
}

func (t *traceWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	ctx = metadata.NewContext(ctx, map[string]string{
		"X-Trace-Id": fmt.Sprintf("%d", time.Now().Unix()),
	})
	return t.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

// Implements client.Wrapper as traceWrapper
func traceWrap(c client.Client) client.Client {
	return &traceWrapper{c}
}

func metricsWrap(cf client.CallFunc) client.CallFunc {
	return func(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
		t := time.Now()
		err := cf(ctx, addr, req, rsp, opts)
		fmt.Printf("[Metrics Wrapper] called: %s %s.%s duration: %v\n", addr, req.Service(), req.Method(), time.Since(t))
		return err
	}
}

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	fmt.Println("\n--- Log Wrapper example ---")

	// Wrap the default client
	client.DefaultClient = logWrap(client.DefaultClient)

	call(0)

	fmt.Println("\n--- Log+Trace Wrapper example ---")

	// Wrap using client.Wrap option
	client.DefaultClient = client.NewClient(
		client.Wrap(traceWrap),
		client.Wrap(logWrap),
	)

	call(1)

	fmt.Println("\n--- Metrics Wrapper example ---")

	// Wrap using client.Wrap option
	client.DefaultClient = client.NewClient(
		client.WrapCall(metricsWrap),
	)

	call(2)
}
