PRAGMA foreign_keys = ON;


CREATE TABLE  if not exists users (
    CodeNumber TEXT PRIMARY KEY,
    NAME TEXT NOT NULL,
    DOI TEXT,
    GEMI TEXT,
    Phone TEXT,
    Mobile_Phone TEXT,
    Email TEXT,
    -- these are for the postaladdress struct
    PostalCellName TEXT,
    PostalCellNumber TEXT,
    PostalCellPostalCode TEXT,
    PostalCellCity TEXT,
    -- these are for the real address struct
    AddStreet TEXT,
    AddNumber TEXT,
    AddPostalCode TEXT,
    AddCity TEXT,
    VatNumber TEXT NOT NULL UNIQUE,
    Country VARCHAR(2),
    Branch   INTEGER
);


CREATE TABLE  if not exists customers (
    CodeNumber TEXT PRIMARY KEY,
    NAME TEXT NOT NULL,
    DOI TEXT,
    GEMI TEXT,
    Phone TEXT,
    Mobile_Phone TEXT,
    Email TEXT,
    -- these are for the postaladdress struct
    PostalCellName TEXT,
    PostalCellNumber TEXT,
    PostalCellPostalCode TEXT,
    PostalCellCity TEXT,
    -- these are for the real address struct
    AddStreet TEXT,
    AddNumber TEXT,
    AddPostalCode TEXT,
    AddCity TEXT,
    VatNumber TEXT NOT NULL UNIQUE,
    Country VARCHAR(2),
    Branch   INTEGER,
    -- these are attributes only for the customers
    Balance REAL  CHECK (Balance = round(Balance, 2)),
    Discount INTEGER CHECK (Discount BETWEEN 1 and 100)
);

CREATE TABLE  IF NOT EXISTS USERS_BANK_ACCOUNTS (
    USERCODE TEXT NOT NULL,
    BANK_NAME TEXT NOT NULL,
    IBAN TEXT NOT NULL UNIQUE,
    FOREIGN KEY (USERCODE)
        REFERENCES users(CodeNumber)
        ON DELETE CASCADE
);

-- we might not need that table
CREATE TABLE  IF NOT EXISTS CUSTOMER_BANK_ACCOUNTS (
    CUSTOMERCODE TEXT NOT NULL,
    BANK_NAME TEXT NOT NULL,
    IBAN TEXT NOT NULL UNIQUE,
    FOREIGN KEY (CUSTOMERCODE)
        REFERENCES customers(CodeNumber)
        ON DELETE CASCADE
);


CREATE TABLE  if not exists BranchCompanies (
    BranchCode TEXT PRIMARY KEY,
    CompanyCode TEXT NOT NULL,
    NAME TEXT NOT NULL,
    Phone TEXT,
    Mobile_Phone TEXT,
    Email TEXT,
    -- these are for the real address struct
    AddStreet TEXT,
    AddNumber TEXT,
    AddPostalCode TEXT,
    AddCity TEXT,
    Country VARCHAR(2),
    Branch   INTEGER,
    -- these are attributes only for the customers
    Balance REAL NOT NULL CHECK (Balance = round(Balance, 2)),
    Discount INTEGER CHECK (Discount BETWEEN 1 and 100)
);






CREATE TABLE if not exists products (
    CodeNumber TEXT PRIMARY KEY,
    name            TEXT NOT NULL,
    description     TEXT,
    unit_net_price      REAL NOT NULL CHECK (unit_net_price >= 0),
    measurmentUnit INTEGER NOT NULL REFERENCES measurementUnits(id),
    vat_category INTEGER NOT NULL references vat_categories(id)
);



CREATE TABLE IF NOT EXISTS measurementUnits (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    unit TEXT NOT NULL UNIQUE
);

INSERT INTO measurementUnits (unit) VALUES
('Τεμάχια'),
('Κιλά'),
('Λίτρα'),
('Μέτρα'),
('Τετραγωνικά Μέτρα'),
('Κυβικά Μέτρα'),
('Τεμάχια_Λοιπές Περιπτώσεις')
;





CREATE TABLE IF NOT EXISTS vat_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    rate REAL CHECK(rate >= 0)
);

INSERT INTO vat_categories(name, rate) VALUES
('ΦΠΑ ΣΥΝΤΕΛΕΣΤΗΣ 24%', 24),
('ΦΠΑ ΣΥΝΤΕΛΕΣΤΗΣ 13%', 13),
('ΦΠΑ ΣΥΝΤΕΛΕΣΤΗΣ 6%', 6),
('ΦΠΑ ΣΥΝΤΕΛΕΣΤΗΣ 17%', 17),
('ΦΠΑ ΣΥΝΤΕΛΕΣΤΗΣ 9%', 9),
('ΦΠΑ ΣΥΝΤΕΛΕΣΤΗΣ 4%', 4),
('Ανευ ΦΠΑ', 0),
('Εγγραφές χωρίς ΦΠΑ (πχ Μισθοδοσία, Αποσβέσεις) ',0 ),
('ΦΠΑ συντελεστής 3% (αρ.31 ν.5057/2023)',3 ),
('ΦΠΑ συντελεστής 4% (αρ.31 ν.5057/2023)',4 );


CREATE TABLE  if not exists categoriesforproducts  (
    id              integer NOT NULL PRIMARY KEY autoincrement,
    name            TEXT NOT NULL UNIQUE,
    description     TEXT
);

CREATE TABLE  if not exists product_categories  (
    product_id TEXT NOT NULL ,
    category_id integer NOT NULL ,
    FOREIGN KEY (product_id) REFERENCES products(CodeNumber) on delete cascade,
    FOREIGN KEY (category_id) REFERENCES categoriesforproducts(id) on delete cascade
);


