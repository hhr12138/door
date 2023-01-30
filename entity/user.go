package entity

import "database/sql"

type User struct {
	BasicItem
	Name sql.NullString `db:"name"`
	Password sql.NullString `db:"password"`
	Salt sql.NullString `db:"salt"`
	HeadImage sql.NullString `db:"head_image"`
	Sex sql.NullString `db:"sex"`
	Phone sql.NullString `db:"phone"`
	Birthday sql.NullString `db:"birthday"`
	Signature sql.NullString `db:"signature"`
	RealName sql.NullString `db:"real_name"`
	IdentityCard sql.NullString `db:"identity_card"`
	Country sql.NullString `db:"country"`
	Province sql.NullString `db:"province"`
	City sql.NullString `db:"city"`
	Location sql.NullString `db:"location"`
}
