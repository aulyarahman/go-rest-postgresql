CREATE TABLE IF NOT EXISTS items(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users(
    id_user SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    phone_number VARCHAR(18) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payment(
    id SERIAL PRIMARY KEY,
    id_user INT,
    amount INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS otp(
    id VARCHAR ,
    otp VARCHAR ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
)