package functions

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	pb "google.golang.org/genproto/googleapis/firestore/v1"
)

type FirestoreEvent struct {
	OldValue   *Document
	Value      *Document
	UpdateMask *pb.DocumentMask
}

type Document pb.Document

func (doc *Document) UnmarshalJSON(content []byte) error {
	return jsonpb.Unmarshal(bytes.NewReader(content), (*pb.Document)(doc))
}
