package controllers

func AddAddress() gin.Handlerfunc {

	return func(c *gin.Context) {

		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search param"})
			c.Abort()
			return
		}

		address, err := ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var addresses models.Address

		address.Address = primitive.NewObjectID()

		if err := c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var addressinfo []bson.M
		if err = pointcursor.All(ctx, &addressinfo); err != nil {
			panic(err)
		}

		var size int32
		for _, address_no := range addressinfo {
		count := address_no["count"]
		size = count.(int32)
		}
		if size < 2 {
			fileter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(err)
			}
		} 
		else {
			c.IndentedJSON(400, "Not Allowed")
		}
		defer cancel()
		ctx.Done()

	}
}

func EditHomeAddress() gin.Handlerfunc {
	return func(c *gin.Context) {

		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search param"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		 var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		 defer cancel()

		 filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		 update := bson.D{{key:"$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: editaddress.House}, {key:"address.0.street_name", Value: editaddress.Street}, {key:"address.0.city_name", Value: editaddress.City}, {key:"address.0.state", Value: editaddress.State}, {key:"address.0.pin_code", Value: editaddress.Pincode}}}}
		 _, err = UserCollection.UpdateOne(ctx, filter, update)
		 if err != nil {
			 c.IndentedJSON(500, "Something went wrong")
			 return
		 }
		 defer cancel()

		 ctx.Done()
		 c.IndentedJSON(200, "Address updated successfully")

	}
}

func EditWorkAddress() gin.Handlerfunc {
		
	return func(c *gin.Context) {
			user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search param"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		 var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		 defer cancel()

		 filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		 update := bson.D{{key:"$set", Value: bson.D{primitive.E{Key: "address.1.house_name", Value: editaddress.House}, {key:"address.1.street_name", Value: editaddress.Street}, {key:"address.1.city_name", Value: editaddress.City}, {key:"address.1.state", Value: editaddress.State}, {key:"address.1.pin_code", Value: editaddress.Pincode}}}}
		 _, err = UserCollection.UpdateOne(ctx, filter, update)
		 if err != nil {
			 c.IndentedJSON(500, "Something went wrong")
			 return
		 }
		 defer cancel()

		 ctx.Done()
		 c.IndentedJSON(200, "Address updated the work successfully")

	}
}

func DeleteAddress() gin.Handlerfunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search param"})
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
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}

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
