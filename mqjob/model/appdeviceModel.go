package model

import (
	"smartpower/common/sql/builder"
	"time"

	"xorm.io/xorm"
)

const (
	UserApp = "user"
	OpsApp  = "ops"
)

type (
	AppDeviceTokenModel interface {
		Transaction(f func(*xorm.Session) (interface{}, error)) error
		DelDeviceToken(session *xorm.Session, token string) error
		Upsert(session *xorm.Session, data Upsertable) (int64, error)
		GetDeviceToken(app string, userId int64) (*AppDeviceToken, error)
	}

	AppDeviceToken struct {
		Id         int64     `db:"id"`
		App        string    `db:"app"`
		UserId     int64     `db:"user_id"`
		Token      string    `db:"token"`
		Platform   string    `db:"platform"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}

	defalutAppDeviceTokenModel struct {
		db *xorm.Engine
	}
	
	Upsertable interface {
		UpsertToSql() (query string, values []interface{}, err error)
	}
)

func (d *defalutAppDeviceTokenModel) Transaction(f func(*xorm.Session) (interface{}, error)) error {
	_, err := d.db.Transaction(f)
	return err
}

func (d *defalutAppDeviceTokenModel) DelDeviceToken(session *xorm.Session, token string) error {
	if session != nil {
		_, err := session.Where("token = ?", token).Delete(&AppDeviceToken{})
		return err
	}
	_, err := d.db.Where("token = ?", token).Delete(&AppDeviceToken{})
	return err
}

func (d *defalutAppDeviceTokenModel) GetDeviceToken(app string, userId int64) (*AppDeviceToken, error) {
	var data AppDeviceToken
	ok, err := d.db.Where("app = ?", app).And("user_id = ?", userId).Get(&data)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNotFound
	}
	return &data, nil
}

func (d *defalutAppDeviceTokenModel) Upsert(session *xorm.Session, data Upsertable) (int64, error) {
	query, values, err := data.UpsertToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	if session != nil {
		_, err = session.SQL(query, values...).Get(&id)
		return id, err
	}
	_, err = d.db.SQL(query, values...).Get(&id)
	return id, err
}

func NewAppDeviceTokenModel(db *xorm.Engine) AppDeviceTokenModel {
	return &defalutAppDeviceTokenModel{db: db}
}

func (*AppDeviceToken) TableName() string {
	return "app_device_token"
}

func (data *AppDeviceToken) UpsertToSql() (query string, values []interface{}, err error) {
	return builder.ToUpsertSQL(data, data.TableName(), []string{"id"}, []string{"app", "user_id"}, []string{"create_time"})
}
