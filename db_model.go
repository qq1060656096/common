package common

import "time"

type DbModel struct {

}

type CreatedAtDbModel struct {
	CreatedAt int64 `gorm:"type:int;"`
}

type UpdatedAtDbModel struct {
	UpdatedAt int64 `gorm:"type:int;"`
}

type DeletedAtDbModel struct {
	DeletedAt int64 `gorm:"type:int;"`
}


func (c *CreatedAtDbModel) BeforeSave() (err error) {
	if c.CreatedAt < 1 {
		c.CreatedAt = time.Now().Unix()
	}
	return
}

func (u *UpdatedAtDbModel) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now().Unix()
	return
}

func (d *DeletedAtDbModel) BeforeDelete() (err error) {
	d.DeletedAt = time.Now().Unix()
	return
}