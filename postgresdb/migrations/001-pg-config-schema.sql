-- Table db_config holds key-value pairs, including 'schema_version'.
CREATE TABLE db_config (
    "key" TEXT NOT NULL PRIMARY KEY,
    "value" TEXT
);

INSERT INTO db_config (key, value) VALUES ('schema_version', '1');

-- start_schema_update must be called by db update scripts.
-- An exception is raised if next_version is not 1 greater than
-- the value of config.schema_version.
create or replace procedure start_schema_update(next_version int)
    language plpgsql
as $$
declare current_version int;
begin
    select value into current_version from db_config where key = 'schema_version';
    raise info 'starting schema update: % --> %', current_version, next_version;

    if next_version != current_version+1 then
        raise exception
            'requested next version must be the next integer value, but was: % --> %', current_version, next_version;
    end if;

    UPDATE db_config SET value = next_version WHERE key = 'schema_version';
end;$$;