package dao

import (
	"citicup-admin/internal/config"
	"context"
)

// Dao struct
type Dao struct {
	//config
	c *config.Config
	//database
	db *config.DBLink
}

func New(c *config.Config) (dao *Dao) {
	dao = &Dao{
		c:  c,
		db: config.SetUpDBLink(c.DataBase),
	}
	return
}

func (d *Dao) Ping(ctx context.Context) (err error) {
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}
