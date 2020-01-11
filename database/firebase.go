package database

import (
	"context"

	"github.com/zerofelx/Behemoth/key"
	"google.golang.org/api/option"
)

type Data struct {
	Nombre     string
	Correo     string
	Contraseña []byte
	Activo     bool
}

var (
	FKey = key.ReturnKey()
	ctx  = context.Background()
	sa   = option.WithCredentialsJSON(FKey)
)
