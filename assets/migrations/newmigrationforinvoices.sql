CREATE TABLE IF NOT EXISTS SellerCompanies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    entity_type INTEGER NOT NULL,
    branch INTEGER NOT NULL,
    vat_number TEXT NOT NULL,
    country TEXT NOT NULL,
    street TEXT,
    number TEXT,
    postal_code TEXT,
    city TEXT
);

--create table entity-types

--maybe create table branch int-description

--create table vat_categories

CREATE TABLE IF NOT EXISTS BuyerCompanies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    entity_type INTEGER NOT NULL,
    branch INTEGER NOT NULL,
    vat_number TEXT NOT NULL,
    country TEXT NOT NULL,
    street TEXT,
    number TEXT,
    postal_code TEXT,
    city TEXT
);

CREATE TABLE IF NOT EXISTS Products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    default_unit_price REAL NOT NULL,
    vat_category INTEGER NOT NULL,
    --standard 24%, reduced 13%, super-reduced 6%, zero-rated 0%
    sku TEXT UNIQUE,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

