package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	desc "notes/pkg/api/notes/v1"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := desc.NewNotesClient(conn)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-auth", "user1")
	req := &desc.SaveNoteRequest{
		Info: &desc.NoteInfo{
			Title:   "tdddddddddd",
			Content: "cdddddddddd",
		},
	}

	response, err := client.SaveNote(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.NoteId)

}
