# Sistema de Gestion de Bibliotecas

Sistema de gestion de bibliotecas multi-sede.

## Estado del Proyecto

Fase 1: Base de datos - Implementada

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

Esto ejecutara el launcher de base de datos que creara:
- Base de datos biblioteca
- Tablas necesarias (sedes, inventario, copias, miembros, prestamos)
- Datos iniciales (2 sedes, 10 libros, 30 copias, 5 miembros)

## Estructura del Proyecto

```
redbibliotecas/
├── launcher.go                          # Launcher principal
├── launchers/
│   └── launcher_database/
│       ├── launcher_db.go               # Inicializador de BD
│       ├── schema.sql                   # Esquema de tablas
│       └── seed.sql                     # Datos iniciales
├── .env.example                         # Plantilla de configuracion
└── README.md
```

## Base de Datos

### Tablas

- **sedes**: Sedes de la biblioteca (2 registros)
- **inventario**: Catalogo de libros (10 registros)
- **copias**: Copias fisicas de libros (30 registros)
- **miembros**: Miembros registrados (5 registros)
- **prestamos**: Registro de prestamos (vacia inicialmente)

### Verificacion

Conectarse a la base de datos:

```bash
psql -U postgres -d biblioteca
```

Verificar datos:

```sql
SELECT COUNT(*) FROM sedes;      -- 2
SELECT COUNT(*) FROM inventario; -- 10
SELECT COUNT(*) FROM copias;     -- 30
SELECT COUNT(*) FROM miembros;   -- 5
```

## Licencia

Prototipo educativo
