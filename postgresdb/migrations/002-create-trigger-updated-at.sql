BEGIN TRANSACTION;
LOCK TABLE db_config IN EXCLUSIVE MODE;
CALL start_schema_update(2);

-- trigger_updated_at updates the updated_at column
-- to the current time. This ensures that the column is always
-- updated, even when working directly using SQL, or if the calling
-- code forgets to update it.
CREATE OR REPLACE FUNCTION trigger_updated_at()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS
$$
BEGIN
    NEW.updated_at = timezone('utc', now());
    RETURN NEW;
END;
$$
;

COMMIT;