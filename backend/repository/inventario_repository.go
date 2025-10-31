package repository

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"redbibliotecas/backend/models"
	"strings"
)

func GetAllInventario(db *sql.DB) ([]models.Inventario, error) {
	query := "SELECT id, titulo, autor, isbn, categoria, descripcion, hash_busqueda, created_at FROM inventario ORDER BY titulo"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventarios []models.Inventario
	for rows.Next() {
		var inv models.Inventario
		err := rows.Scan(&inv.ID, &inv.Titulo, &inv.Autor, &inv.ISBN, &inv.Categoria, &inv.Descripcion, &inv.HashBusqueda, &inv.CreatedAt)
		if err != nil {
			return nil, err
		}
		inventarios = append(inventarios, inv)
	}

	return inventarios, nil
}

func BuscarLibros(db *sql.DB, termino string) ([]models.Inventario, error) {
	hash := fmt.Sprintf("%x", md5Hash(termino))

	query := `
		SELECT id, titulo, autor, isbn, categoria, descripcion, hash_busqueda, created_at 
		FROM inventario 
		WHERE hash_busqueda = $1
		   OR LOWER(titulo) LIKE LOWER($2)
		   OR LOWER(autor) LIKE LOWER($2)
		   OR isbn LIKE $2
		ORDER BY titulo
	`

	searchTerm := "%" + termino + "%"
	rows, err := db.Query(query, hash, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventarios []models.Inventario
	for rows.Next() {
		var inv models.Inventario
		err := rows.Scan(&inv.ID, &inv.Titulo, &inv.Autor, &inv.ISBN, &inv.Categoria, &inv.Descripcion, &inv.HashBusqueda, &inv.CreatedAt)
		if err != nil {
			return nil, err
		}
		inventarios = append(inventarios, inv)
	}

	return inventarios, nil
}

func GetInventarioByID(db *sql.DB, id int) (*models.Inventario, error) {
	query := "SELECT id, titulo, autor, isbn, categoria, descripcion, hash_busqueda, created_at FROM inventario WHERE id = $1"

	var inv models.Inventario
	err := db.QueryRow(query, id).Scan(&inv.ID, &inv.Titulo, &inv.Autor, &inv.ISBN, &inv.Categoria, &inv.Descripcion, &inv.HashBusqueda, &inv.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(strings.ToLower(text)))
	return fmt.Sprintf("%x", hash)
}
