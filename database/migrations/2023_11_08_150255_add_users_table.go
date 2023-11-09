package migrations

import (
    "database/sql"
    "gohub/app/models"
    "gohub/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type AddUsersTable struct {
        models.BaseModel

        Name     string `gorm:"type:varchar(255);not null;index"`
        Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(20);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&AddUsersTable{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&AddUsersTable{})
    }

    migrate.Add("2023_11_08_150255_add_users_table", up, down)
}