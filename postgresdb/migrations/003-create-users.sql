BEGIN TRANSACTION;
LOCK TABLE db_config IN EXCLUSIVE MODE;

CALL start_schema_update(3);

create table Users
(
    id serial not null
        constraint Users_pkey
            primary key,
    first_name varchar(100) not null default '',
    last_name varchar(100) not null default '',
    email varchar(100) unique not null default '',
    email_activated boolean not null default FALSE,
    encr_password varchar(100),
    salt_password varchar(100),
    temp_password varchar(100),
    created_at timestamp default (now() at time zone 'utc') not null,
    updated_at timestamp default null,
    deleted_at timestamp default null,
    is_deleted boolean not null default FALSE
);

alter table Users owner to goauth;

-- create updated by trigger

create trigger users_trigger_updated_at
    before update on Users for each row
    execute procedure trigger_updated_at();

COMMIT;
