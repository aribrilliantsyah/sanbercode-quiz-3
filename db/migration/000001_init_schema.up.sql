CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR,
    image_url VARCHAR,
    release_year INTEGER,
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR,
    category_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR,
    CONSTRAINT fk_category
      FOREIGN KEY (category_id) 
      REFERENCES categories (id)
      ON DELETE CASCADE
)

