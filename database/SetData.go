package database

import (
	"crypto/sha256"
	"fmt"
	"strings"

	firebase "firebase.google.com/go"
)

func SetData(data *Data) bool {
	usuario := data.Nombre
	usuario = strings.ToLower(usuario)
	usuario = strings.Title(usuario)

	correo := strings.ToLower(data.Correo)

	psw := data.Contraseña

	password := sha256.New()
	password.Write(psw)
	shpsw := password.Sum(nil)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	get, err := client.Collection("users").Doc(usuario).Get(ctx)
	if len(get.Data()) == 0 {
		_, err = client.Collection("users").Doc(usuario).Set(ctx, map[string]interface{}{
			"Nombre":     usuario,
			"Correo":     correo,
			"Contraseña": shpsw,
			"Activo":     false,
		})
		if err != nil {
			panic(err)
		}
		er := false
		return er
	} else {
		a := get.Data()
		fmt.Println(a)
		er := true
		return er
	}
}
