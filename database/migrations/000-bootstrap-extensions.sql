BEGIN TRANSACTION;

-- We need pgcrypto for UUID support.
CREATE EXTENSION "pgcrypto";

COMMIT;
