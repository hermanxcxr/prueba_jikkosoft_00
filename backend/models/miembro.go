package models

import "time"

type Miembro struct {
	ID            int
	Nombre        string
	Apellido      string
	Email         string
	Telefono      string
	Direccion     string
	FechaRegistro time.Time
	Activo        bool
	CreatedAt     time.Time
}
