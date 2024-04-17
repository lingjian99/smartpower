package builder

import (
	"testing"
	"time"
)

func TestPostgreSqlJoinDoUpdatesetExclude(t *testing.T) {
	uqs := []string{"bid", "date"}
	elems := []string{"bid", "date", "name", "age", "create_time"}
	exclude := []string{"create_time"}
	s := PostgreSqlJoinDoUpdatesetExclude(uqs, elems, exclude)
	t.Log(s)
}

func TestToUpsertSQL(t *testing.T) {
	type DataSt struct {
		Id         int64     `db:"id" xorm:"id"`
		Bid        string    `db:"bid" xorm:"bid"`
		Code       int64     `db:"code" xorm:"code"`
		Flag       int       `db:"flag" xorm:"flag"`
		EventTime  time.Time `db:"event_time" xorm:"event_time" `
		Data       string    `db:"data" xorm:"data"` // json raw string
		CreateTime time.Time `db:"create_time" xorm:"created 'create_time'"`
		UpdateTime time.Time `db:"update_time" xorm:"updated 'update_time'"`
	}
	n := time.Now()
	table := "data_tbl"
	ds := DataSt{
		Id:         0,
		Bid:        "XXXXXXXXXXXXXX",
		Code:       12,
		Flag:       0x01,
		EventTime:  n,
		Data:       "[1,1,1,1]",
		CreateTime: n,
		UpdateTime: n,
	}

	query, values, err := ToUpsertSQL(ds, table, []string{"id"}, []string{"bid", "code", "flag"}, []string{"create_time"})
	t.Log(query, values)
	if err != nil {
		t.Fail()
	}
}

type McAccount struct {
	Id         int64     `db:"id"`
	UserId     int64     `db:"user_id"`
	Banlance   int64     `db:"banlance"` // 余额 分
	Version    int64     `db:"version"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	DelState   int64     `db:"del_state"`
	DeleteTime time.Time `db:"delete_time"`
}

func TestFieldsToSetMap(t *testing.T) {
	data := McAccount{
		Id:         0,
		UserId:     12,
		Banlance:   222,
		Version:    3,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DelState:   0,
		DeleteTime: time.Unix(0, 0),
	}

	res := FieldsToSetMap(data, "id", "create_time")
	t.Logf("%+v", res)

}
