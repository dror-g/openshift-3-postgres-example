package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"flag"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

var dbConn *sql.DB

var HOST string
var PORT string 
var USER string
var PASS string

func init() {
	fmt.Println("init called to open dbConn")
  
	flag.StringVar(&HOST, "h", "", "postgresql host to connect to")
	flag.StringVar(&PORT, "p", "", "postgresql port to listen on")
	flag.StringVar(&USER, "u", "", "postgresql user to connect as")
	flag.StringVar(&PASS, "w", "", "postgresql user password to connect with")
        flag.Parse()

	fmt.Println("connecting to postgres HOST=" + HOST + " PORT=" + PORT + " USER=" + USER + " PASSWORD=" + PASS)
	var err error
	dbConn, err = sql.Open("postgres", "password=" + PASS + " sslmode=disable user="+USER+" host="+HOST+" port="+PORT+" dbname="+USER)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	handler := rest.ResourceHandler{
		PreRoutingMiddlewares: []rest.Middleware{
			&rest.CorsMiddleware{
				RejectNonCorsRequests: false,
				OriginValidator: func(origin string, request *rest.Request) bool {
					return true
				},
				AllowedMethods: []string{"DELETE", "GET", "POST", "PUT"},
				AllowedHeaders: []string{
					"Accept", "Content-Type", "X-Custom-Header", "Origin"},
				AccessControlAllowCredentials: true,
				AccessControlMaxAge:           3600,
			},
		},
		EnableRelaxedContentType: true,
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/test", GetAllDatabases},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":12000", &handler))
}


type DBDatabase struct {
	Name        string
}

func GetAllDatabases(w rest.ResponseWriter, r *rest.Request) {
	var rows *sql.Rows
	var err error
	rows, err = dbConn.Query("select datname from pg_stat_database order by datname")
	if err != nil {
                rest.Error(w, err.Error(), 400)
	}
	defer rows.Close()
	databases := make([]DBDatabase, 0)
	for rows.Next() {
		database := DBDatabase{}
		if err = rows.Scan(&database.Name); err != nil {
                	rest.Error(w, err.Error(), 400)
		}
		databases = append(databases, database)
	}
	if err = rows.Err(); err != nil {
                rest.Error(w, err.Error(), 400)
	}

        w.WriteJson(&databases)
}

