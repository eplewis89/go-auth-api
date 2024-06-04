BEGIN TRANSACTION;
LOCK TABLE db_config IN EXCLUSIVE MODE;

CALL start_schema_update(6);

CREATE TABLE AccessTokens
(
    id serial not null
        constraint AccessTokens_pkey
            primary key,
    expiration_time timestamp default now() not null,
    token varchar(255) not null default '',
    user_id integer
);

alter table AccessTokens owner to goauth;

-- create foreign keys

alter table AccessTokens
    add constraint AccessTokens_user_id_fkey
        foreign key (user_id) references Users (id)
            on delete cascade;

COMMIT;