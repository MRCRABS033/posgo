CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT
);

CREATE TABLE IF NOT EXISTS permissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL UNIQUE,
    create_product BOOLEAN NOT NULL DEFAULT 0,
    delete_product BOOLEAN NOT NULL DEFAULT 0,
    modified_unit_cost_price_product BOOLEAN NOT NULL DEFAULT 0,
    modified_unit_sell_price_product BOOLEAN NOT NULL DEFAULT 0,
    modified_stock_product BOOLEAN NOT NULL DEFAULT 0,
    modified_discount BOOLEAN NOT NULL DEFAULT 0,
    modified_product_name BOOLEAN NOT NULL DEFAULT 0,
    modified_available_discount BOOLEAN NOT NULL DEFAULT 0,
    modified_department BOOLEAN NOT NULL DEFAULT 0,
    modified_user_name BOOLEAN NOT NULL DEFAULT 0,
    modified_user_last_name BOOLEAN NOT NULL DEFAULT 0,
    modified_user_phone_number BOOLEAN NOT NULL DEFAULT 0,
    modified_user_permissions BOOLEAN NOT NULL DEFAULT 0,
    modified_user_password BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS departments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    department_id INT NOT NULL,
    code TEXT NOT NULL,
    product_name TEXT NOT NULL,
    unit_cost_price REAL NOT NULL,
    unit_sell_price REAL NOT NULL,
    discount REAL,
    is_single_product BOOLEAN NOT NULL DEFAULT 1,
    available_discount BOOLEAN NOT NULL DEFAULT 0,
    stock REAL NOT NULL,
    available BOOLEAN NOT NULL DEFAULT 1,
    FOREIGN KEY (department_id) REFERENCES departments(id)
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    uuid TEXT NOT NULL UNIQUE,
    login_at DATETIME NOT NULL,
    logout_at DATETIME,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS cash_ins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    session_id INTEGER NOT NULL,
    concept TEXT NOT NULL,
    quantity REAL NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (session_id) REFERENCES sessions(id)
);

CREATE TABLE IF NOT EXISTS cash_out (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    session_id INTEGER NOT NULL,
    concept TEXT NOT NULL,
    quantity REAL NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (session_id) REFERENCES sessions(id)
);

CREATE TABLE IF NOT EXISTS tickets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    session_id INTEGER NOT NULL,
    create_at DATETIME,
    total REAL NOT NULL,
    is_completed BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY (session_id) REFERENCES sessions(id)
);

CREATE TABLE IF NOT EXISTS ticket_items (
    ticket_id INT NOT NULL,
    session_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity REAL NOT NULL,
    subtotal REAL NOT NULL,
    FOREIGN KEY (ticket_id) REFERENCES tickets(id),
    FOREIGN KEY (session_id) REFERENCES sessions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);