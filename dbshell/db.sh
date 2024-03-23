#!/bin/bash
set -e
export PGPASSWORD=postgres123;
psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "userdb" <<-EOSQL
  CREATE DATABASE userdb;
  GRANT ALL PRIVILEGES ON DATABASE userdb TO "postgres";
  \c userdb;  -- Switch to the new database
  CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;  -- Create the extension
  ALTER EXTENSION plpgsql SET SCHEMA pg_catalog;  -- Set the extension's schema
EOSQL
