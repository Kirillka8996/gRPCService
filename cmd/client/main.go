package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "notes/pkg/api/notes/v1"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	client := desc.NewNotesClient(conn)

	ctx := context.Background()

	response, err := client.SaveNote(ctx, &desc.SaveNoteRequest{
		Info: &desc.NoteInfo{
			Title:   "t1",
			Content: "c1",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.NoteId)

}
