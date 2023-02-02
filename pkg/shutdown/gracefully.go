package shutdown

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/xrexy/togo/pkg/database"
)

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer func() {
		close(quit)
		database.ClosePostgresDB()
	}()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
