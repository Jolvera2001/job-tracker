package applications

import (
	"context"
	"job-tracker/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const(
	colName = "Applications"
)
func GetAppService(appId primitive.ObjectID) (AppModel, error) {

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

}

func UpdateAppService(appDto AppDto) (AppModel, error) {

}

func DeleteAppService(appId primitive.ObjectID) {

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