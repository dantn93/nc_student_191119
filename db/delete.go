package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteStudentById(id int) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := make(map[string]interface{})
	err := Client.Database(DbName).Collection(ColName).FindOneAndDelete(ctx, bson.M{"id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteStudentByIdReq(id int) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := make(map[string]interface{})
	err := Client.Database(DbName).Collection(ColName).FindOneAndDelete(ctx, bson.M{"id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
