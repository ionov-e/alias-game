package setup

import (
	"fmt"
	"log"
	"os" //nolint:nolintlint,goimports
	"time"
)

func Logging() error {
	t := time.Now()
	logDir := "log/" + t.Format("2006/01")
	if err := os.MkdirAll(logDir, 0777); err != nil {
		return fmt.Errorf("creating log dir failed: %w", err)
	}
	logPath := logDir + "/" + t.Format("02") + ".log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("opening log file failed: %w", err)
	}
	log.SetOutput(file) // TODO slog or implement similar
	return nil
}
