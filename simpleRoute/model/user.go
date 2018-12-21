package model

type User struct {
	Id        int    `db:"id"`
	Nickname  string `db:"nickname"`
	Avatar    string `db:"avatar"`
	Gender    int    `db:"gender"`
	Stage     int    `db:"stage"`
	City      string `db:"city"`
	Province  string `db:"province"`
	Openid    string `db:"openid"`
	Gold      int    `db:"gold"`
	Max_sign  int    `db:"max_sign"`
	Add_app   int    `db:"add_app"`
	Audio     int    `db:"audio"`
	M         string `db:"m"`
	Eid       int    `db:"eid"`
	CreatedAt string `db:"createdAt"`
	UpdatedAt string `db:"updatedAt"`
}