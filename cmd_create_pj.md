// step 1
1 go mod init  name project
2 go get github.com/gofiber/fiber/v2
3 go get gorm.io/gorm
4 you will to go get driver of only DB as postgrade,mysql ...

// step 2
 
 1 set env

-  [code for gen struct gorm]

// g := gen.NewGenerator(gen.Config{
	// 	OutPath: "./query",
	// 	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	// })

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	//g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	// g.ApplyBasic(
	// 	// Generate struct `User` based on table `users`
	// 	g.GenerateModel("users"),

	// 	// Generate struct `Employee` based on table `users`
	// 	g.GenerateModelAs("users", "Employee"),

	// 	// Generate struct `User` based on table `users` and generating options
	// 	g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	// )
	// g.ApplyBasic(
	// 	// Generate structs from all tables of current database
	// 	g.GenerateAllTable()...,
	// )
	// Generate the code
	// g.Execute()