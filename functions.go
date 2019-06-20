package functions

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	pb "google.golang.org/genproto/googleapis/firestore/v1"
)

// FirestoreEvent is received as an argument in functions invoked by Firestore events.
type FirestoreEvent struct {
	OldValue   *Document
	Value      *Document
	UpdateMask *pb.DocumentMask
}

// Document wraps the native Firestore document to help with the JSON unmarshal
// in Cloud Functions.
type Document pb.Document

// UnmarshalJSON implements the custom unmarshal from JSON of the document
// using jsonpb instead of the normal JSON procedures.
func (doc *Document) UnmarshalJSON(content []byte) error {
	return jsonpb.Unmarshal(bytes.NewReader(content), (*pb.Document)(doc))
}
