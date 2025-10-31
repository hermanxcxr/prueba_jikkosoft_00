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

	launchers := []string{
		"launchers/launcher_database/launcher_db.go",
	}

	ejecutarLaunchers(launchers)
}

func ejecutarLaunchers(launchers []string) {
	for i, launcher := range launchers {
		fmt.Printf("[%d/%d] Ejecutando: %s\n", i+1, len(launchers), launcher)
		fmt.Println()

		if err := ejecutarLauncher(launcher); err != nil {
			log.Fatalf("Error ejecutando %s: %v\n", launcher, err)
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
