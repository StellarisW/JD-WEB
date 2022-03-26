package g

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"main/app/internal/model/config"
	"sync"
	"time"
)

type Model struct {
	Id         int    `json:"id" form:"id" db:"id"`                            //`gorm:"primarykey"` // 主键ID
	UpdateTime string `json:"update_time" form:"update_time" db:"update_time"` // 更新时间
	CreateTime string `json:"create_time" form:"create_time" db:"create_time"` // 创建时间
}

var (
	DB *gorm.DB
	//DBList map[string]*gorm.DB
	Redis  *redis.Client
	Config config.Get
	VP     *viper.Viper
	Logger *zap.SugaredLogger

	//GVA_Timer           timer.Timer = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{} //做并发控制防止缓存击穿

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

func (m Model) GetUpdatedTime() time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", m.UpdateTime, time.Local)
	return t
}

func (m Model) GetCreatedTime() time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", m.CreateTime, time.Local)
	return t
}
