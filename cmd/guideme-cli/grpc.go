package main

import (
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
	cli "guide.me/gen/grpc/cli/guideme"
)

func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server at %s: %v\n", host, err)
	}
	return cli.ParseEndpoint(conn)
}
