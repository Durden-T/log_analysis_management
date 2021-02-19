package model

import (
	"context"
	"gin-vue-admin/global"
	"github.com/segmentio/kafka-go"
)

type App struct {
	global.GVA_MODEL
	Name             string     `json:"name" form:"name" gorm:"comment:app名;unique;'<-:create"`
	KafkaInputTopic  string     `json:"kafkaInputTopic" form:"kafkaInputTopic" gorm:"comment:kafka输入topic"`
	KafkaOutputTopic string     `json:"kafkaOutputTopic" form:"kafkaOutputTopic" gorm:"comment:kafka输出topic"`
	EnableAlarm      bool       `json:"enableAlarm" form:"enableAlarm" gorm:"comment:启用报警"`
	LogParser        *logParser `json:"-" gorm:"-"`

	*OriginalLogAlarmManager
	*TemplateAlarmManager `json:"-" gorm:"-"`
}

const (
	AppTemplateTableSuffix = "_log_templates"
	GroupID                = "log_log_analysis_management"
)

func (a *App) Init() error {
	if err := a.initLogParser(); err != nil {
		return err
	}
	if a.EnableAlarm {
		if err := a.InitAlarm(); err != nil {
			return err
		}
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
	if a.TemplateAlarmManager != nil {
		a.TemplateAlarmManager.Stop()
	}
	if a.OriginalLogAlarmManager != nil {
		a.OriginalLogAlarmManager.Stop()
	}
}

func (a *App) initLogParser() error {
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

	if err := writer.WriteMessages(context.Background(), kafka.Message{Value: []byte{}}); err != nil {
		writer.Close()
		return err
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:          a.KafkaOutputTopic,
		Brokers:        cfg.Hosts,
		GroupID:        GroupID,
		MinBytes:       cfg.ReadMinBytes,
		MaxBytes:       cfg.ReadMaxBytes,
		CommitInterval: cfg.CommitInterval,
		StartOffset:    kafka.LastOffset,
	})

	a.LogParser = NewLogParser(a.Name, writer, reader)
	go a.LogParser.Run()
	return nil
}

func (a *App) initTemplateAlarmManager() (err error) {
	a.TemplateAlarmManager, err = NewTemplateAlarmManager(a.Name)
	if err != nil {
		return
	}
	return
}

func (a *App) initOriginalLogAlarmManager() (err error) {
	cfg := global.GVA_CONFIG.Kafka

	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:          a.KafkaOutputTopic,
		Brokers:        cfg.Hosts,
		GroupID:        GroupID,
		MinBytes:       cfg.ReadMinBytes,
		MaxBytes:       cfg.ReadMaxBytes,
		CommitInterval: cfg.CommitInterval,
		StartOffset:    kafka.LastOffset,
	})

	a.OriginalLogAlarmManager, err = NewOriginalLogAlarmManager(a.Name, reader)
	if err != nil {
		return
	}
	go a.OriginalLogAlarmManager.Run()
	return nil
}

func (a *App) Stop() {
	a.LogParser.Stop()
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
