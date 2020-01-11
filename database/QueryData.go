package database

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"

	firebase "firebase.google.com/go"
)

func QueryPassword(username string, psw string) bool {
	usuario := username
	usuario = strings.ToLower(usuario)
	usuario = strings.Title(usuario)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	get, err := client.Collection("users").Doc(usuario).Get(ctx)
	if err != nil {
		fmt.Println(err)
		return false
	}

	datos := new(Data)
	get.DataTo(datos)

	defer client.Close()
	if datos.Activo != false {
		contraseña := datos.Contraseña
		sum := sha256.Sum256([]byte(psw))

		verify := bytes.Equal(sum[:], contraseña[:])
		fmt.Println(verify)
		fmt.Println(usuario)
		return verify
	} else {
		fmt.Println("falso")
		return false
	}
}
