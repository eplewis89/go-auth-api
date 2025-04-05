create table Users
(
    id   INTEGER PRIMARY KEY,
    first_name text    NOT NULL,
    last_name text    NOT NULL,
    email text    NOT NULL,
    email_activated boolean BOOLEAN DEFAULT(FALSE),
    encr_password text,
    salt_password text,
    temp_password text,
    is_deleted boolean BOOLEAN DEFAULT(FALSE)
);

create table AccessTokens
(
    id   INTEGER PRIMARY KEY,
    expiration_time integer not null,
    token text,
    user_id integer not null
);

create table EmailActivationCodes
(
    id   INTEGER PRIMARY KEY,
    activation_code integer not null,
    expiration_time integer not null,
    user_id integer not null
);