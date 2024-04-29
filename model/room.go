package model

import (
	"github.com/RianIhsan/go-auth-rbac/db"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Name     string `gorm:"not null;unique" json:"name"`
	Location string `gorm:"not null" json:"location"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// add a room
func (room *Room) Save() (*Room, error) {
	err := db.Db.Create(&room).Error
	if err != nil {
		return &Room{}, err
	}
	return room, nil
}

// get all rooms
func GetRooms(Room *[]Room) (err error) {
	err = db.Db.Find(Room).Error
	if err != nil {
		return err
	}
	return nil
}

// get room by id
func GetRoom(Room *Room, id int) (err error) {
	err = db.Db.Where("id = ?", id).First(Room).Error
	if err != nil {
		return err
	}
	return nil
}

// update room
func UpdateRoom(Room *Room) (err error) {
	db.Db.Save(Room)
	return nil
}
