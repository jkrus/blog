package main

import (
	"context"
	"log"

	"blog/models/grpc/proto"
)

func AddNote() {
	ctx := context.Background()
	defer terminate()
	if err := connect(); err != nil {
		log.Println(err)
	}

	client := proto.NewNoteServiceClient(con)

	req := proto.NoteDTO{Body: "grpc-body", Header: "grpc-header", Tags: []string{"grpc-t1", "grpc-t2", "grpc-t3"}}

	resp, err := client.AddNote(ctx, &req)

	if err != nil {
		log.Println("ERR = ", err)
	}

	log.Println("RESP = ", resp)

}

func GetNotes() {
	ctx := context.Background()
	defer terminate()
	if err := connect(); err != nil {
		log.Println(err)
	}

	client := proto.NewNoteServiceClient(con)

	req := proto.EmptyRequest{}

	resp, err := client.GetNotes(ctx, &req)

	if err != nil {
		log.Println("ERR = ", err)
	}

	log.Println("RESP = ", resp.GetNotes())

}
