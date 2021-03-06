package serve

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
	"github.com/justinas/alice"

	"github.com/flimzy/kivik/serve/couchserver"
	"github.com/flimzy/kivik/serve/logger"
)

func (s *Service) setupRoutes() (http.Handler, error) {
	h := couchserver.Handler{
		Client:        s.Client,
		Vendor:        s.VendorName,
		VendorVersion: s.VendorVersion,
		Favicon:       s.Favicon,
		SessionKey:    SessionKey,
	}

	rlog := s.RequestLogger
	if rlog == nil {
		rlog = logger.DefaultLogger
	}

	return alice.New(
		setContext(s),
		setSession(),
		loggerMiddleware(rlog),
		gzipHandler(s),
		authHandler,
	).Then(h.Main()), nil
}

func gzipHandler(s *Service) func(http.Handler) http.Handler {
	level := s.Conf().GetInt("httpd.compression_level")
	if level == 0 {
		level = 8
	}
	gzipHandler, err := gziphandler.NewGzipLevelHandler(int(level))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid httpd.compression_level %d\n", level)
		return func(h http.Handler) http.Handler {
			return h
		}
	}
	fmt.Fprintf(os.Stderr, "Enabling gzip compression, level %d\n", level)
	return gzipHandler
}
