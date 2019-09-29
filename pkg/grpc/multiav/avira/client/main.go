// Copyright 2019 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/saferwall/saferwall/pkg/grpc/multiav"
	pb "github.com/saferwall/saferwall/pkg/grpc/multiav/avira/proto"
	"google.golang.org/grpc"
	"log"
)

// ScanFile scans file
func ScanFile(client pb.AviraScannerClient, path string) (multiav.ScanResult, error) {
	scanFileRequest := &pb.ScanFileRequest{Filepath: path}
	res, err := client.ScanFile(context.Background(), scanFileRequest)
	if err != nil {
		return multiav.ScanResult{}, err
	}

	return multiav.ScanResult{
		Output:   res.Output,
		Infected: res.Infected,
		Update:   res.Update,
	}, nil
}

func main() {
	serverAddr, opts, filePath := multiav.ParseFlags()
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAviraScannerClient(conn)

	// ScanFile
	res, err := ScanFile(client, filePath)
	if err != nil {
		log.Fatalf("fail to scanfile: %v", err)
	}
	log.Println(res)
}