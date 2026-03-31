package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Jnmendza/uh/backend/graph"
	"github.com/Jnmendza/uh/backend/graph/generated"
	"github.com/Jnmendza/uh/backend/internal/db"
	"github.com/Jnmendza/uh/backend/internal/repository"
)

func main() {
	log.Println("Starting Union Hub Backend Server...")

	// 1. Initialize Database Connection
	client := db.ConnectDB()

	// Ensure we clean up the connection on exit
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
		log.Println("Gracefully disconnected from MongoDB.")
	}()

	// 2. Initialize Repositories
	orgRepo := repository.NewOrganizationRepo(client)
	userRepo := repository.NewUserRepo(client)
	membershipRepo := repository.NewMembershipRepo(client)
	membershipTierRepo := repository.NewMembershipTierRepo(client)

	// 3. Setup GraphQL Resolver
	resolver := &graph.Resolver{
		OrgRepo:            orgRepo,
		UserRepo:           userRepo,
		MembershipRepo:     membershipRepo,
		MembershipTierRepo: membershipTierRepo,
	}

	// 4. Configure GraphQL Handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// 5. Mount Routes
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// 6. Start Server
	log.Println("✅ Database connection is up!")
	log.Printf("🚀 Connect to http://localhost:8080/ for GraphQL playground\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
