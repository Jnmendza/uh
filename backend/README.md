# Union Hub Backend API

The Union Hub Backend is a robust, strongly-typed GraphQL API built with **Go (1.25.0)**, leveraging **gqlgen** for schema-first code generation, and driven by **MongoDB** for flexible persistence.

## Tech Stack Overview
- **Language**: Go 1.25.0
- **API Protocol**: GraphQL
- **Generator framework**: [99designs/gqlgen](https://gqlgen.com/)
- **Database**: MongoDB (Local or Atlas)
- **Dependency Management**: Go Modules (`go.mod`)

## Directory Map

- `/cmd/api/` - The entry point for the application. Starts the HTTP server and wires the database connections and repositories into the GraphQL query handlers.
- `/graph/` - Contains the raw GraphQL schema files (`schema.graphqls`), alongside auto-generated structs and the manually implemented resolver logic mappings.
- `/internal/models/` - BSON-tagged Go structs mapping natively to MongoDB documents, serving as the system's core entities.
- `/internal/repository/` - Database logic, wrapping standard raw MongoDB queries. Injected into the GraphQL `Resolver` for query operations.
- `/internal/db/` - Manages database connections and MongoDB client polling.

## Quick Start Configuration

1. Make sure you are in the `backend/` directory:
   ```bash
   cd backend
   ```
2. Create your `.env` configuration file in the `backend/` root (this is excluded from git):
   ```env
   MONGO_URI=mongodb://localhost:27017 # Your DB string
   PORT=8080                          # Optional, default is 8080
   ```
3. Boot the API:
   ```bash
   go run cmd/api/main.go
   ```
4. Access the GraphQL Playground in your web browser:
   `http://localhost:8080/`

## Modifying the GraphQL Schema

This project relies on **gqlgen** for type-safety and boilerplate GraphQL generation. We operate on a 'Schema-First' approach.

1. Modify your raw schema definitions in `graph/schema.graphqls`.
2. Generate the corresponding Go boilerplate bindings by running:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```
3. Update `graph/schema.resolvers.go` manually if `gqlgen` drops any new unimplemented function stubs at the bottom of the file!

*Note: You do not need to install `gqlgen` globally; `go run` leverages the project's tracked module perfectly.*
