package service

import (
	"database/sql"
	"redbibliotecas/backend/models"
	"redbibliotecas/backend/repository"
)

type ResultadoBusqueda struct {
	Inventario            models.Inventario
	CopiasDisponibles     int
	CopiasDisponiblesSede int
}

func BuscarLibros(db *sql.DB, termino string, sedeID int) ([]ResultadoBusqueda, error) {
	inventarios, err := repository.BuscarLibros(db, termino)
	if err != nil {
		return nil, err
	}

	var resultados []ResultadoBusqueda
	for _, inv := range inventarios {
		disponiblesSede, err := repository.CountCopiasDisponibles(db, inv.ID, sedeID)
		if err != nil {
			continue
		}

		resultado := ResultadoBusqueda{
			Inventario:            inv,
			CopiasDisponiblesSede: disponiblesSede,
		}
		resultados = append(resultados, resultado)
	}

	return resultados, nil
}

func GetCopiasDisponibles(db *sql.DB, inventarioID int, sedeID int) ([]models.Copia, error) {
	return repository.GetCopiasDisponiblesPorSede(db, inventarioID, sedeID)
}
