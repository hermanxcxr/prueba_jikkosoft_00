package models

import "time"

type Prestamo struct {
	ID                      int
	CopiaID                 int
	MiembroID               int
	SedePrestamoID          int
	FechaPrestamo           time.Time
	FechaDevolucionEsperada time.Time
	FechaDevolucionReal     *time.Time
	SedeDevolucionID        *int
	Estado                  string
	CreatedAt               time.Time
}
