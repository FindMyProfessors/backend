package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/FindMyProfessors/backend/pagination"
	"github.com/FindMyProfessors/backend/repository/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/FindMyProfessors/backend/graph"
	"github.com/FindMyProfessors/backend/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pool, err := ConnectWithRetries(GetEnvOrDie("DATABASE_URI"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	config := generated.Config{
		Resolvers: &graph.Resolver{
			Repository: &database.Repository{DatabasePool: pool},
		},
		Directives: generated.DirectiveRoot{Pagination: pagination.Pagination},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	srv.SetRecoverFunc(func(ctx context.Context, iErr interface{}) error {
		err := fmt.Errorf("%v", iErr)

		log.Printf("runtime error: %v\n\n%v\n", err, string(debug.Stack()))

		return gqlerror.Errorf("Internal server error! Check logs for more details!")
	})
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Println("Error presented: ", err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ConnectWithRetries(databaseUri string) (pool *pgxpool.Pool, err error) {
	for i := 0; i < 10; i++ {
		pool, err = pgxpool.New(context.Background(), databaseUri)
		if err == nil {
			return pool, nil
		}
		time.Sleep(time.Second * 1)
	}
	return nil, err
}

func GetEnvOrDie(key string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("You must provide the %s environmental variable\n", key)
	}
	return env
}
