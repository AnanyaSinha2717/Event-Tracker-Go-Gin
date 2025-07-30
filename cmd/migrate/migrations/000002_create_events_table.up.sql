CREATE TABLE IF NOT EXISTS events(
    id INT PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    date TEXT NOT NULL,
    location TEXT NOT NULL,
    type TEXT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);