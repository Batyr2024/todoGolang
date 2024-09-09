CREATE TABLE tasks
(
    id      int     NOT NULL UNIQUE,
    text    TEXT NOT NULL,
    checked BOOLEAN DEFAULT(false),
    PRIMARY KEY (id)
)