// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"fmt"
	"github.com/ra9dev/safe-and-sound/internal/configs"
	"github.com/ra9dev/safe-and-sound/internal/police-server/daemons"
	"github.com/ra9dev/safe-and-sound/internal/police-server/http"
	"github.com/ra9dev/safe-and-sound/pkg/config"
	"github.com/ra9dev/safe-and-sound/pkg/log"
	"github.com/ra9dev/safe-and-sound/pkg/os"
	stdlog "log"

	"github.com/ra9dev/safe-and-sound/internal/police-server/database"
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
)

var version = "unknown"

func main() {
	fmt.Printf("safe-and-sound %s\n", version)

	appCtx, cancelAppCtx := context.WithCancel(context.Background())
	defer cancelAppCtx()
	go os.CatchTermination(cancelAppCtx)

	appConfig := config.Parse(new(configs.DefaultConfig)).(*configs.DefaultConfig)
	log.Setup(appConfig.LogMode)

	ds := database.New(drivers.DataStoreConfig{
		URL:  appConfig.DataStoreURL,
		Name: appConfig.DataStoreName,
		DB:   appConfig.DataStoreDB,
	})
	if err := ds.Connect(); err != nil {
		stdlog.Printf("[ERROR] cannot connect to datastore %s: %v", ds.Name(), err)
		return
	}
	stdlog.Printf("Connected to %s", ds.Name())
	defer ds.Close(appCtx)

	sensorsWatcher := daemons.NewSensorsWatcher(appCtx, []string{"http://localhost:8081"}, ds.Incidents())
	go sensorsWatcher.Run()

	httpSrv := http.NewServer(
		appCtx,
		http.WithVersion(version),
		http.WithCustomAddress(appConfig.ListenAddr),
		http.WithSSL(appConfig.CertFile, appConfig.KeyFile),
		http.WithFiles(appConfig.FilesDir),
		http.WithTestingMode(appConfig.IsTesting),
		http.WithDS(ds),
	)
	if err := httpSrv.Run(); err != nil {
		stdlog.Println(err)
		return
	}

	httpSrv.WaitForGracefulTermination()
	stdlog.Printf("[WARN] process terminated")
}
