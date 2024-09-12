CREATE TABLE tasks
(
    id    serial NOT NULL UNIQUE,
    text    TEXT NOT NULL,
    checked BOOLEAN DEFAULT(false),
    PRIMARY KEY (id)

)