package sql

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"go.elastic.co/apm/module/apmgorm"

	"github.com/jinzhu/gorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/mysql"
)

//
//global pool for mysql connection will be registered while we boot up our api
//
var ObjDb *MySqlPool

//
// default time zone for mysql
//
const DEFAULTTIMEZONE = "Asia/Jakarta"

// A pool maintains a set of connections.
// Bydefault no connection is created. The connection is created only when query is fired.
type MySqlPool struct {
	Pool *gorm.DB
}

//GetConnection returns a fresh *MySqlConnection object which can be further used to perform queries.
func (this *MySqlPool) GetConnection() *MySqlConnection {
	connection := &MySqlConnection{Conn: this.Pool}
	return connection
}

//Close the DBpool
func (this *MySqlPool) Close() error {
	return this.Pool.Close()
}

//On a broader level this can be seen as a Mysql connection.
type MySqlConnection struct {
	Conn *gorm.DB
}

//A dummy function which pretends to close the MySqlConnection
//but actually MySqlConnection is a virtual entity that does not make any connection, thus does not needs to be closed.
// Its actually the stmt and tx that needs to be closed. Closing of stmt and tx is internally handled by this wrapper.
// SO, its safe if the user does not call this close method.But for clarity purpose user should call this method.
func (this *MySqlConnection) Close() error {
	return nil
}

func Initiate(this MysqlConfig) (*MySqlPool, error) {

	conn := FormatDSN(this)
	db, err := apmgorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(this.MaxIdleConnections)
	db.DB().SetMaxOpenConns(this.MaxOpenConnections)
	//Check if we have encountered any errors
	if err != nil {
		return nil, err
	}
	//On DB error
	if err = db.Error; err != nil {
		return nil, err
	}

	ObjDb = &MySqlPool{Pool: db}
	return ObjDb, nil
}

//converts the configuration to the format understood by go sql driver.
func FormatDSN(this MysqlConfig) string {

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		this.User,
		this.Password,
		this.Host,
		this.Port,
		this.DBName,
		url.QueryEscape(this.DefaultTimeZone),
	)
}
