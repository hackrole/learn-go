package main

import (
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivateAt   sql.NullTime
	CreatedAt    time.Time
	UpdateAt     time.Time
}

// no-conventions
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:autoCreateTime`
	UpdatedAt time.Time      `gorm:autoUpdateTime:nano`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// embedded
type U struct {
	gorm.Model
	Name string
}

// equals above
type U2 struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

// embedded
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

// equals above
type Blog2 struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}

// embeddedprefix
type B2 struct {
	ID      int
	Author  Author `gorm:'embedded;embeddedPrefix:author_`
	Upvotes int32
}

// equals
type B3 struct {
	ID          int64
	AuthorName  string
	AuthorEmail string
	Upvotes     int32
}
