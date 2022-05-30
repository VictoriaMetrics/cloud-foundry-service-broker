package config

import (
	"flag"
	"github.com/VictoriaMetrics/cloud-foundry-service-broker/lib/logger"
)

// PrintFlags prints all flags to the log.
func PrintFlags() {
	configLogger := logger.DefaultLogger
	configLogger.Info("dumping runtime flags")

	flag.VisitAll(func(f *flag.Flag) {
		configLogger.Infow("parsed flag", "name", f.Name, "value", f.Value.String())
	})
}
