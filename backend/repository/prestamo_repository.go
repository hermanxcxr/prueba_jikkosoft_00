package repository

import (
	"database/sql"
	"redbibliotecas/backend/models"
	"time"
)

func CreatePrestamo(db *sql.DB, p *models.Prestamo) error {
	query := `
		INSERT INTO prestamos (copia_id, miembro_id, sede_prestamo_id, fecha_devolucion_esperada) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, fecha_prestamo, created_at
	`
	err := db.QueryRow(query, p.CopiaID, p.MiembroID, p.SedePrestamoID, p.FechaDevolucionEsperada).
		Scan(&p.ID, &p.FechaPrestamo, &p.CreatedAt)
	return err
}

func GetPrestamoActivoByCopia(db *sql.DB, copiaID int) (*models.Prestamo, error) {
	query := `
		SELECT id, copia_id, miembro_id, sede_prestamo_id, fecha_prestamo, 
		       fecha_devolucion_esperada, fecha_devolucion_real, sede_devolucion_id, estado, created_at 
		FROM prestamos 
		WHERE copia_id = $1 AND estado = 'activo'
	`

	var p models.Prestamo
	err := db.QueryRow(query, copiaID).Scan(
		&p.ID, &p.CopiaID, &p.MiembroID, &p.SedePrestamoID, &p.FechaPrestamo,
		&p.FechaDevolucionEsperada, &p.FechaDevolucionReal, &p.SedeDevolucionID, &p.Estado, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func RegistrarDevolucion(db *sql.DB, prestamoID int, sedeDevolucionID int) error {
	query := `
		UPDATE prestamos 
		SET fecha_devolucion_real = $1, sede_devolucion_id = $2, estado = 'devuelto' 
		WHERE id = $3
	`
	_, err := db.Exec(query, time.Now(), sedeDevolucionID, prestamoID)
	return err
}

func CountPrestamosActivosByMiembro(db *sql.DB, miembroID int) (int, error) {
	query := "SELECT COUNT(*) FROM prestamos WHERE miembro_id = $1 AND estado = 'activo'"

	var count int
	err := db.QueryRow(query, miembroID).Scan(&count)
	return count, err
}

func GetPrestamosActivos(db *sql.DB) ([]models.Prestamo, error) {
	query := `
		SELECT id, copia_id, miembro_id, sede_prestamo_id, fecha_prestamo, 
		       fecha_devolucion_esperada, fecha_devolucion_real, sede_devolucion_id, estado, created_at 
		FROM prestamos 
		WHERE estado = 'activo'
		ORDER BY fecha_prestamo DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prestamos []models.Prestamo
	for rows.Next() {
		var p models.Prestamo
		err := rows.Scan(&p.ID, &p.CopiaID, &p.MiembroID, &p.SedePrestamoID, &p.FechaPrestamo,
			&p.FechaDevolucionEsperada, &p.FechaDevolucionReal, &p.SedeDevolucionID, &p.Estado, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, p)
	}

	return prestamos, nil
}
