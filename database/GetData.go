package database

import (
	"log"

	firebase "firebase.google.com/go"
)

func GetData(document string) *Data {
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Panic(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Panic(err)
	}

	defer client.Close()

	get, err := client.Collection("users").Doc(document).Get(ctx)
	if err != nil {
		panic(err)
	}
	estructDatos := new(Data)
	get.DataTo(estructDatos)

	return estructDatos
}
