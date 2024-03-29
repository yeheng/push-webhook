package app

import (
	"fmt"
	"strings"

	"github.com/YeHeng/push-webhook/common/model"
	"github.com/gin-gonic/gin"
)

type TransformerStrategy interface {
	Transform(c *gin.Context) (*model.PushMessage, error)
}

var transformers = map[string]TransformerStrategy{}

func GetTransformer(channel string) (TransformerStrategy, error) {
	s, ok := transformers[strings.ToUpper(channel)]
	if !ok || s == nil {
		return nil, fmt.Errorf("找不到推送渠道: %s", channel)
	}
	return s, nil
}

func RegisterTransformer(channel string, service TransformerStrategy) {
	transformers[strings.ToUpper(channel)] = service
}
