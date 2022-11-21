package greeter

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/keepalive"
	"log"
	"testing"
	"time"
)

func TestGrpcClient(t *testing.T) {
	conn, err := grpc.Dial("192.168.0.103:8088", grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			PermitWithoutStream: true,
			Timeout:             1 * time.Second,
			Time:                10 * time.Second,
		}),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s","MethodConfig": [{"Name": [{"Service": "greeter.GreeterService"}], "RetryPolicy": {"MaxAttempts":2, "InitialBackoff": "0.1s", "MaxBackoff": "1s", "BackoffMultiplier": 2.0, "RetryableStatusCodes": ["UNAVAILABLE"]}}]}`, roundrobin.Name)),
		grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("---1-----")
	if resp, err := NewGreeterServiceClient(conn).SayHelloWord(context.Background(), &HelloRequest{
		Name: "test",
	}); err != nil {
		t.Fatal(err)
	} else {
		t.Log("vvvvvv", resp.Time)
	}
	go func() {
		for {
			log.Println(conn.GetState().String())
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(8 * time.Second)
	log.Println("--------start------")
	start := time.Now()
	ctxTime, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println(conn.GetState().String())
	log.Println(conn.GetState().String())
	if resp, err := NewGreeterServiceClient(conn).SayHelloWord(ctxTime, &HelloRequest{
		Name: "test",
	}, grpc.WaitForReady(true)); err != nil {
		t.Log("errrrrrr", time.Now().Sub(start))
		t.Fatal(err)
	} else {
		t.Log("xxxxxx", resp.Time, time.Now().Sub(start))
	}
}
