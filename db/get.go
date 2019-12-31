package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllStudent() ([]*Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{} //map[string]interface{}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []*Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return students, nil
}

func SearchStudent(req StudentSearchRequest) (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{}
	if req.FirstName != "" {
		filter["first_name"] = primitive.Regex{Pattern: req.FirstName, Options: "i"}
	}
	if req.LastName != "" {
		filter["last_name"] = primitive.Regex{Pattern: req.LastName, Options: "i"}
	}
	if req.Name != "" {
		filter["$or"] = []bson.M{
			bson.M{"first_name": primitive.Regex{Pattern: req.Name, Options: "i"}},
			bson.M{"last_name": primitive.Regex{Pattern: req.Name, Options: "i"}},
		}
	}
	if req.ClassName != "" {
		filter["class_name"] = primitive.Regex{Pattern: req.ClassName, Options: "i"}
	}
	if req.Email != "" {
		filter["email"] = primitive.Regex{Pattern: req.Email, Options: "i"}
	}

	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}
	return &students, nil
}

func GroupByLastName() (interface{}, error) {
	students, err := GetAllStudent()
	if err != nil {
		log.Printf("Get all student error: %v", err)
		return nil, err
	}
	groupLastName := make(map[string][]interface{})
	for _, value := range students {
		groupLastName[value.LastName] = append(groupLastName[value.LastName], value)
	}
	return groupLastName, err
}

func SearchStudentSimple(req StudentSearchRequest) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"first_name": req.FirstName, "last_name": req.LastName}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func GetStudentById(id int) (*Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var student Student
	fmt.Println("ID: ", id)
	err := Client.Database(DbName).Collection(ColName).FindOne(ctx, bson.M{"id": id}).Decode(&student)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &student, nil
}
