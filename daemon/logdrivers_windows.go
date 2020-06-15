package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/helmutkemper/moby/daemon/logger/awslogs"
	_ "github.com/helmutkemper/moby/daemon/logger/etwlogs"
	_ "github.com/helmutkemper/moby/daemon/logger/fluentd"
	_ "github.com/helmutkemper/moby/daemon/logger/gcplogs"
	_ "github.com/helmutkemper/moby/daemon/logger/gelf"
	_ "github.com/helmutkemper/moby/daemon/logger/jsonfilelog"
	_ "github.com/helmutkemper/moby/daemon/logger/logentries"
	_ "github.com/helmutkemper/moby/daemon/logger/loggerutils/cache"
	_ "github.com/helmutkemper/moby/daemon/logger/splunk"
	_ "github.com/helmutkemper/moby/daemon/logger/syslog"
)
