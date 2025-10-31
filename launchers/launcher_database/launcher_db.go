package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func main() {
	fmt.Println(" INICIALIZACION DE BASE DE DATOS ")
	fmt.Println()

	// Cargar configuración
	config := loadConfig()

	// Conectar a PostgreSQL (sin especificar base de datos)
	fmt.Println("Conectando a PostgreSQL...")
	db := connectPostgres(config)
	defer db.Close()

	// Crear base de datos si no existe
	fmt.Println("Verificando base de datos...")
	createDatabase(db, config.DBName)

	// Cerrar conexión inicial
	db.Close()

	// Conectar a la base de datos específica
	fmt.Println("Conectando a base de datos biblioteca...")
	db = connectDatabase(config)
	defer db.Close()

	// Ejecutar schema.sql
	fmt.Println("Ejecutando schema.sql...")
	executeScript(db, "launchers/launcher_database/schema.sql")

	// Ejecutar seed.sql
	fmt.Println("Ejecutando seed.sql...")
	executeScript(db, "launchers/launcher_database/seed.sql")

	// Validar instalación
	fmt.Println()
	fmt.Println("Validando instalacion...")
	validateInstallation(db)

	fmt.Println()
	fmt.Println("✓ Base de datos inicializada correctamente")
	fmt.Println()
}

func loadConfig() Config {
	envPath, _ := filepath.Abs("../../.env")
	if err := godotenv.Load(envPath); err != nil {
		godotenv.Load(".env")
	}

	return Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "biblioteca"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func connectPostgres(config Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error abriendo conexion a PostgreSQL:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}

	return db
}

func connectDatabase(config Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error abriendo conexion a base de datos:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error conectando a base de datos:", err)
	}

	return db
}

func createDatabase(db *sql.DB, dbName string) {
	// Verificar si la base de datos existe
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname='%s')", dbName)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatal("Error verificando existencia de base de datos:", err)
	}

	if exists {
		fmt.Printf("Base de datos '%s' ya existe\n", dbName)

		// Preguntar si desea recrearla
		fmt.Print("¿Desea recrear la base de datos? (s/n): ")
		var respuesta string
		fmt.Scanln(&respuesta)

		if respuesta == "s" || respuesta == "S" {
			fmt.Println("Eliminando base de datos existente...")
			// Terminar conexiones activas
			terminateQuery := fmt.Sprintf(`
				SELECT pg_terminate_backend(pg_stat_activity.pid)
				FROM pg_stat_activity
				WHERE pg_stat_activity.datname = '%s'
				AND pid <> pg_backend_pid()
			`, dbName)
			_, _ = db.Exec(terminateQuery)

			// Eliminar base de datos
			dropQuery := fmt.Sprintf("DROP DATABASE %s", dbName)
			if _, err := db.Exec(dropQuery); err != nil {
				log.Fatal("Error eliminando base de datos:", err)
			}

			// Crear nueva base de datos
			createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
			if _, err := db.Exec(createQuery); err != nil {
				log.Fatal("Error creando base de datos:", err)
			}
			fmt.Printf("Base de datos '%s' recreada exitosamente\n", dbName)
		} else {
			fmt.Println("Continuando con base de datos existente...")
		}
	} else {
		// Crear base de datos
		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if _, err := db.Exec(createQuery); err != nil {
			log.Fatal("Error creando base de datos:", err)
		}
		fmt.Printf("Base de datos '%s' creada exitosamente\n", dbName)
	}
}

func executeScript(db *sql.DB, scriptPath string) {
	// Leer archivo SQL
	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		log.Fatalf("Error obteniendo ruta absoluta de %s: %v", scriptPath, err)
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatalf("Error leyendo %s: %v", scriptPath, err)
	}

	// Ejecutar script
	if _, err := db.Exec(string(content)); err != nil {
		log.Fatalf("Error ejecutando %s: %v", scriptPath, err)
	}

	fmt.Printf("✓ %s ejecutado exitosamente\n", filepath.Base(scriptPath))
}

func validateInstallation(db *sql.DB) {
	tables := []string{"sedes", "inventario", "copias", "miembros", "prestamos"}

	for _, table := range tables {
		var count int
		query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
		err := db.QueryRow(query).Scan(&count)
		if err != nil {
			log.Fatalf("Error validando tabla %s: %v", table, err)
		}
		fmt.Printf("  ✓ Tabla %-12s: %d registros\n", table, count)
	}

	// Validaciones específicas
	var sedeCount, libroCount, copiaCount, miembroCount int
	db.QueryRow("SELECT COUNT(*) FROM sedes").Scan(&sedeCount)
	db.QueryRow("SELECT COUNT(*) FROM inventario").Scan(&libroCount)
	db.QueryRow("SELECT COUNT(*) FROM copias").Scan(&copiaCount)
	db.QueryRow("SELECT COUNT(*) FROM miembros").Scan(&miembroCount)

	fmt.Println()
	if sedeCount == 2 && libroCount == 10 && copiaCount == 30 && miembroCount == 5 {
		fmt.Println("✓ Todos los datos iniciales se insertaron correctamente")
	} else {
		fmt.Println("⚠ Advertencia: Los conteos no coinciden con los esperados")
		fmt.Printf("  Esperado: 2 sedes, 10 libros, 30 copias, 5 miembros\n")
		fmt.Printf("  Obtenido: %d sedes, %d libros, %d copias, %d miembros\n",
			sedeCount, libroCount, copiaCount, miembroCount)
	}
}
