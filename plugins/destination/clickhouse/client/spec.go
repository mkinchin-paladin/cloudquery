package client

import (
	"crypto/x509"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
)

type Spec struct {
	Cluster          string `json:"cluster,omitempty"`
	ConnectionString string `json:"connection_string,omitempty"`
	CACert           string `json:"ca_cert,omitempty"`

	Engine *queries.Engine `json:"engine,omitempty"`

	BatchSize      int           `json:"batch_size,omitempty"`
	BatchSizeBytes int           `json:"batch_size_bytes,omitempty"`
	BatchTimeout   time.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) Options() (*clickhouse.Options, error) {
	options, err := clickhouse.ParseDSN(s.ConnectionString)
	if err != nil {
		return nil, err
	}

	// set database name to "default", if empty
	if len(options.Auth.Database) == 0 {
		options.Auth.Database = "default"
	}

	if tlsConfig := options.TLS; tlsConfig != nil && len(s.CACert) > 0 {
		if tlsConfig.RootCAs == nil {
			tlsConfig.RootCAs, err = x509.SystemCertPool()
			if err != nil {
				return nil, err
			}
		}

		if ok := tlsConfig.RootCAs.AppendCertsFromPEM([]byte(s.CACert)); !ok {
			return nil, fmt.Errorf("failed to append \"ca_cert\" value")
		}
	}

	return options, nil
}

func (s *Spec) SetDefaults() {
	if s.Engine == nil {
		s.Engine = queries.DefaultEngine()
	}

	if s.BatchSize == 0 {
		const defaultBatchSize = 10_000 // 10K
		s.BatchSize = defaultBatchSize
	}

	if s.BatchSizeBytes == 0 {
		const defaultBatchSizeBytes = 5 << 20 // 5 MiB
		s.BatchSizeBytes = defaultBatchSizeBytes
	}

	if s.BatchTimeout == 0 {
		const defaultBatchTimeout = 20 * time.Second // 20s
		s.BatchTimeout = defaultBatchTimeout
	}
}

func (s *Spec) Validate() error {
	return s.Engine.Validate()
}
