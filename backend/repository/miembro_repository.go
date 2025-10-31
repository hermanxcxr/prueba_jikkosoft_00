package repository

import (
	"database/sql"
	"redbibliotecas/backend/models"
)

func GetAllMiembros(db *sql.DB) ([]models.Miembro, error) {
	query := "SELECT id, nombre, apellido, email, telefono, direccion, fecha_registro, activo, created_at FROM miembros ORDER BY apellido, nombre"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miembros []models.Miembro
	for rows.Next() {
		var m models.Miembro
		err := rows.Scan(&m.ID, &m.Nombre, &m.Apellido, &m.Email, &m.Telefono, &m.Direccion, &m.FechaRegistro, &m.Activo, &m.CreatedAt)
		if err != nil {
			return nil, err
		}
		miembros = append(miembros, m)
	}

	return miembros, nil
}

func GetMiembroByID(db *sql.DB, id int) (*models.Miembro, error) {
	query := "SELECT id, nombre, apellido, email, telefono, direccion, fecha_registro, activo, created_at FROM miembros WHERE id = $1"

	var m models.Miembro
	err := db.QueryRow(query, id).Scan(&m.ID, &m.Nombre, &m.Apellido, &m.Email, &m.Telefono, &m.Direccion, &m.FechaRegistro, &m.Activo, &m.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func CreateMiembro(db *sql.DB, m *models.Miembro) error {
	query := "INSERT INTO miembros (nombre, apellido, email, telefono, direccion) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := db.QueryRow(query, m.Nombre, m.Apellido, m.Email, m.Telefono, m.Direccion).Scan(&m.ID)
	return err
}

func UpdateMiembro(db *sql.DB, m *models.Miembro) error {
	query := "UPDATE miembros SET nombre = $1, apellido = $2, email = $3, telefono = $4, direccion = $5 WHERE id = $6"
	_, err := db.Exec(query, m.Nombre, m.Apellido, m.Email, m.Telefono, m.Direccion, m.ID)
	return err
}

func DeleteMiembro(db *sql.DB, id int) error {
	query := "DELETE FROM miembros WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}

func MiembroTienePrestamosActivos(db *sql.DB, miembroID int) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM prestamos WHERE miembro_id = $1 AND estado = 'activo')"

	var exists bool
	err := db.QueryRow(query, miembroID).Scan(&exists)
	return exists, err
}
