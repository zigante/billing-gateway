CREATE TABLE TRANSACTIONS (
    ID            TEXT NOT NULL,
    ACCOUNT_ID    TEXT NOT NULL,
    AMOUNT        REAL NOT NULL,
    STATUS        TEXT NOT NULL,
    ERROR_MESSAGE TEXT NOT NULL,
    CREATED_AT    TEXT NOT NULL,
    UPDATED_AT    TEXT NOT NULL
);