package models

// AutoMigrate auto migrate
func AutoMigrate() {
	UserAutoMigrate()
	PostAutoMigrate()
}
