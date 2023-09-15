package services

import (
	"context"
	"errors"
	"hng-stage2/resource"
	"fmt"
	"time"

	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type HumanServiceImpl struct {
	humanCollection *mongo.Collection
	ctx             context.Context
}

func NewHumanService(humanCollection  *mongo.Collection, ctx context.Context ) HumanService{
	return &HumanServiceImpl{humanCollection, ctx}
}

func (h *HumanServiceImpl) CreateHuman(human *resource.Human) (*resource.DbHuman, error)  {
	fmt.Println("creating new human name")
	human.CreatedAT = time.Now()
	human.UpdatedAt = time.Now()

	res, err := h.humanCollection.InsertOne(h.ctx, human)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000{
			return nil, errors.New("human name already exist")
		}

		return nil, err
	}
		// adding a unique constraint  on the name field so that no two human will have the same name
		opt := options.Index()
		opt.SetUnique(true)

		index := mongo.IndexModel{Keys: bson.M{"user_name": 1}, Options: opt}
		if _, err := h.humanCollection.Indexes().CreateOne(h.ctx, index); err != nil {
			return nil, errors.New("could not create index for name")
		}

		var newHuman *resource.DbHuman
		filter := bson.M{"_id": res.InsertedID}
		if err = h.humanCollection.FindOne(h.ctx, filter).Decode(&newHuman); err != nil{
			return nil, err
		}

		return newHuman, nil
}

func (h *HumanServiceImpl) GetAllHuman() ([]*resource.DbHuman, error) {
	fmt.Println("getting all human names")
	filter := bson.M{}

	cursor, err := h.humanCollection.Find(h.ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(h.ctx)

	var humans []*resource.DbHuman

	for cursor.Next(h.ctx){
		human := &resource.DbHuman{}
		err := cursor.Decode(human)

		if err != nil {
			return nil, err
		}
	
		humans = append(humans, human)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(humans) == 0 {
		return []*resource.DbHuman{}, nil
	}

	return humans, nil

	
}


func (h *HumanServiceImpl) GetHumanbyID(id string) (*resource.DbHuman, error) {
	fmt.Println("getting all human name for the id: ", id)
	huId, _ := primitive.ObjectIDFromHex(id)


	filter := bson.M{"_id": huId}

	var human *resource.DbHuman

	if err := h.humanCollection.FindOne(h.ctx, filter).Decode(&human); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return human, nil


}

func (h *HumanServiceImpl) UpdateHuman(id string, data *resource.Human) (*resource.DbHuman, error)  {
	fmt.Println("update human name for id: ", id)
	dataArr, err := bson.Marshal(data)
	if err != nil {
		return nil, errors.New("Unable to Marshal the post object to bson")
	}
	var doc *bson.D
	err = bson.Unmarshal(dataArr, &doc)
	if err != nil {
		return nil, err
	}

	huID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: huID}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := h.humanCollection.FindOneAndUpdate(h.ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *resource.DbHuman
	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no human with that Id exists")
	}

	return updatedPost, nil
}

func (h *HumanServiceImpl) DeleteHuman(id string) error {
	fmt.Println("delete human name oof id: ", id)
		huId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": huId}
	
		res, err := h.humanCollection.DeleteOne(h.ctx, filter)
		if err != nil {
			return err
		}
	
		if res.DeletedCount == 0 {
			return errors.New("no document with that Id exists")
		}
	
		return nil

}
