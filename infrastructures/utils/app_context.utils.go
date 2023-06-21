package utils

import "gorm.io/gorm"

type AppContext interface {
	GetReadDBConnection() *gorm.DB
	GetWriteDBConnection() *gorm.DB
}

type appCtx struct {
	dbRead  *gorm.DB
	dbWrite *gorm.DB
}

func NewAppContext(dbRead *gorm.DB, dbWrite *gorm.DB) *appCtx {
	return &appCtx{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}

func (ctx *appCtx) GetReadDBConnection() *gorm.DB {
	return ctx.dbRead
}

func (ctx *appCtx) GetWriteDBConnection() *gorm.DB {
	return ctx.dbWrite
}
