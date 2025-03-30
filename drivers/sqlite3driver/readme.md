# Obtaining db schema

```sql
SELECT name FROM sqlite_master WHERE type='table' ORDER BY name;
-- sqlite_sequence
PRAGMA table_info(...);
PRAGMA index_list(...);
PRAGMA index_info(...);
```

# column types

- INTEGER
- TEXT
- REAL
- BLOB
- NUMERIC

# PK

```sql
SELECT last_insert_rowid();
```