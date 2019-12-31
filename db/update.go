package db

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetNextSequenceValue(sequenceName string) (id int, err error) {
	collection := Client.Database(DbName).Collection("counters")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	doc := make(map[string]interface{})
	collection.FindOneAndUpdate(ctx,
		bson.M{"id": sequenceName},
		bson.M{"$inc": bson.M{"sequence_value": 1}},
	).Decode(&doc)

	if doc["sequence_value"] == nil {
		_, err := collection.InsertOne(ctx, bson.M{"id": sequenceName, "sequence_value": 0})
		if err != nil {
			return 0, err
		}
		return 0, nil
	}

	counters := Counters{}
	byteCounter, err := json.Marshal(doc)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(byteCounter, &counters)
	return counters.SequenceValue + 1, nil
}

func AddStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id, err := GetNextSequenceValue("student_id")
	if err != nil {
		return nil, err
	}
	student.ID = id
	if err != nil {
		return nil, err
	}
	res, err := collection.InsertOne(ctx, student)
	return res, err
}

func UpdateStudent(student *StudentUpdateRequest) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": student.Email}
	update := bson.M{"$set": student}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
