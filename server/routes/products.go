package routes

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var validate = validator.New()
var ProductCollection *mongo.Collection = OpenCollention(Client,os.Getenv("PRODUCTS_COLLECTION_NAME"))


func GetProducts(c *gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var products []bson.M

	cursor, err := ProductCollection.Find(ctx,bson.M{})
	if err != nil{
		defer cancel()
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	if err := cursor.All(ctx, &products); err != nil{
		defer cancel()
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK,products)
}

func GetProductById(c *gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	productID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(productID)
	var product bson.M

	result := ProductCollection.FindOne(ctx, bson.M{"_id":docID})

	if result.Err() != nil{
		defer cancel()
		c.JSON(http.StatusInternalServerError, gin.H{"error":result.Err().Error()})
		return
	}

	if err := result.Decode(&product); err != nil{
		
		if err == mongo.ErrNoDocuments{
			defer cancel()
			c.JSON(http.StatusNotFound, gin.H{"error":"Product Not Found!"})
			return
		}

		defer cancel()
    	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	defer cancel()
	c.JSON(http.StatusOK, product)

}