-- +migrate Up
-- Enable UUID extension for PostgreSQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Down
-- Note: Dropping extension might fail if tables are using UUID functions
DROP EXTENSION IF EXISTS "uuid-ossp";