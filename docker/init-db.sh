#!/bin/sh -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    SELECT 'CREATE DATABASE "orderfaz"' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname='orderfaz')\gexec
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
