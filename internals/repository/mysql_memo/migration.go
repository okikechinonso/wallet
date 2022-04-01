package mysqlmemo

import (
	"github.com/okikechinonso/internals/domain/entities"
	"log"
)

func (d *sqLRepository) Migrate() {
	err := d.db.AutoMigrate(&entities.Wallet{})
	if err != nil {
		log.Printf("%s", err)
	}

}
