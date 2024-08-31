package applications

import (
	"context"
	"job-tracker/internal/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const(
	colName = "Applications"
)
func GetAppService(appId primitive.ObjectID) (AppModel, error) {
	filter := bson.M{"_id": appId}

	var res AppModel
	err := database.GetCollection(colName).FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return AppModel{}, err
	}

	return res, nil
}

func GetAppAllService(batchId primitive.ObjectID) ([]AppModel, error) {
	filter := bson.M{"batch_id": batchId}

	var appList []AppModel
	cur, err := database.GetCollection(colName).Find(context.Background(), filter)
	if err != nil {
		return []AppModel{}, err
	}

	defer func() {
		if closeError := cur.Close(context.Background()); closeError != nil {
			err = closeError
		}
	}()

	err = parseCursor(&appList, cur)
	if err != nil {
		return []AppModel{}, err
	}

	return appList, nil
}

func CreateAppService(appDto AppDto) (AppModel, error) {
	appDoc := AppModel{
		ID: primitive.NewObjectID(),
		BatchId: appDto.BatchId,
		Name: appDto.Name,
		Company: appDto.Company,
		Description: appDto.Description,
		Status: appDto.Status,
		RoundCount: appDto.RoundCount,
		DateCreated: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := database.GetCollection(colName).InsertOne(context.Background(), appDoc)
	if err != nil {
		return AppModel{}, err
	}

	return appDoc, nil
}

func UpdateAppService(update AppModel) (AppModel, error) {
	filter := bson.M{"_id": update.ID}
	updateDoc := bson.M{"$set": update}

	_, err := database.GetCollection(colName).UpdateOne(context.Background(), filter, updateDoc)
	if err != nil {
		return AppModel{}, err
	}

	var res AppModel
	err = database.GetCollection(colName).FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return AppModel{}, err
	}

	return res, nil
}

func DeleteAppService(appId primitive.ObjectID) error {
	filter := bson.M{"_id": appId}

	_, err := database.GetCollection(colName).DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func parseCursor(appList *[]AppModel, cursor *mongo.Cursor) error {
	for cursor.Next(context.Background()) {
		var app AppModel
		if err := cursor.Decode(&app); err != nil {
			return err
		}

		*appList = append(*appList, app)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}