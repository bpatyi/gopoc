package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap/zapcore"
)

const (
	ConsoleEncoding = "console"
	JsonEncoding    = "json"
)

var (
	ListOfLevel = []interface{}{
		zapcore.DebugLevel.String(),
		zapcore.InfoLevel.String(),
		zapcore.WarnLevel.String(),
		zapcore.ErrorLevel.String(),
		zapcore.DPanicLevel.String(),
		zapcore.PanicLevel.String(),
		zapcore.FatalLevel.String(),
	}
	ListOfEncoding = []interface{}{
		ConsoleEncoding,
		JsonEncoding,
	}
)

type Logger struct {
	Level            string                 `json:"level"`
	Encoding         string                 `json:"encoding,omitempty"`
	InitialFields    map[string]interface{} `json:"initial_fields,omitempty"`
	OutputPaths      []string               `json:"output_paths,omitempty"`
	ErrorOutputPaths []string               `json:"error_output_paths,omitempty"`
	Encoder          *LoggerEncoder         `json:"encoder,omitempty"`
	RequestLogOn     bool                   `json:"request_log_on"`
}

func (l Logger) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Level, validation.Required, validation.In(ListOfLevel...)),
		validation.Field(&l.Encoding, validation.In(ListOfEncoding...)),
		validation.Field(&l.OutputPaths, validation.NilOrNotEmpty),
		validation.Field(&l.ErrorOutputPaths, validation.NilOrNotEmpty),
		validation.Field(&l.Encoder),
	)
}

const (
	EncoderLevelLowercase      = "lowercase"
	EncoderLevelLowercaseColor = "lowercase_color"
	EncoderLevelCapital        = "capital"
	EncoderLevelCapitalColor   = "capital_color"
	EncoderTimeEpoch           = "epoch"
	EncoderTimeEpochMillis     = "epoch_millis"
	EncoderTimeEpochNanos      = "epoch_nanos"
	EncoderTimeISO8601         = "iso8601"
	EncoderTimeRFC3339         = "rfc3339"
	EncoderTimeRFC3339Nano     = "rfc3339_nano"
)

var (
	ListOfEncoderLevel = []interface{}{
		EncoderLevelLowercase,
		EncoderLevelLowercaseColor,
		EncoderLevelCapital,
		EncoderLevelCapitalColor,
	}
	ListOfEncoderTime = []interface{}{
		EncoderTimeEpoch,
		EncoderTimeEpochMillis,
		EncoderTimeEpochNanos,
		EncoderTimeISO8601,
		EncoderTimeRFC3339,
		EncoderTimeRFC3339Nano,
	}
)

type LoggerEncoder struct {
	EncodeLevel string `json:"encode_level"`
	EncodeTime  string `json:"encode_time"`
}

func (l LoggerEncoder) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.EncodeLevel, validation.When(l.EncodeLevel != "", validation.In(ListOfEncoderLevel...))),
		validation.Field(&l.EncodeTime, validation.When(l.EncodeTime != "", validation.In(ListOfEncoderTime...))),
	)
}
