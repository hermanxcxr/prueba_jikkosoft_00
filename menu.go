package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"redbibliotecas/backend/config"
	"redbibliotecas/backend/database"
	"redbibliotecas/backend/models"
	"redbibliotecas/backend/repository"
	"redbibliotecas/backend/service"
	"runtime"
)

var (
	db         *sql.DB
	sedeActual int
	nombreSede string
)

func main() {
	fmt.Println("=== SISTEMA DE GESTION DE BIBLIOTECAS ===")
	fmt.Println()

	cfg := config.Load()
	db = database.Connect(cfg)
	defer db.Close()

	seleccionarSede()

	for {
		mostrarMenuPrincipal()
		opcion := leerOpcion()

		switch opcion {
		case 1:
			prestarLibro()
		case 2:
			devolverLibro()
		case 3:
			buscarLibro()
		case 4:
			gestionarMiembros()
		case 5:
			verEstadisticas()
		case 6:
			seleccionarSede()
		case 7:
			fmt.Println("\nSaliendo del sistema...")
			return
		default:
			fmt.Println("\nOpcion invalida")
		}

		esperarEnter()
	}
}

func seleccionarSede() {
	limpiarPantalla()
	fmt.Println("=== SELECCION DE SEDE ===")
	fmt.Println()

	sedes, err := repository.GetAllSedes(db)
	if err != nil {
		fmt.Println("Error obteniendo sedes:", err)
		return
	}

	for _, sede := range sedes {
		fmt.Printf("%d. %s\n", sede.ID, sede.Nombre)
	}

	fmt.Print("\nSeleccione sede: ")
	fmt.Scan(&sedeActual)

	for _, sede := range sedes {
		if sede.ID == sedeActual {
			nombreSede = sede.Nombre
			break
		}
	}

	fmt.Printf("\n✓ Trabajando desde: %s\n", nombreSede)
	esperarEnter()
}

func mostrarMenuPrincipal() {
	limpiarPantalla()
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Printf("║  SISTEMA DE BIBLIOTECAS - %-28s║\n", nombreSede)
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("1. Prestar Libro")
	fmt.Println("2. Devolver Libro")
	fmt.Println("3. Buscar Libro")
	fmt.Println("4. Gestionar Miembros")
	fmt.Println("5. Ver Estadisticas")
	fmt.Println("6. Cambiar Sede")
	fmt.Println("7. Salir")
	fmt.Println()
	fmt.Print("Seleccione opcion: ")
}

func prestarLibro() {
	limpiarPantalla()
	fmt.Println("=== PRESTAR LIBRO ===")
	fmt.Println()

	var codigoCopia string
	var miembroID int

	fmt.Print("Codigo de copia: ")
	fmt.Scan(&codigoCopia)

	fmt.Print("ID de miembro: ")
	fmt.Scan(&miembroID)

	err := service.PrestarLibro(db, codigoCopia, miembroID, sedeActual)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Println("\n✓ Prestamo registrado exitosamente")
}

func devolverLibro() {
	limpiarPantalla()
	fmt.Println("=== DEVOLVER LIBRO ===")
	fmt.Println()

	var codigoCopia string

	fmt.Print("Codigo de copia: ")
	fmt.Scan(&codigoCopia)

	err := service.DevolverLibro(db, codigoCopia, sedeActual)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Println("\n✓ Devolucion registrada exitosamente")
	fmt.Printf("✓ El libro ahora esta en: %s\n", nombreSede)
}

func buscarLibro() {
	limpiarPantalla()
	fmt.Println("=== BUSCAR LIBRO ===")
	fmt.Println()

	var termino string
	fmt.Print("Termino de busqueda: ")
	fmt.Scanln()
	fmt.Scanln(&termino)

	resultados, err := service.BuscarLibros(db, termino, sedeActual)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	if len(resultados) == 0 {
		fmt.Println("\nNo se encontraron resultados")
		return
	}

	fmt.Printf("\n=== RESULTADOS (%d libros) ===\n\n", len(resultados))
	for _, r := range resultados {
		fmt.Printf("ID: %d\n", r.Inventario.ID)
		fmt.Printf("Titulo: %s\n", r.Inventario.Titulo)
		fmt.Printf("Autor: %s\n", r.Inventario.Autor)
		fmt.Printf("ISBN: %s\n", r.Inventario.ISBN)
		fmt.Printf("Categoria: %s\n", r.Inventario.Categoria)
		fmt.Printf("Disponibles en %s: %d\n", nombreSede, r.CopiasDisponiblesSede)
		fmt.Println("---")
	}

	fmt.Print("\n¿Desea ver copias disponibles de algun libro? (s/n): ")
	var respuesta string
	fmt.Scan(&respuesta)

	if respuesta == "s" || respuesta == "S" {
		var inventarioID int
		fmt.Print("ID del libro: ")
		fmt.Scan(&inventarioID)

		copias, err := service.GetCopiasDisponibles(db, inventarioID, sedeActual)
		if err != nil {
			fmt.Printf("\n❌ Error: %v\n", err)
			return
		}

		if len(copias) == 0 {
			fmt.Println("\nNo hay copias disponibles en esta sede")
			return
		}

		fmt.Printf("\n=== COPIAS DISPONIBLES ===\n\n")
		for _, c := range copias {
			fmt.Printf("Codigo: %s\n", c.CodigoUnico)
		}
	}
}

