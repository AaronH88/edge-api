// FIXME: golangci-lint
// nolint:revive
package clients

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/edge-api/pkg/routes/common"
	"github.com/redhatinsights/platform-go-middlewares/request_id"
	log "github.com/sirupsen/logrus"
)

// GetOutgoingHeaders returns Red Hat Insights headers from our request to use it
// in other request that will be used when communicating in Red Hat Insights services
func GetOutgoingHeaders(ctx context.Context) map[string]string {
	requestID := request_id.GetReqID(ctx)
	headers := map[string]string{"x-rh-insights-request-id": requestID}
	if config.Get().Auth {
		xrhid, err := common.GetOriginalIdentity(ctx)

		if err != nil {
			log.Error("Error getting x-rh-identity")
		} else {
			headers["x-rh-identity"] = xrhid
		}
	}
	return headers
}

func ConfigureHttpClient(client *http.Client) *http.Client {
	cfg := config.Get()
	timeout, err := time.ParseDuration(fmt.Sprintf("%ds", cfg.HTTPClientTimeout))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("Failed to parse duration")
		return client
	}
	client.Timeout = timeout
	if cfg.TlsCAPath == "" {
		return client
	}
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("Failed to cert from system cert pool")
		rootCAs = x509.NewCertPool()
	}
	certs, err := os.ReadFile(cfg.TlsCAPath)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("Failed to append CA to RootCAs")
		return client
	}
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Warn("adding certificate from PEM failed")
	}
	httpConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: false,
		RootCAs:            rootCAs,
	}
	client.Transport = &http.Transport{TLSClientConfig: httpConfig}
	return client
}
