package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToPrimitiveId(Id string) primitive.ObjectID {
	result, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return primitive.NilObjectID
	}

	return result
}
