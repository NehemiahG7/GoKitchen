\c GoKitchen
CREATE TABLE accounts(
    ID serial NOT NULL PRIMARY KEY,
    username VARCHAR (64) NOT NULL,
    password VARCHAR (64) NOT NULL,
    items json NOT NULL
);