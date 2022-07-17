package app

import (
	"github.com/bpatyi/gopoc/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	Config *config.Config
	Logger *zap.SugaredLogger
}

func NewApp(configPath, configEnvPrefix *string) (*App, error) {
	cfg, err := config.Load(configPath, configEnvPrefix)
	if err != nil {
		return nil, err
	}

	var logger *zap.SugaredLogger
	var zapCfg zap.Config
	var zapEncoder zapcore.EncoderConfig
	if cfg.Env == config.EnvDebug || cfg.Env == config.EnvStaging {
		zapCfg = zap.NewDevelopmentConfig()
		zapEncoder = zap.NewDevelopmentEncoderConfig()
	} else if cfg.Env == config.EnvProduction {
		zapCfg = zap.NewProductionConfig()
		zapEncoder = zap.NewProductionEncoderConfig()
	}
	zapCfg.Level = getZapLevel(cfg.Logger.Level)
	if cfg.Logger.Encoding != "" {
		zapCfg.Encoding = cfg.Logger.Encoding
	}
	if len(cfg.Logger.OutputPaths) != 0 {
		zapCfg.OutputPaths = cfg.Logger.OutputPaths
	}
	if len(cfg.Logger.ErrorOutputPaths) != 0 {
		zapCfg.ErrorOutputPaths = cfg.Logger.ErrorOutputPaths
	}
	if len(cfg.Logger.InitialFields) != 0 {
		zapCfg.InitialFields = cfg.Logger.InitialFields
	}
	if cfg.Logger.Encoder != nil {
		if cfg.Logger.Encoder.EncodeLevel != "" {
			zapEncoder.EncodeLevel = getZapEncoderLevel(cfg.Logger.Encoder.EncodeLevel)
		}
		if cfg.Logger.Encoder.EncodeTime != "" {
			zapEncoder.EncodeTime = getZapEncoderTime(cfg.Logger.Encoder.EncodeTime)
		}
	}
	zapCfg.EncoderConfig = zapEncoder
	l, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}
	logger = l.Sugar()
	defer logger.Sync() //nolint

	return &App{
		Config: cfg,
		Logger: logger,
	}, nil
}

func getZapLevel(level string) zap.AtomicLevel {
	switch level {
	case zap.DebugLevel.String():
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case zap.ErrorLevel.String():
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case zap.InfoLevel.String():
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case zap.WarnLevel.String():
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case zap.PanicLevel.String():
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case zap.DPanicLevel.String():
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	case zap.FatalLevel.String():
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		panic("unknown log level")
	}
}
func getZapEncoderLevel(encoderLevel string) zapcore.LevelEncoder {
	switch encoderLevel {
	case config.EncoderLevelLowercase:
		return zapcore.LowercaseLevelEncoder
	case config.EncoderLevelLowercaseColor:
		return zapcore.LowercaseColorLevelEncoder
	case config.EncoderLevelCapital:
		return zapcore.CapitalLevelEncoder
	case config.EncoderLevelCapitalColor:
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.CapitalColorLevelEncoder
	}
}
func getZapEncoderTime(encoderTime string) zapcore.TimeEncoder {
	switch encoderTime {
	case config.EncoderTimeEpoch:
		return zapcore.EpochTimeEncoder
	case config.EncoderTimeISO8601:
		return zapcore.ISO8601TimeEncoder
	case config.EncoderTimeEpochNanos:
		return zapcore.EpochNanosTimeEncoder
	case config.EncoderTimeEpochMillis:
		return zapcore.EpochMillisTimeEncoder
	case config.EncoderTimeRFC3339:
		return zapcore.RFC3339TimeEncoder
	case config.EncoderTimeRFC3339Nano:
		return zapcore.RFC3339NanoTimeEncoder
	default:
		return zapcore.ISO8601TimeEncoder
	}
}
