PRAGMA foreign_keys = ON;

CREATE TABLE  if not exists companies (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    -- code            TEXT NOT NULL,
    name            TEXT NOT NULL,
    entity_type INTEGER NOT NULL REFERENCES entity_types(code) ON DELETE RESTRICT 
    ON UPDATE CASCADE,
    branch INTEGER NOT NULL DEFAULT 0,
    vat_number TEXT NOT NULL unique,
    address_street   TEXT,
    address_number   text,
    city            TEXT,
    postal_code     TEXT,
    country         TEXT,
    email           TEXT,
    phone           TEXT,
    mobile_phone           TEXT,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS entity_types(
    id              integer NOT NULL PRIMARY KEY autoincrement,
    code integer not null unique,
    name text not null
);


INSERT INTO entity_types(code,name) values (1,'Business'),(2,'Private Individual'),(3,'Public Sector Entity'),(4,'Foreign Entity'),(5,'Non-Profit Organization'),(6,'Intra-EU VAT Registered Entity VIES');




CREATE TABLE  if not exists categoriesforproducts  (
    id              integer NOT NULL PRIMARY KEY autoincrement,
    name            TEXT NOT NULL UNIQUE,
    description     TEXT
);

CREATE TABLE  if not exists product_categories  (
    product_id TEXT NOT NULL ,
    category_id integer NOT NULL ,
    FOREIGN KEY (product_id) REFERENCES products(id) on delete cascade,
    FOREIGN KEY (category_id) REFERENCES categoriesforproducts(id) on delete cascade
);




CREATE TABLE if not exists products (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name            TEXT NOT NULL,
    description     TEXT,
    sku             TEXT UNIQUE,
    unit_price      REAL NOT NULL CHECK (unit_price >= 0),
    active          INTEGER NOT NULL DEFAULT 1, 
    vat_category INTEGER NOT NULL references vat_categories(id),
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS vat_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    rate REAL CHECK(rate >= 0)
);

INSERT INTO vat_categories(name, rate) VALUES
('Standard', 24),
('Reduced', 13),
('Super-reduced', 6),
('Exempt', 0);




CREATE TABLE IF NOT EXISTS invoices (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    seller_id TEXT NOT NULL REFERENCES companies(id) ON DELETE CASCADE, -- Your company issuing the invoice
    buyer_id TEXT NOT NULL REFERENCES companies(id), -- Or separate customers table
    series text,
    aa text,
    invoice_type text,
    invoice_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date DATETIME,
    currency TEXT NOT NULL DEFAULT 'EUR',
    total_net REAL NOT NULL CHECK(total_net >= 0),
    total_vat REAL NOT NULL CHECK(total_vat >= 0),
    total_with_vat REAL NOT NULL CHECK(total_with_vat >= 0),
    status TEXT NOT NULL DEFAULT 'Draft', -- Draft, Issued, Paid
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS invoice_lines (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    invoice_id TEXT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE, -- Belongs to which invoice
    line_number INTEGER NOT NULL,
    rec_type INTEGER DEFAULT 3,
    net_value REAL NOT NULL CHECK(net_value >= 0),
    vat_amount REAL NOT NULL CHECK(vat_amount >= 0),
    product_id TEXT REFERENCES products(id), -- Optional: link to a product
    description TEXT, -- Free text description of product/service
    quantity REAL NOT NULL CHECK(quantity > 0),
    unit_price REAL NOT NULL CHECK(unit_price >= 0),
    vat_category INTEGER NOT NULL REFERENCES vat_categories(id), -- VAT category
    line_total REAL NOT NULL CHECK(line_total >= 0), -- Quantity * Unit Price + VAT
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS invoice_payments (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    invoice_id TEXT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    type INTEGER NOT NULL, -- 1=Cash, 2=Bank, 3=Card, 4=POS
    amount REAL NOT NULL CHECK(amount >= 0),
    tid TEXT -- optional POS terminal ID
);

CREATE TABLE IF NOT EXISTS invoice_classifications (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    invoice_id TEXT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    classification_type TEXT NOT NULL,
    classification_category TEXT NOT NULL,
    amount REAL NOT NULL CHECK(amount >= 0)
);
