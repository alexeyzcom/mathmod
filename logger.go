package mathmod

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"os"
)

func Info(message string, summ int) {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger = logger.With(zap.String("app", "myapp")).With(zap.String("environment", "psm"))

	logger.Info(message, zap.Int("summ", summ))
}
