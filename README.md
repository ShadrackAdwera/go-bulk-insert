# go-bulk-insert

Optimized insertion query using parameterized queries to insert records in batches into a PostgreSQL database

### Process

- Open file with multiple records
- Read json data
- Send this data to a redis queue
- Inserting records in batches of 1000
- Iterate through the records, creates a prepared statement, and execute it for each batch.
- The process continues until all records have been inserted.

### Technologies

- Gin
- Redis
- Postgres
