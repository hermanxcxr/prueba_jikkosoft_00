# Sistema de Gestion de Bibliotecas

Sistema de gestion de bibliotecas multi-sede.

## Estado del Proyecto

Fase 1: Base de datos - Completada
Fase 2: Backend + Menu Terminal - Completada

## Requisitos

- PostgreSQL 13
- Go 1.19 o superior

## Instalacion

### 1. Clonar el repositorio

```bash
git clone https://github.com/usuario/redbibliotecas.git
cd redbibliotecas
```

### 2. Configurar variables de entorno

Copiar el archivo de ejemplo y ajustar los valores:

```bash
cp .env.example .env
```

Editar .env con los datos de PostgreSQL:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=biblioteca
```

### 3. Instalar dependencias

```bash
go mod download
```

### 4. Inicializar la base de datos

```bash
go run launcher.go
```

Esto creara la base de datos con tablas y datos iniciales.

### 5. Ejecutar el menu terminal

```bash
go run menu.go
```

## Estructura del Proyecto

```
redbibliotecas/
├── launcher.go                          # Inicializador
├── menu.go                              # Menu terminal
├── backend/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── connection.go
│   ├── models/
│   │   ├── sede.go
│   │   ├── inventario.go
│   │   ├── copia.go
│   │   ├── miembro.go
│   │   └── prestamo.go
│   ├── repository/
│   │   ├── sede_repository.go
│   │   ├── inventario_repository.go
│   │   ├── copia_repository.go
│   │   ├── miembro_repository.go
│   │   └── prestamo_repository.go
│   └── service/
│       ├── biblioteca_service.go
│       ├── busqueda_service.go
│       └── miembro_service.go
├── launchers/
│   └── launcher_database/
│       ├── launcher_db.go
│       ├── schema.sql
│       └── seed.sql
└── README.md
```

## Funcionalidades

### Menu Terminal

Al ejecutar menu.go se presenta un menu interactivo con las siguientes opciones:

1. Prestar Libro
2. Devolver Libro
3. Buscar Libro
4. Gestionar Miembros
5. Ver Estadisticas
6. Cambiar Sede
7. Salir

### Prestar Libro

- Solicita codigo de copia
- Solicita ID de miembro
- Valida disponibilidad
- Valida que el miembro no tenga mas de 3 prestamos activos
- Registra el prestamo
- Actualiza estado de la copia a prestado

### Devolver Libro

- Solicita codigo de copia
- Busca el prestamo activo
- Registra la devolucion
- Actualiza la ubicacion de la copia a la sede actual
- Actualiza estado de la copia a disponible

### Buscar Libro

- Solicita termino de busqueda
- Busca por titulo, autor o ISBN
- Muestra resultados con disponibilidad por sede
- Permite ver codigos de copias disponibles

### Gestionar Miembros

- Listar todos los miembros
- Crear nuevo miembro
- Editar miembro existente
- Eliminar miembro (solo si no tiene prestamos activos)

## Base de Datos

### Tablas

- sedes: 2 registros
- inventario: 10 registros
- copias: 30 registros
- miembros: 5 registros
- prestamos: registros de prestamos

### Trigger Automatico

La tabla inventario tiene un trigger que calcula automaticamente el hash_busqueda al insertar o actualizar libros.

## Licencia

Prototipo educativo
