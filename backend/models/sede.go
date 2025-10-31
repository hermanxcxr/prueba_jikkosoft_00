package models

import "time"

type Sede struct {
	ID        int
	Nombre    string
	Direccion string
	Telefono  string
	CreatedAt time.Time
}
