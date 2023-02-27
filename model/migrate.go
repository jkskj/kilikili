package model

// 执行数据迁移
func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Video{}).AutoMigrate(&BulletComment{}).
		AutoMigrate(&Comment{}).AutoMigrate(&Reply{}).AutoMigrate(&Interaction{})
	DB.Model(&Video{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&BulletComment{}).AddForeignKey("vid", "Video(id)", "CASCADE", "CASCADE").
		AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Comment{}).AddForeignKey("vid", "Video(id)", "CASCADE", "CASCADE").
		AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Reply{}).AddForeignKey("vid", "Video(id)", "CASCADE", "CASCADE").
		AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Interaction{}).AddForeignKey("vid", "Video(id)", "CASCADE", "CASCADE").
		AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
