package repository

import (
	"database/sql"
	"redbibliotecas/backend/models"
)

func GetAllSedes(db *sql.DB) ([]models.Sede, error) {
	query := "SELECT id, nombre, direccion, telefono, created_at FROM sedes ORDER BY id"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sedes []models.Sede
	for rows.Next() {
		var s models.Sede
		err := rows.Scan(&s.ID, &s.Nombre, &s.Direccion, &s.Telefono, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		sedes = append(sedes, s)
	}

	return sedes, nil
}

func GetSedeByID(db *sql.DB, id int) (*models.Sede, error) {
	query := "SELECT id, nombre, direccion, telefono, created_at FROM sedes WHERE id = $1"

	var s models.Sede
	err := db.QueryRow(query, id).Scan(&s.ID, &s.Nombre, &s.Direccion, &s.Telefono, &s.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
