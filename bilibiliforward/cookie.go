package bilibiliforward

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/FloatTech/floatbox/file"
	sql "github.com/FloatTech/sqlite"
)

// Storage cookie库系统
type Storage struct {
	sync.RWMutex
	db *sql.Sqlite
}

// Record 记录
type Record struct {
	UID    int64
	Cookie string
}

var (
	sdb = &Storage{
		db: &sql.Sqlite{
			DBPath: "data/bilibiliforward/record.db",
		},
	}
)

func init() {
	if file.IsNotExist("data/bilibiliforward") {
		err := os.MkdirAll("data/bilibiliforward", 0755)
		if err != nil {
			panic(err)
		}
	}
	err := sdb.db.Open(time.Hour * 24)
	if err != nil {
		panic(err)
	}
	err = sdb.db.Create("storage", &Record{})
	if err != nil {
		panic(err)
	}
}

// GetCookieOf 获取某人cookie
func GetCookieOf(uid int64) (cookie string) {
	return sdb.getCookieOf(uid).Cookie
}

// InsertCookieOf 更新cookie(uid QQ号, cookie 用户提供的cookie)
func InsertCookieOf(uid int64, cookie string) error {
	return sdb.updateCookieOf(uid, cookie)
}

// 获取用户cookie
func (sql *Storage) getCookieOf(uid int64) (record Record) {
	sql.RLock()
	defer sql.RUnlock()
	uidstr := strconv.FormatInt(uid, 10)
	_ = sql.db.Find("storage", &record, "where uid is "+uidstr)
	return
}

// 更新cookie
func (sql *Storage) updateCookieOf(uid int64, cookie string) (err error) {
	sql.Lock()
	defer sql.Unlock()
	return sql.db.Insert("storage", &Record{
		UID:    uid,
		Cookie: cookie,
	})
}
