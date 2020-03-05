package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/viper"

	articleHandler "github.com/golang/learn-blog/handlers/article"

	"github.com/golang/learn-blog/usecases/article"

	"github.com/golang/learn-blog/domains/repository"

	"github.com/gorilla/mux"
)

var articleRepo repository.ArticleRepository
var addArticleUsecase article.AddArticleUsecase
var updateArticleUsecase article.UpdateArticleUsecase
var fetchArticleUsecase article.FetchArticleUsecase
var deleteArticleUsecase article.DeleteArticleUsecase
var findArticleUsecase article.FindArticleUsecase
var addArticleHandler articleHandler.AddHandler

func init() {
	viper.SetConfigFile(`app.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	db := dbconfig()

	articleRepo = repository.NewArticleRepository(db)
	addArticleUsecase = article.NewAddArticleUsecase(articleRepo)
	updateArticleUsecase = article.NewUpdateArticleUsecase(articleRepo)
	fetchArticleUsecase = article.NewFetchArticleUsecase(articleRepo)
	findArticleUsecase = article.NewFindArticleUsecase(articleRepo)
	deleteArticleUsecase = article.NewDeleteArticleUsecase(articleRepo)

	addArticleHandler = articleHandler.NewAddHandler(addArticleUsecase)
}

func dbconfig() *sql.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//lazy, menunda sampai object ga kepake
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return dbConn
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	return r
}

func main() {

	r := newRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
