package repository

import (
	"database/sql"
	"redbibliotecas/backend/models"
)

func GetCopiaByCodigoUnico(db *sql.DB, codigoUnico string) (*models.Copia, error) {
	query := "SELECT id, inventario_id, codigo_unico, sede_id, estado, created_at FROM copias WHERE codigo_unico = $1"

	var c models.Copia
	err := db.QueryRow(query, codigoUnico).Scan(&c.ID, &c.InventarioID, &c.CodigoUnico, &c.SedeID, &c.Estado, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetCopiasDisponiblesPorSede(db *sql.DB, inventarioID int, sedeID int) ([]models.Copia, error) {
	query := `
		SELECT id, inventario_id, codigo_unico, sede_id, estado, created_at 
		FROM copias 
		WHERE inventario_id = $1 AND sede_id = $2 AND estado = 'disponible'
		ORDER BY codigo_unico
	`

	rows, err := db.Query(query, inventarioID, sedeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var copias []models.Copia
	for rows.Next() {
		var c models.Copia
		err := rows.Scan(&c.ID, &c.InventarioID, &c.CodigoUnico, &c.SedeID, &c.Estado, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		copias = append(copias, c)
	}

	return copias, nil
}

func UpdateCopiaEstado(db *sql.DB, copiaID int, estado string) error {
	query := "UPDATE copias SET estado = $1 WHERE id = $2"
	_, err := db.Exec(query, estado, copiaID)
	return err
}

func UpdateCopiaSede(db *sql.DB, copiaID int, sedeID int) error {
	query := "UPDATE copias SET sede_id = $1 WHERE id = $2"
	_, err := db.Exec(query, sedeID, copiaID)
	return err
}

func CountCopiasDisponibles(db *sql.DB, inventarioID int, sedeID int) (int, error) {
	query := "SELECT COUNT(*) FROM copias WHERE inventario_id = $1 AND sede_id = $2 AND estado = 'disponible'"

	var count int
	err := db.QueryRow(query, inventarioID, sedeID).Scan(&count)
	return count, err
}
