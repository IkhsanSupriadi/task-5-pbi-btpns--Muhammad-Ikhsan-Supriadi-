#!/bin/sh -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    SELECT 'CREATE DATABASE "test_orderfaz"' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname='test_orderfaz')\gexec
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
