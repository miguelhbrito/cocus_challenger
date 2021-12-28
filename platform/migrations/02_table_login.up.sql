CREATE TABLE IF NOT EXISTS login
(
    username CHARACTER varying(36) NOT NULL,
    password CHARACTER varying(300) NOT NULL,
    PRIMARY KEY (username)
);
