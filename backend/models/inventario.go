package models

import "time"

type Inventario struct {
	ID           int
	Titulo       string
	Autor        string
	ISBN         string
	Categoria    string
	Descripcion  string
	HashBusqueda string
	CreatedAt    time.Time
}
