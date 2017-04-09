
package lib

import (
	log "go.uber.org/zap/zapcore"
)


func One() int {
	log.Info("One was called")
	return 1
}
