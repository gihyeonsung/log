package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gihyeonsung/log/internal/application"
	"github.com/gihyeonsung/log/internal/infrastructure"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	configService := infrastructure.NewYamlConfigService("config.yaml")
	config, err := configService.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", config.Elasticsearch.Host, config.Elasticsearch.Port)},
	})
	if err != nil {
		log.Fatalf("create elasticsearch client: %v", err)
	}
	postDocumentRepository := infrastructure.NewEsPostDocumentRepository(esClient)

	sqliteDB, err := sql.Open("sqlite3", config.Sqlite.Path)
	if err != nil {
		log.Fatalf("open sqlite database: %v", err)
	}
	defer sqliteDB.Close()
	postRepository := infrastructure.NewSqlitePostRepository(sqliteDB)

	authnService := application.NewEnvVarAuthnService(config.AuthnService.Key)
	postCreate := application.NewPostCreate(authnService, postRepository)
	postDelete := application.NewPostDelete(authnService, postRepository)
	postDocumentSearch := application.NewPostDocumentSearch(postDocumentRepository)
	postFind := application.NewPostFind(postRepository)
	postUpdate := application.NewPostUpdate(authnService, postRepository)
	postDocumentSync := application.NewPostDocumentSync(postDocumentRepository, postRepository)

	mux := http.NewServeMux()
	infrastructure.NewPostController(mux, postCreate, postDelete, postDocumentSearch, postDocumentSync, postFind, postUpdate)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: mux,
	}

	log.Printf("listening %s", fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
