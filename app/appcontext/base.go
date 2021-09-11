package appcontext

import "github.com/murbagus/gin-simple-mvc/config"

type ContextInterface interface {
	DatabaseContextInterface
}

// Context aplikasi berisi data data yang digunakan oleh aplikasi
// sepertid data koneksi-koneksi database
type Context struct {
	DatabaseContext
}

// Membuat context baru
func NewContext() *Context {
	context := &Context{}

	context.SetDatabaseConnection(config.GetDatabaseConnection())

	return context
}
