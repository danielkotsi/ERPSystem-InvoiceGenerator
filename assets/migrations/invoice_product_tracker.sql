PRAGMA foreign_keys = ON;

CREATE TABLE  if not exists companies (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    code            TEXT NOT NULL,
    name            TEXT NOT NULL,
    address_line1   TEXT,
    address_num1   int,
    address_line2   TEXT,
    address_num2   int,
    city            TEXT,
    state           TEXT,
    postal_code     TEXT,
    country         TEXT,
    email           TEXT,
    phone           TEXT,
    mobile_phone           TEXT,
    tax_id          TEXT,          -- VAT, EIN, etc.
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE  if not exists categoriesforproducts  (
    id              integer NOT NULL PRIMARY KEY autoincrement,
    name            TEXT NOT NULL UNIQUE,
    description     TEXT
);

CREATE TABLE  if not exists product_categories  (
    id              integer NOT NULL PRIMARY KEY autoincrement,
    product_id TEXT NOT NULL ,
    category_id TEXT NOT NULL ,
    FOREIGN KEY (product_id) REFERENCES products(id) on delete cascade
    FOREIGN KEY (category_id) REFERENCES categoriesforproducts(id) 
);

CREATE TABLE if not exists products (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name            TEXT NOT NULL,
    description     TEXT,
    sku             TEXT UNIQUE,        -- Optional stock-keeping code
    unit_price      REAL NOT NULL CHECK (unit_price >= 0),
    currency        TEXT NOT NULL DEFAULT 'EUR',
    active          INTEGER NOT NULL DEFAULT 1,  -- 1 = active, 0 = discontinued

    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (category_id) REFERENCES product_categories(id)
);

CREATE TABLE  if not exists invoices (
    id                  TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    invoice_number      TEXT NOT NULL UNIQUE,
    seller_id           TEXT NOT NULL,
    customer_id         TEXT NOT NULL,
    issue_date          DATE NOT NULL,
    due_date            DATE NOT NULL,
    status              TEXT NOT NULL DEFAULT 'pending', 
    currency            TEXT NOT NULL DEFAULT 'USD',
    notes               TEXT,

    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (seller_id) REFERENCES companies(id),
    FOREIGN KEY (customer_id) REFERENCES companies(id)
);

CREATE TABLE if not exists invoice_items (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    invoice_id      TEXT NOT NULL,
    product_id      TEXT,               -- optional: allow custom lines without products
    description     TEXT NOT NULL,
    quantity        REAL NOT NULL CHECK (quantity > 0),
    unit_price      REAL NOT NULL CHECK (unit_price >= 0),
    total           REAL NOT NULL,      -- quantity * unit_price, stored for stability

    FOREIGN KEY (invoice_id) REFERENCES invoices(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

