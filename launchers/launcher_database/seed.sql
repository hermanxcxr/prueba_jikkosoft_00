INSERT INTO sedes (nombre, direccion, telefono) VALUES
('Sede Central', 'Calle 10 #20-30, Centro', '555-0100'),
('Sede Norte', 'Avenida 45 #80-15, Norte', '555-0200');

INSERT INTO inventario (titulo, autor, isbn, categoria, descripcion) VALUES
('Cien Años de Soledad', 'Gabriel García Márquez', '9780307474728', 'Ficción', 'Obra maestra del realismo mágico que narra la historia de la familia Buendía'),
('Don Quijote de la Mancha', 'Miguel de Cervantes', '9788424936464', 'Clásico', 'La obra cumbre de la literatura española'),
('El Principito', 'Antoine de Saint-Exupéry', '9780156012195', 'Infantil', 'Fábula filosófica sobre la naturaleza humana'),
('1984', 'George Orwell', '9780451524935', 'Distopía', 'Novela distópica sobre un futuro totalitario'),
('Crónica de una Muerte Anunciada', 'Gabriel García Márquez', '9780307387387', 'Ficción', 'Relato de un asesinato anunciado en un pueblo colombiano'),
('El Amor en los Tiempos del Cólera', 'Gabriel García Márquez', '9780307387738', 'Romance', 'Historia de amor que trasciende el tiempo'),
('La Sombra del Viento', 'Carlos Ruiz Zafón', '9780143034902', 'Misterio', 'Misterio ambientado en la Barcelona de posguerra'),
('Rayuela', 'Julio Cortázar', '9788420471891', 'Experimental', 'Novela experimental que puede leerse en múltiples órdenes'),
('Pedro Páramo', 'Juan Rulfo', '9780802133908', 'Realismo Mágico', 'Novela corta sobre un pueblo fantasma en México'),
('El Túnel', 'Ernesto Sabato', '9788432217326', 'Psicológico', 'Novela psicológica sobre obsesión y locura');

INSERT INTO copias (inventario_id, codigo_unico, sede_id, estado) VALUES
(1, 'CIEN-CENT-001', 1, 'disponible'),
(1, 'CIEN-CENT-002', 1, 'disponible'),
(1, 'CIEN-CENT-003', 1, 'disponible'),
(1, 'CIEN-NORT-001', 2, 'disponible'),
(1, 'CIEN-NORT-002', 2, 'disponible'),
(2, 'QUIJ-CENT-001', 1, 'disponible'),
(2, 'QUIJ-CENT-002', 1, 'disponible'),
(2, 'QUIJ-NORT-001', 2, 'disponible'),
(2, 'QUIJ-NORT-002', 2, 'disponible'),
(3, 'PRIN-CENT-001', 1, 'disponible'),
(3, 'PRIN-CENT-002', 1, 'disponible'),
(3, 'PRIN-NORT-001', 2, 'disponible'),
(4, '1984-CENT-001', 1, 'disponible'),
(4, '1984-CENT-002', 1, 'disponible'),
(4, '1984-NORT-001', 2, 'disponible'),
(4, '1984-NORT-002', 2, 'disponible'),
(5, 'CRON-CENT-001', 1, 'disponible'),
(5, 'CRON-NORT-001', 2, 'disponible'),
(6, 'AMOR-CENT-001', 1, 'disponible'),
(6, 'AMOR-CENT-002', 1, 'disponible'),
(6, 'AMOR-NORT-001', 2, 'disponible'),
(7, 'SOMB-CENT-001', 1, 'disponible'),
(7, 'SOMB-NORT-001', 2, 'disponible'),
(7, 'SOMB-NORT-002', 2, 'disponible'),
(8, 'RAYU-CENT-001', 1, 'disponible'),
(8, 'RAYU-NORT-001', 2, 'disponible'),
(9, 'PEDR-CENT-001', 1, 'disponible'),
(9, 'PEDR-NORT-001', 2, 'disponible'),
(10, 'TUNE-CENT-001', 1, 'disponible'),
(10, 'TUNE-NORT-001', 2, 'disponible');

INSERT INTO miembros (nombre, apellido, email, telefono, direccion, activo) VALUES
('Juan', 'Pérez', 'juan.perez@email.com', '555-1001', 'Calle 15 #25-40', true),
('María', 'González', 'maria.gonzalez@email.com', '555-1002', 'Carrera 20 #30-50', true),
('Carlos', 'Rodríguez', 'carlos.rodriguez@email.com', '555-1003', 'Avenida 25 #40-60', true),
('Ana', 'Martínez', 'ana.martinez@email.com', '555-1004', 'Calle 30 #35-45', true),
('Luis', 'Hernández', 'luis.hernandez@email.com', '555-1005', 'Carrera 40 #50-70', true);