CREATE TABLE IF NOT EXISTS triangle
(
    id CHARACTER varying(36) NOT NULL,
    side1 INTEGER NOT NULL,
    side2 INTEGER NOT NULL,
    side3 INTEGER NOT NULL,
    type CHARACTER varying(100) NOT NULL,
    PRIMARY KEY (id)
);