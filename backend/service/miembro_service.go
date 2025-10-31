package service

import (
	"database/sql"
	"errors"
	"redbibliotecas/backend/models"
	"redbibliotecas/backend/repository"
)

func GetAllMiembros(db *sql.DB) ([]models.Miembro, error) {
	return repository.GetAllMiembros(db)
}

func GetMiembroByID(db *sql.DB, id int) (*models.Miembro, error) {
	return repository.GetMiembroByID(db, id)
}

func CreateMiembro(db *sql.DB, m *models.Miembro) error {
	if m.Nombre == "" || m.Apellido == "" {
		return errors.New("nombre y apellido son requeridos")
	}

	return repository.CreateMiembro(db, m)
}

func UpdateMiembro(db *sql.DB, m *models.Miembro) error {
	if m.Nombre == "" || m.Apellido == "" {
		return errors.New("nombre y apellido son requeridos")
	}

	existing, err := repository.GetMiembroByID(db, m.ID)
	if err != nil {
		return errors.New("miembro no encontrado")
	}

	if existing == nil {
		return errors.New("miembro no existe")
	}

	return repository.UpdateMiembro(db, m)
}

func DeleteMiembro(db *sql.DB, id int) error {
	tienePrestamos, err := repository.MiembroTienePrestamosActivos(db, id)
	if err != nil {
		return err
	}

	if tienePrestamos {
		return errors.New("no se puede eliminar un miembro con prestamos activos")
	}

	return repository.DeleteMiembro(db, id)
}
