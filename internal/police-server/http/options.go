package http

import (
	"github.com/ra9dev/safe-and-sound/internal/police-server/database/drivers"
	"net/http"
)

type ServerOption func(srv *Server)

func WithDS(ds drivers.DataStore) ServerOption {
	return func(srv *Server) {
		srv.ds = ds
	}
}

func WithSSL(certFile, keyFile string) ServerOption {
	return func(srv *Server) {
		if certFile != "" {
			srv.CertFile = &certFile
		}

		if keyFile != "" {
			srv.KeyFile = &keyFile
		}
	}
}

func WithVersion(version string) ServerOption {
	return func(srv *Server) {
		srv.version = version
	}
}

func WithCustomAddress(addr string) ServerOption {
	return func(srv *Server) {
		srv.Address = addr
	}
}

func WithFiles(dir http.Dir) ServerOption {
	return func(srv *Server) {
		srv.FilesDir = dir
	}
}

func WithTestingMode(isTesting bool) ServerOption {
	return func(srv *Server) {
		srv.IsTesting = isTesting
	}
}
