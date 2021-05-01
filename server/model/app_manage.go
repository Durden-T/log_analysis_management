package model

import (
	"context"
	"gin-vue-admin/global"
	"github.com/segmentio/kafka-go"
)

type App struct {
	global.GVA_MODEL
	Name             string `json:"name" form:"name" gorm:"comment:app名;unique;'<-:create"`
	KafkaInputTopic  string `json:"kafkaInputTopic" form:"kafkaInputTopic" gorm:"comment:kafka输入topic"`
	KafkaOutputTopic string `json:"kafkaOutputTopic" form:"kafkaOutputTopic" gorm:"comment:kafka输出topic"`
	EnableAlarm      bool   `json:"enableAlarm" form:"enableAlarm" gorm:"comment:启用报警"`

	TemplateCollector *templateCollector `json:"-" gorm:"-"`

	*originalLogAlarmManager `json:"-" gorm:"-"`
	*templateAlarmManager    `json:"-" gorm:"-"`
}

const (
	AppTemplateTableSuffix = "_log_templates"              // 数据库表名后缀
	GroupID                = "log_log_analysis_management" // kafka group id
)

func (a *App) Init() error {
	if a.EnableAlarm {
		if err := a.InitAlarm(); err != nil {
			return err
		}
	}

	if err := a.initTemplateCollector(); err != nil {
		return err
	}
	global.APP_MANAGER.Store(a.Name, a)
	return nil
}

func (a *App) InitAlarm() error {
	if err := a.initTemplateAlarmManager(); err != nil {
		return err
	}
	if err := a.initOriginalLogAlarmManager(); err != nil {
		return err
	}
	return nil
}

func (a *App) DisableAlarm() {
	if a.templateAlarmManager != nil {
		a.templateAlarmManager.Stop()
	}
	if a.originalLogAlarmManager != nil {
		a.originalLogAlarmManager.Stop()
	}
}

func (a *App) initTemplateCollector() error {
	templateTableName := GetTemplateTableName(a.Name)
	if err := global.GVA_DB.Table(templateTableName).AutoMigrate(&LogTemplate{}); err != nil {
		return err
	}

	levelAlarmTableName := GetLevelAlarmTableName(a.Name)
	if err := global.GVA_DB.Table(levelAlarmTableName).AutoMigrate(&LevelAlarmStrategy{}); err != nil {
		return err
	}

	cfg := global.GVA_CONFIG.Kafka

	writer := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Hosts...),
		Topic:        a.KafkaInputTopic,
		RequiredAcks: kafka.RequireOne,
		Compression:  kafka.Lz4,
		BatchTimeout: 0, // 实时写 不缓冲
	}
	// 尝试写一条空信息，出错说明连接可能失败
	if err := writer.WriteMessages(context.Background(), kafka.Message{Value: []byte{}}); err != nil {
		writer.Close()
		return err
	}

	//rand.Seed(time.Now().Unix())
	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:          a.KafkaOutputTopic,
		Brokers:        cfg.Hosts,
		//GroupID:        GroupID+strconv.Itoa(rand.Int()),
		GroupID:        GroupID,
		MinBytes:       cfg.ReadMinBytes,
		MaxBytes:       cfg.ReadMaxBytes,
		CommitInterval: cfg.CommitInterval,
		StartOffset:    kafka.LastOffset,
		//StartOffset:    kafka.FirstOffset,
	})

	a.TemplateCollector = NewTemplateCollector(a.Name, writer, reader, a.originalLogAlarmManager)
	go a.TemplateCollector.Run()
	return nil
}

func (a *App) initTemplateAlarmManager() (err error) {
	a.templateAlarmManager, err = NewTemplateAlarmManager(a.Name)
	if err != nil {
		return
	}
	a.templateAlarmManager.Start()
	return
}

func (a *App) initOriginalLogAlarmManager() (err error) {
	a.originalLogAlarmManager, err = NewOriginalLogAlarmManager(a.Name)
	if err != nil {
		return
	}
	a.originalLogAlarmManager.Start()
	return nil
}

func (a *App) Stop() {
	a.TemplateCollector.Stop()
	a.DisableAlarm()
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetTemplateTableName
//@description: 获取app的日志模版对应的表名
//@param: app *model.App
//@return: tableName string

func GetTemplateTableName(app string) string {
	return app + AppTemplateTableSuffix
}
