package mysqlmemo

import (
	"log"

	"github.com/okikechinonso/wallet/internals/domain/entities"
)

func (d *MySqlDb) Migrate() {

	err := d.MyDB.AutoMigrate(&entities.Wallet{})
	if err != nil {
		log.Printf("%s", err)
	}

}
