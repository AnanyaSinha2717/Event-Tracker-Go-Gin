CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    password TEXT NOT NULL
);

-- TODO: solve error: 
-- sql: unknown driver "sqlite3/" (forgotten import?)
-- exit status 1