PRAGMA foreign_keys = ON;

-- ============================================
--  Companies table (seller + customers)
-- ============================================
CREATE TABLE companies (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name            TEXT NOT NULL,
    address_line1   TEXT,
    address_line2   TEXT,
    city            TEXT,
    state           TEXT,
    postal_code     TEXT,
    country         TEXT,
    email           TEXT,
    phone           TEXT,
    tax_id          TEXT,          -- VAT, EIN, etc.
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
--  Product Categories (optional but useful)
-- ============================================
CREATE TABLE product_categories (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name            TEXT NOT NULL UNIQUE,
    description     TEXT
);

-- ============================================
--  Products (available for sale)
-- ============================================
CREATE TABLE products (
    id              TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    category_id     TEXT,
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

-- Trigger: auto-update timestamp
CREATE TRIGGER update_products_updated_at
AFTER UPDATE ON products
BEGIN
    UPDATE products SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;

-- ============================================
--  Invoices table
-- ============================================
CREATE TABLE invoices (
    id                  TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    invoice_number      TEXT NOT NULL UNIQUE,
    seller_id           TEXT NOT NULL,
    customer_id         TEXT NOT NULL,
    issue_date          DATE NOT NULL,
    due_date            DATE NOT NULL,
    status              TEXT NOT NULL DEFAULT 'pending', 
        -- statuses: pending, paid, cancelled

    currency            TEXT NOT NULL DEFAULT 'USD',
    notes               TEXT,

    created_at          DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (seller_id) REFERENCES companies(id),
    FOREIGN KEY (customer_id) REFERENCES companies(id)
);

-- Trigger: auto-update timestamp
CREATE TRIGGER update_invoices_updated_at
AFTER UPDATE ON invoices
BEGIN
    UPDATE invoices SET updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.id;
END;

-- ============================================
--  Invoice line items
--  (can reference products OR custom description)
-- ============================================
CREATE TABLE invoice_items (
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

-- ============================================
--  Payments for invoic

