-- Script de inicialización para PostgreSQL
-- Este script se ejecuta automáticamente cuando se levanta el contenedor por primera vez

-- Habilitar extensión para UUIDs (opcional, por si decidimos usar UUIDs en el futuro)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Configurar zona horaria
SET timezone = 'UTC';
