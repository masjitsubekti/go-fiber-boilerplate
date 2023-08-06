package models

import (
	"time"

	"github.com/gofrs/uuid"
)

const tableNameAuthor = "author"

type Author struct {
	ID        uuid.UUID  `db:"id" json:"id" gorm:"column:id"`
	Name      string     `db:"name" json:"name" gorm:"column:name"`
	Address   *string    `db:"address" json:"address" gorm:"column:address"`
	CreatedAt time.Time  `db:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at" gorm:"column:deleted_at"`
}

type AuthorRequest struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" validate:"required,lte=255"`
	Address *string   `json:"address"`
}

func (*Author) TableName() string {
	return tableNameAuthor
}

func (i *Author) BindFromRequest(req AuthorRequest) {
	var now = time.Now()
	if req.ID == uuid.Nil {
		newID, _ := uuid.NewV4()
		i.ID = newID
		i.CreatedAt = now
	} else {
		i.ID = req.ID
		i.UpdatedAt = &now
	}

	i.Name = req.Name
	i.Address = req.Address
}

func (i *Author) SoftDelete() {
	var now = time.Now()
	i.DeletedAt = &now
}
