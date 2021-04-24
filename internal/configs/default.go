package configs

import (
	"github.com/hashicorp/logutils"
	"net/http"
)

type DefaultConfig struct {
	DataStoreName string `short:"n" long:"ds" env:"DATASTORE" description:"DataStore name (format: dgraph/null)" required:"false" default:"mongo"`
	DataStoreDB   string `short:"d" long:"ds-db" env:"DATASTORE_DB" description:"DataStore database name (format: inventory)" required:"false" default:"safe-and-sound"`
	DataStoreURL  string `short:"u" long:"ds-url" env:"DATASTORE_URL" description:"DataStore URL (format: mongodb://localhost:27017)" required:"false" default:"mongodb://localhost:27017"`

	ListenAddr string   `short:"l" long:"listen" env:"LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`
	FilesDir   http.Dir `long:"files-directory" env:"FILES_DIR" description:"Directory where all static files are located" required:"false" default:"./api"`
	CertFile   string   `short:"c" long:"cert" env:"CERT_FILE" description:"Location of the SSL/TLS cert file" required:"false" default:""`
	KeyFile    string   `short:"k" long:"key" env:"KEY_FILE" description:"Location of the SSL/TLS key file" required:"false" default:""`

	LogMode   logutils.LogLevel `long:"log-mode" env:"LOG_MODE" description:"log mode" default:"WARN"`
	IsTesting bool              `long:"testing" env:"APP_TESTING" description:"testing mode"`
}

type SensorConfig struct {
	ID string `long:"sensor-id" env:"SENSOR_ID" description:"sensor id" default:"myID"`

	ListenAddr string `short:"l" long:"listen" env:"LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":8080"`

	LogMode logutils.LogLevel `long:"log-mode" env:"LOG_MODE" description:"log mode" default:"WARN"`
}