func gestionarMiembros() {
	for {
		limpiarPantalla()
		fmt.Println("=== GESTIONAR MIEMBROS ===")
		fmt.Println()
		fmt.Println("1. Listar Miembros")
		fmt.Println("2. Crear Miembro")
		fmt.Println("3. Editar Miembro")
		fmt.Println("4. Eliminar Miembro")
		fmt.Println("5. Volver")
		fmt.Println()
		fmt.Print("Seleccione opcion: ")

		opcion := leerOpcion()

		switch opcion {
		case 1:
			listarMiembros()
		case 2:
			crearMiembro()
		case 3:
			editarMiembro()
		case 4:
			eliminarMiembro()
		case 5:
			return
		default:
			fmt.Println("\nOpcion invalida")
		}

		esperarEnter()
	}
}

func listarMiembros() {
	limpiarPantalla()
	fmt.Println("=== LISTA DE MIEMBROS ===")
	fmt.Println()

	miembros, err := service.GetAllMiembros(db)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	if len(miembros) == 0 {
		fmt.Println("No hay miembros registrados")
		return
	}

	for _, m := range miembros {
		estado := "Activo"
		if !m.Activo {
			estado = "Inactivo"
		}
		fmt.Printf("ID: %d | %s %s | %s | %s\n", m.ID, m.Nombre, m.Apellido, m.Email, estado)
	}
}

func crearMiembro() {
	limpiarPantalla()
	fmt.Println("=== CREAR MIEMBRO ===")
	fmt.Println()

	var m models.Miembro

	fmt.Print("Nombre: ")
	fmt.Scanln()
	fmt.Scanln(&m.Nombre)

	fmt.Print("Apellido: ")
	fmt.Scanln(&m.Apellido)

	fmt.Print("Email: ")
	fmt.Scanln(&m.Email)

	fmt.Print("Telefono: ")
	fmt.Scanln(&m.Telefono)

	fmt.Print("Direccion: ")
	fmt.Scanln(&m.Direccion)

	err := service.CreateMiembro(db, &m)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Printf("\n✓ Miembro creado exitosamente con ID: %d\n", m.ID)
}

func editarMiembro() {
	limpiarPantalla()
	fmt.Println("=== EDITAR MIEMBRO ===")
	fmt.Println()

	var id int
	fmt.Print("ID del miembro: ")
	fmt.Scan(&id)

	miembro, err := service.GetMiembroByID(db, id)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Printf("\nDatos actuales:\n")
	fmt.Printf("Nombre: %s\n", miembro.Nombre)
	fmt.Printf("Apellido: %s\n", miembro.Apellido)
	fmt.Printf("Email: %s\n", miembro.Email)
	fmt.Printf("Telefono: %s\n", miembro.Telefono)
	fmt.Printf("Direccion: %s\n", miembro.Direccion)
	fmt.Println()

	fmt.Print("Nuevo nombre (Enter para mantener): ")
	fmt.Scanln()
	var nombre string
	fmt.Scanln(&nombre)
	if nombre != "" {
		miembro.Nombre = nombre
	}

	fmt.Print("Nuevo apellido (Enter para mantener): ")
	var apellido string
	fmt.Scanln(&apellido)
	if apellido != "" {
		miembro.Apellido = apellido
	}

	fmt.Print("Nuevo email (Enter para mantener): ")
	var email string
	fmt.Scanln(&email)
	if email != "" {
		miembro.Email = email
	}

	fmt.Print("Nuevo telefono (Enter para mantener): ")
	var telefono string
	fmt.Scanln(&telefono)
	if telefono != "" {
		miembro.Telefono = telefono
	}

	fmt.Print("Nueva direccion (Enter para mantener): ")
	var direccion string
	fmt.Scanln(&direccion)
	if direccion != "" {
		miembro.Direccion = direccion
	}

	err = service.UpdateMiembro(db, miembro)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Println("\n✓ Miembro actualizado exitosamente")
}

func eliminarMiembro() {
	limpiarPantalla()
	fmt.Println("=== ELIMINAR MIEMBRO ===")
	fmt.Println()

	var id int
	fmt.Print("ID del miembro: ")
	fmt.Scan(&id)

	miembro, err := service.GetMiembroByID(db, id)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Printf("\n¿Esta seguro de eliminar a %s %s? (s/n): ", miembro.Nombre, miembro.Apellido)
	var confirmacion string
	fmt.Scan(&confirmacion)

	if confirmacion != "s" && confirmacion != "S" {
		fmt.Println("\nOperacion cancelada")
		return
	}

	err = service.DeleteMiembro(db, id)
	if err != nil {
		fmt.Printf("\n❌ Error: %v\n", err)
		return
	}

	fmt.Println("\n✓ Miembro eliminado exitosamente")
}

func verEstadisticas() {
	limpiarPantalla()
	fmt.Println("=== ESTADISTICAS ===")
	fmt.Println()

	prestamos, err := service.GetPrestamosActivos(db)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	fmt.Printf("Prestamos activos: %d\n", len(prestamos))

	if len(prestamos) > 0 {
		fmt.Println("\n=== PRESTAMOS ACTIVOS ===\n")
		for _, p := range prestamos {
			copia, _ := repository.GetCopiaByCodigoUnico(db, "")
			if copia != nil {
				inv, _ := repository.GetInventarioByID(db, copia.InventarioID)
				miembro, _ := repository.GetMiembroByID(db, p.MiembroID)

				if inv != nil && miembro != nil {
					fmt.Printf("Libro: %s\n", inv.Titulo)
					fmt.Printf("Miembro: %s %s\n", miembro.Nombre, miembro.Apellido)
					fmt.Printf("Fecha prestamo: %s\n", p.FechaPrestamo.Format("2006-01-02"))
					fmt.Printf("Devolucion esperada: %s\n", p.FechaDevolucionEsperada.Format("2006-01-02"))
					fmt.Println("---")
				}
			}
		}
	}
}

func leerOpcion() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

func esperarEnter() {
	fmt.Println()
	fmt.Print("Presione ENTER para continuar...")
	fmt.Scanln()
}

func limpiarPantalla() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
