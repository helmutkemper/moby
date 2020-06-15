// +build !windows

package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"net"
	"net/http"
	"path/filepath"
	"strings"

	metrics "github.com/docker/go-metrics"
	"github.com/helmutkemper/moby/pkg/plugingetter"
	"github.com/helmutkemper/moby/pkg/plugins"
	"github.com/helmutkemper/moby/plugin"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func (daemon *Daemon) listenMetricsSock() (string, error) {
	path := filepath.Join(daemon.configStore.ExecRoot, "metrics.sock")
	unix.Unlink(path)
	l, err := net.Listen("unix", path)
	if err != nil {
		return "", errors.Wrap(err, "error setting up metrics plugin listener")
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", metrics.Handler())
	go func() {
		logrus.Debugf("metrics API listening on %s", l.Addr())
		if err := http.Serve(l, mux); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			logrus.WithError(err).Error("error serving metrics API")
		}
	}()
	daemon.metricsPluginListener = l
	return path, nil
}

func registerMetricsPluginCallback(store *plugin.Store, sockPath string) {
	store.RegisterRuntimeOpt(metricsPluginType, func(s *specs.Spec) {
		f := plugin.WithSpecMounts([]specs.Mount{
			{Type: "bind", Source: sockPath, Destination: "/run/docker/metrics.sock", Options: []string{"bind", "ro"}},
		})
		f(s)
	})
	store.Handle(metricsPluginType, func(name string, client *plugins.Client) {
		// Use lookup since nothing in the system can really reference it, no need
		// to protect against removal
		p, err := store.Get(name, metricsPluginType, plugingetter.Lookup)
		if err != nil {
			return
		}

		adapter, err := makePluginAdapter(p)
		if err != nil {
			logrus.WithError(err).WithField("plugin", p.Name()).Error("Error creating plugin adapter")
		}
		if err := adapter.StartMetrics(); err != nil {
			logrus.WithError(err).WithField("plugin", p.Name()).Error("Error starting metrics collector plugin")
		}
	})
}
