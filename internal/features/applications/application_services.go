package applications

import (
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAppService(appId primitive.ObjectID) (AppModel, error) {

}

func GetAppAllService(batchId primitive.ObjectID) ([]AppModel, error) {

}

func CreateAppService(appDto AppDto) (AppModel, error) {

}

func UpdateAppService(appDto AppDto) (AppModel, error) {

}

func DeleteAppService(appId primitive.ObjectID) {

}