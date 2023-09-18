package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG Config
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
)
