// Copyright (C) 2017 ScyllaDB

package dynamoping

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
	"github.com/scylladb/scylla-manager/pkg/ping"
	"github.com/scylladb/scylla-manager/pkg/util/timeutc"
)

// Config specifies the ping configuration, note that timeout is mandatory and
// zero timeout will result in errors.
type Config struct {
	Addr                   string
	RequiresAuthentication bool
	Username               string
	Password               string
	Timeout                time.Duration
	TLSConfig              *tls.Config
}

// SimplePing sends GET request to alternator port and expects 200 response code.
func SimplePing(ctx context.Context, config Config) (rtt time.Duration, err error) {
	t := timeutc.Now()
	defer func() {
		rtt = timeutc.Since(t)
		if rtt >= config.Timeout {
			err = ping.ErrTimeout
		}
	}()

	c := httpClient(config)
	ctx, cancel := context.WithDeadline(ctx, t.Add(config.Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.Addr, http.NoBody)
	if err != nil {
		return 0, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, errors.Errorf("unexpected response code %d", resp.StatusCode)
	}

	return 0, nil
}

var unauthorisedMessage = []string{
	"User not found",
	"The security token included in the request is invalid.",
}

// QueryPing checks if host is available, it returns RTT and error. Special errors
// are ErrTimeout and ErrUnauthorised. Ping is based on executing
// a real query.
func QueryPing(ctx context.Context, config Config) (rtt time.Duration, err error) {
	t := timeutc.Now()
	defer func() {
		rtt = timeutc.Since(t)
		if rtt >= config.Timeout {
			err = ping.ErrTimeout
		} else if err != nil {
			for _, m := range unauthorisedMessage {
				if strings.Contains(err.Error(), m) {
					err = ping.ErrUnauthorised
					break
				}
			}
		}
	}()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint:   aws.String(config.Addr),
			Region:     aws.String("scylla"),
			HTTPClient: httpClient(config),
		},
	}))

	if config.RequiresAuthentication && config.Username != "" && config.Password != "" {
		sess.Config.Credentials = credentials.NewStaticCredentials(config.Username, config.Password, "")
	} else {
		sess.Config.Credentials = credentials.AnonymousCredentials
	}

	svc := dynamodb.New(sess)

	ctx, cancel := context.WithDeadline(ctx, t.Add(config.Timeout))
	defer cancel()

	const tableName = ".scylla.alternator.system_schema.tables"
	_, err = svc.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		Limit:     aws.Int64(1),
	})
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func httpClient(config Config) *http.Client {
	// Disable keepalive to kill idle connections immediately.
	t := &http.Transport{
		DisableKeepAlives: true,
	}
	if config.TLSConfig != nil {
		t.TLSClientConfig = config.TLSConfig.Clone()
	}

	return &http.Client{
		Timeout:   config.Timeout,
		Transport: t,
	}
}
