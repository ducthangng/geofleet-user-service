package postgresql

/* To follow Best Practices with sqlc:
- Schema First: sqlc generates Go structs from your SQL schema. You typically do not define the struct manually.
- Use pgx/v5: This is the most performant PostgreSQL driver for Go.
- Atomic Transactions: Always pass context to ensure requests can be cancelled.
- Type Mapping: Your struct uses string for DateCreated. Best practice is to use TIMESTAMPTZ in Postgres and time.Time in Go.
*/
