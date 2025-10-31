package models

import "time"

type Copia struct {
	ID           int
	InventarioID int
	CodigoUnico  string
	SedeID       int
	Estado       string
	CreatedAt    time.Time
}
