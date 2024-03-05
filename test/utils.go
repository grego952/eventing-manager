package test

import (
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	kcorev1 "k8s.io/api/core/v1"
	kmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kyma-project/eventing-manager/pkg/logger"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func NewEventingLogger() (*logger.Logger, error) {
	return logger.New("json", "info")
}

func NewLogger() (*zap.Logger, error) {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.Encoding = "json"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000")

	return loggerConfig.Build()
}

func NewSugaredLogger() (*zap.SugaredLogger, error) {
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

func NewNamespace(name string) *kcorev1.Namespace {
	namespace := kcorev1.Namespace{
		TypeMeta: kmetav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: kmetav1.ObjectMeta{
			Name: name,
		},
	}
	return &namespace
}

// GetRandK8sName returns a valid name for K8s objects.
func GetRandK8sName(length int) string {
	return fmt.Sprintf("name-%s", GetRandString(length))
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec,gochecknoglobals // used in tests

// GetRandString returns a random string of the given length.
func GetRandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// findEnvVar returns the env variable which has `name == envVar.Name`,
// or `nil` if there is no such env variable.
func FindEnvVar(envVars []kcorev1.EnvVar, name string) *kcorev1.EnvVar {
	for _, envvar := range envVars {
		if name == envvar.Name {
			return &envvar
		}
	}
	return nil
}
