package main

import (
    "net/http"
    "os"
    "log"
    "os/signal"
    "syscall"
    "github.com/spf13/pflag"
    "go.uber.org/zap"
    "github.com/myk4040okothogodo/EvolveAPI/pkg/api"
    "github.com/myk4040okothogodo/EvolveAPI/pkg/db"
    "github.com/myk4040okothogodo/EvolveAPI/docs"
    "github.com/joho/godotenv"

  )

const defaultConnection ="host=localhost port=5432 user=postgres dbname=evolveapi password=mykokothe sslmode=disable"



var ( 
    version string
    addr string
)

func parseEnvVars(key string) string {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
      log.Print("Error loading DB_CONNECTION Variable from.env file :")
    }

    return os.Getenv(key)
   }

func init() {
    pflag.StringVarP(&addr, "address", "a", ":8080", "the address for the api to listen on. Host and Port separated by ':'")
    pflag.Parse()
}


func main() {

    // configure logger
    log, _ := zap.NewProduction(zap.WithCaller(false))
    defer func() {
        _ = log.Sync()
    }()    
   
    docs.SwaggerInfo.Version = version

    //gracefully exit on keyboard interupt
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
 
    // print current version
    log.Info("spinning up the Users API .....", zap.String("version :", version))

    dbClient := &db.Client{}
    connectionStr := parseEnvVars("DB_CONNECTION")
    if connectionStr == "" {
        log.Info("Failed to load 'DATABASE_CONNECTION' from environment variables , system is reverting to ", zap.String("default", defaultConnection))
        connectionStr = defaultConnection
    }
    if err := dbClient.Connect(connectionStr); err != nil {
        log.Error("Couldnt connect to database", zap.Error(err))
        os.Exit(1)
    }

    // Start the api server
    //
    r := api.GetRouter(log, dbClient)
    go func() {
        if err := http.ListenAndServe(addr, r); err != nil {
            log.Error("failed to start server", zap.Error(err))
            os.Exit(1)
        }
    }()

    log.Info("Ready to Serve Requests on :" + addr)
    <-c
    log.Info("Gracefully Shutting Down")
    os.Exit(0)
}

