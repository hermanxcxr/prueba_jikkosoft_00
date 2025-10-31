package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("=== LAUNCHER - SISTEMA DE GESTION DE BIBLIOTECAS ===")
	fmt.Println()

	// Lista de launchers a ejecutar
	// El desarrollador puede comentar/descomentar segun necesite
	launchers := []string{
		"launchers/launcher_database/launcher_db.go", // Fase 1: Base de datos
		// "launchers/launcher_api/launcher_api.go",     // Fase 3: API (pendiente)
		// "launchers/launcher_frontend/launcher_frontend.sh", // Fase 4: Frontend (pendiente)
	}

	ejecutarLaunchers(launchers)
}

func ejecutarLaunchers(launchers []string) {
	for i, launcher := range launchers {
		fmt.Printf("[%d/%d] Ejecutando: %s\n", i+1, len(launchers), launcher)
		fmt.Println()

		if err := ejecutarLauncher(launcher); err != nil {
			log.Printf("Error ejecutando %s: %v\n", launcher, err)
			fmt.Println()
			fmt.Print("¿Desea continuar con los siguientes launchers? (s/n): ")
			var respuesta string
			fmt.Scanln(&respuesta)
			if respuesta != "s" && respuesta != "S" {
				fmt.Println("Proceso cancelado por el usuario")
				os.Exit(1)
			}
		}

		fmt.Println()
		fmt.Println("---")
		fmt.Println()
	}

	fmt.Println("✓ Todos los launchers se ejecutaron")
}

func ejecutarLauncher(launcher string) error {
	cmd := exec.Command("go", "run", launcher)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
