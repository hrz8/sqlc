// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import ()

type JoinTable struct {
	ID             int32
	PrimaryTableID int32
	OtherTableID   int32
	IsActive       bool
}

type PrimaryTable struct {
	ID     int32
	UserID int32
}
