package sign

import (
	"github.com/hit9/log"
	env "github.com/yuanblq/go-sign/pkg/env"
)

var logger = log.Get("jd")

type JdKey struct {
	PtPin, PtKey string
}

func JD(k JdKey) Result {
	ptPin, ptKey := k.PtPin, k.PtKey

	if k.PtPin == "" {
		ptPin = env.GetEnv("JD_PT_PIN")
		ptKey = env.GetEnv("JD_PT_KEY")
	}

	logger.Info("JD_PT_PIN: %s, JD_PT_KEY: %s", ptPin, ptKey)

	return Result{}
}
