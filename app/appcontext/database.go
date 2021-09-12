package appcontext

import "gorm.io/gorm"

type DatabaseContextInterface interface {
	// Untuk melakukan koneksi ke database
	ConnectDatabase(connectionName string)
	SetDatabaseConnection(connections map[string]*gorm.DB)
}

type DatabaseContext struct {
	Database map[string]*gorm.DB
}

func (dbc DatabaseContext) ConnectDatabase(connectionName string) *gorm.DB {
	if val, ok := dbc.Database[connectionName]; ok {
		return val
	} else {
		panic("koneksi database tidak ditemukan")
	}
}

func (dbc *DatabaseContext) SetDatabaseConnection(connections map[string]*gorm.DB) {
	dbc.Database = connections
}
