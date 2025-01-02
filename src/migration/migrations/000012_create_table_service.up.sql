CREATE TABLE services (
    id serial PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    version VARCHAR(255) NOT NULL,
    connection json NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)