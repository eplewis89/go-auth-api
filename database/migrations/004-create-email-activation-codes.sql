BEGIN TRANSACTION;
LOCK TABLE db_config IN EXCLUSIVE MODE;

CALL start_schema_update(5);

create table EmailActivationCodes
(
    id serial not null
        constraint EmailActivationCodes_pkey
            primary key,
    activation_code varchar(100) not null default '',
    expiration_time timestamp default (now() at time zone 'utc') not null,
    user_id integer not null default 0
);

alter table EmailActivationCodes owner to edgeworx;

-- create foreign keys

alter table EmailActivationCodes
    add constraint EmailActivationCodes_user_id_fkey
        foreign key (user_id) references Users (id)
            on delete cascade;

COMMIT;