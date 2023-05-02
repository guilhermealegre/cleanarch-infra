-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- create role app_auth
CREATE TABLE auth.users (
    id_users serial primary key,
    email varchar(300) NOT NULL,
    "password" varchar(500) NOT NULL
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE auth.users;