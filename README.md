# Union Hub

Union Hub is a centralized platform currently under active development. This repository operates as a monorepo containing both the frontend and backend applications, designed with modern web architecture stack to provide a scalable and robust experience.

## Project Architecture

This project uses a monorepo architecture leveraging `go.work` for backend multi-module management.

- **`/backend`**: The Go-based API layer. Built with highly typed GraphQL (`gqlgen`) and MongoDB.
- **`/frontend`**: The user interface layer (pending integration structure).

## Prerequisite Tooling

- [Go 1.25+](https://golang.org/dl/)
- MongoDB Connection URI (Local or Atlas)
- Node.js & npm (for frontend dependencies)

## Global Setup

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd union-hub
   ```
2. **Environment Variables:**
   - Both the frontend and backend have their own independent `.env` structures. 
   - Never commit sensitive keys; check `.gitignore` before caching new configuration files.

## Getting Started

Visit the respective subdirectories for instructions on starting the independent servers.
- [Backend Documentation](./backend/README.md)
- Frontend Documentation (Coming soon!)
