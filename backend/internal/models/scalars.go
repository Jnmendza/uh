package models

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectID is a type alias to primitive.ObjectID.
// This allows gqlgen to bind this type to the GraphQL ID scalar seamlessly
// while remaining fully compatible with the BSON tags natively.
type ObjectID = primitive.ObjectID

// MarshalObjectID converts a models.ObjectID (primitive.ObjectID) to a string for GraphQL responses
func MarshalObjectID(id ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(fmt.Sprintf(`"%s"`, id.Hex())))
	})
}

// UnmarshalObjectID parses a string from a GraphQL request into a models.ObjectID
func UnmarshalObjectID(v interface{}) (ObjectID, error) {
	switch v := v.(type) {
	case string:
		return primitive.ObjectIDFromHex(v)
	case primitive.ObjectID:
		return v, nil
	default:
		return primitive.NilObjectID, fmt.Errorf("%T is not a valid ID", v)
	}
}
