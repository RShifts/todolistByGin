package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Init( /*DB *gorm.DB*/ ) {
	var err error
	DB, err = gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/dev_2024?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
}

func UUID() string {
	random, _ := uuid.NewRandom()
	return random.String()
}

type Todo struct {
	Id             string    `gorm:"column:id;primaryKey;comment:主键" ;json:"id"`
	ListId         string    `gorm:"column:list_id;comment:待办列表id" ;json:"listId"`
	Content        string    `gorm:"column:content;comment:内容" ;json:"content"`
	Done           int       `gorm:"column:done;comment:是否已完成(1 完成 0/ NULL 未完成)" ;json:"done"`
	Sort           int       `gorm:"column:sort;comment:排序" ;json:"sort"`
	Deleted        int       `gorm:"column:deleted;comment:已删除" ;json:"deleted"`
	CreateTime     time.Time `gorm:"column:create_time;comment:创建时间" ;json:"createTime"`
	LastUpdateTime time.Time `gorm:"column:last_update_time;comment:最后更新时间" ;json:"lastUpdateTime"`
}

type List struct {
	Id             string    `gorm:"column:id;primaryKey;comment:主键"`
	Title          string    `gorm:"column:title;comment:待办列表标题"`
	Deleted        int       `gorm:"column:deleted;comment:已删除"`
	CreateTime     time.Time `gorm:"column:create_time;comment:创建时间"`
	LastUpdateTime time.Time `gorm:"column:last_update_time;comment:最后更新时间"`
}
