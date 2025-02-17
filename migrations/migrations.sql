CREATE TABLE urls (
    short_code VARCHAR(10) PRIMARY KEY,
    original_url TEXT NOT NULL
);

CREATE INDEX idx_short_code ON urls(short_code);