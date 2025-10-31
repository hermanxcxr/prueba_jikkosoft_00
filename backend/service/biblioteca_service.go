package service

import (
	"database/sql"
	"errors"
	"redbibliotecas/backend/models"
	"redbibliotecas/backend/repository"
	"time"
)

func PrestarLibro(db *sql.DB, codigoCopia string, miembroID int, sedeID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	copia, err := repository.GetCopiaByCodigoUnico(db, codigoCopia)
	if err != nil {
		return errors.New("copia no encontrada")
	}

	if copia.Estado != "disponible" {
		return errors.New("la copia no esta disponible")
	}

	miembro, err := repository.GetMiembroByID(db, miembroID)
	if err != nil {
		return errors.New("miembro no encontrado")
	}

	if !miembro.Activo {
		return errors.New("el miembro no esta activo")
	}

	count, err := repository.CountPrestamosActivosByMiembro(db, miembroID)
	if err != nil {
		return err
	}

	if count >= 3 {
		return errors.New("el miembro ya tiene 3 prestamos activos")
	}

	prestamo := &models.Prestamo{
		CopiaID:                 copia.ID,
		MiembroID:               miembroID,
		SedePrestamoID:          sedeID,
		FechaDevolucionEsperada: time.Now().AddDate(0, 0, 15),
		Estado:                  "activo",
	}

	if err := repository.CreatePrestamo(db, prestamo); err != nil {
		return err
	}

	if err := repository.UpdateCopiaEstado(db, copia.ID, "prestado"); err != nil {
		return err
	}

	return tx.Commit()
}

func DevolverLibro(db *sql.DB, codigoCopia string, sedeDevolucionID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	copia, err := repository.GetCopiaByCodigoUnico(db, codigoCopia)
	if err != nil {
		return errors.New("copia no encontrada")
	}

	if copia.Estado != "prestado" {
		return errors.New("la copia no esta prestada")
	}

	prestamo, err := repository.GetPrestamoActivoByCopia(db, copia.ID)
	if err != nil {
		return errors.New("no se encontro prestamo activo para esta copia")
	}

	if err := repository.RegistrarDevolucion(db, prestamo.ID, sedeDevolucionID); err != nil {
		return err
	}

	if err := repository.UpdateCopiaSede(db, copia.ID, sedeDevolucionID); err != nil {
		return err
	}

	if err := repository.UpdateCopiaEstado(db, copia.ID, "disponible"); err != nil {
		return err
	}

	return tx.Commit()
}

func GetPrestamosActivos(db *sql.DB) ([]models.Prestamo, error) {
	return repository.GetPrestamosActivos(db)
}
