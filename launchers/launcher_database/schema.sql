DROP TABLE IF EXISTS prestamos CASCADE;
DROP TABLE IF EXISTS copias CASCADE;
DROP TABLE IF EXISTS miembros CASCADE;
DROP TABLE IF EXISTS inventario CASCADE;
DROP TABLE IF EXISTS sedes CASCADE;

CREATE TABLE sedes (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE,
    direccion VARCHAR(255),
    telefono VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sedes_nombre ON sedes(nombre);

CREATE TABLE inventario (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    autor VARCHAR(255) NOT NULL,
    isbn VARCHAR(13) UNIQUE,
    categoria VARCHAR(100),
    descripcion TEXT,
    hash_busqueda VARCHAR(32),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inventario_hash ON inventario(hash_busqueda);
CREATE INDEX idx_inventario_titulo ON inventario(titulo);
CREATE INDEX idx_inventario_autor ON inventario(autor);
CREATE INDEX idx_inventario_isbn ON inventario(isbn);

CREATE TABLE copias (
    id SERIAL PRIMARY KEY,
    inventario_id INTEGER NOT NULL REFERENCES inventario(id) ON DELETE CASCADE,
    codigo_unico VARCHAR(50) NOT NULL UNIQUE,
    sede_id INTEGER NOT NULL REFERENCES sedes(id) ON DELETE RESTRICT,
    estado VARCHAR(20) NOT NULL DEFAULT 'disponible',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_estado CHECK (estado IN ('disponible', 'prestado', 'mantenimiento'))
);

CREATE INDEX idx_copias_inventario ON copias(inventario_id);
CREATE INDEX idx_copias_sede ON copias(sede_id);
CREATE INDEX idx_copias_estado ON copias(estado);
CREATE INDEX idx_copias_codigo ON copias(codigo_unico);
CREATE INDEX idx_copias_disponibilidad ON copias(inventario_id, estado, sede_id);

CREATE TABLE miembros (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE,
    telefono VARCHAR(20),
    direccion VARCHAR(255),
    fecha_registro DATE DEFAULT CURRENT_DATE,
    activo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_miembros_email ON miembros(email);
CREATE INDEX idx_miembros_activo ON miembros(activo);
CREATE INDEX idx_miembros_nombre ON miembros(nombre, apellido);

CREATE TABLE prestamos (
    id SERIAL PRIMARY KEY,
    copia_id INTEGER NOT NULL REFERENCES copias(id) ON DELETE RESTRICT,
    miembro_id INTEGER NOT NULL REFERENCES miembros(id) ON DELETE RESTRICT,
    sede_prestamo_id INTEGER NOT NULL REFERENCES sedes(id) ON DELETE RESTRICT,
    fecha_prestamo TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_devolucion_esperada DATE NOT NULL,
    fecha_devolucion_real TIMESTAMP,
    sede_devolucion_id INTEGER REFERENCES sedes(id) ON DELETE RESTRICT,
    estado VARCHAR(20) NOT NULL DEFAULT 'activo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_prestamo_estado CHECK (estado IN ('activo', 'devuelto', 'vencido'))
);

CREATE INDEX idx_prestamos_copia ON prestamos(copia_id);
CREATE INDEX idx_prestamos_miembro ON prestamos(miembro_id);
CREATE INDEX idx_prestamos_estado ON prestamos(estado);
CREATE INDEX idx_prestamos_fecha ON prestamos(fecha_prestamo);
CREATE INDEX idx_prestamos_activos ON prestamos(copia_id, estado);
CREATE INDEX idx_prestamos_vencidos ON prestamos(estado, fecha_devolucion_esperada);

CREATE OR REPLACE FUNCTION calcular_hash_busqueda()
RETURNS TRIGGER AS $$
BEGIN
    NEW.hash_busqueda := MD5(LOWER(
        COALESCE(NEW.titulo, '') || 
        COALESCE(NEW.autor, '') || 
        COALESCE(NEW.isbn, '')
    ));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_hash_busqueda
BEFORE INSERT OR UPDATE ON inventario
FOR EACH ROW
EXECUTE FUNCTION calcular_hash_busqueda();