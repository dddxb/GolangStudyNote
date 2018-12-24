package backup

/*import (
	"database/sql"
	"fmt"
	"GolangStudy/simpleRoute/model"
)

func Connect(sqlType string, host string, port int, user string, password string, dbName string, charset string){
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", user, password, host, port, dbName, charset))
	s.CheckError(err)
	db.Ping()
	s.CheckError(err)
	s.Db = db
}

func Query(sql string, user model.User) {
	rows, err := s.Db.Query(sql)
	s.CheckError(err)
	defer rows.Close()
	var data []model.User
	for rows.Next() {
		var temp model.User
		rows.Scan(&temp.Id, &temp.Nickname, &temp.Avatar, &temp.Gold)
		data = append(data, temp)
	}
	fmt.Println(data)
}

func Insert(sql string, param ...interface{}) int64 {
	prepare, err := s.Db.Prepare(sql)
	s.CheckError(err)
	res, err := prepare.Exec(param...)
	s.CheckError(err)
	id, err := res.LastInsertId()
	s.CheckError(err)
	return id
}

func Update(sql string, param ...interface{}) int64 {
	prepare, err := s.Db.Prepare(sql)
	s.CheckError(err)
	res, err := prepare.Exec(param...)
	s.CheckError(err)
	affect, err := res.RowsAffected()
	s.CheckError(err)
	return affect
}

func Delete(sql string, param ...interface{}) int64 {
	prepare, err := s.Db.Prepare(sql)
	s.CheckError(err)
	res, err := prepare.Exec(param...)
	s.CheckError(err)
	affect, err := res.RowsAffected()
	s.CheckError(err)
	return affect
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
*/