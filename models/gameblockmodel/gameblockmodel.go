package gameblockmodel

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameOfLifeUnit struct {
	Alive bool `json:"alive"`
}

type GameBlockDoc struct {
	Location string              `bson:"location"`
	Units    [][]*GameOfLifeUnit `bson:"units"`
}

var collection *mongo.Collection

func getBlockLocation(rowIdx int, colIdx int) string {
	return fmt.Sprintf("%v-%v", colIdx, rowIdx)
}

func CreateOrUpdateGameBlocks(rowIdx int, colIdx int, units [][]*GameOfLifeUnit) error {
	location := getBlockLocation(rowIdx, colIdx)
	filter := bson.D{{"location", location}}
	update := bson.D{{"$set", GameBlockDoc{
		location,
		units,
	}}}
	options := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(
		context.TODO(),
		filter,
		update,
		options,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetGameBlock(rowIdx int, colIdx int) ([][]*GameOfLifeUnit, error) {
	gameBlock := GameBlockDoc{}

	location := getBlockLocation(rowIdx, colIdx)
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"location", getBlockLocation(rowIdx, colIdx)}},
	).Decode(&gameBlock)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("Game block was not found, location: %s\n", location)
		return nil, err
	}

	return gameBlock.Units, nil
}

func IniitalizeGameBlocks(coll *mongo.Collection) {
	collection = coll
}
