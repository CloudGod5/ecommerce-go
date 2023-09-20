package controllers

import ()

func AddAddress() gin.Handlerfunc {

}

func EditHomeAddress() gin.Handlerfunc {

}

func EditWorkAddress() gin.Handlerfunc {

}

func DeleteAddress() gin.Handlerfunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error":"Invalid search param"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D(primitive.E{Key: "_id", Value: user_id})
		update := bson.D{{Key:"$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(400, "Bad Request")
			return
		}
		defer cancel()
		
		ctx.Done()
		c.IndentedJSON(200, "Address deleted successfully")

	
	}
}
