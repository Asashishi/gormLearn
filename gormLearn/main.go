package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gormLearn/entity"
	// "gorm.io/driver/sqlite"
	"log"
	"os"
)

/*
gorm 用于与SQL进行结构体与表结构映射
*/

//var db *gorm.DB

//func init() {
//
//	// 配置数据源
//	dsn := "root:Asashishi107QS.@tcp(8.218.247.195:3306)/gorm_learn?charset=utf8&parseTime=True&loc=Local"
//
//	// gorm处理MySQL链接和操作
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// 设置数据库连接池
//	sqlDB, _ := db.DB()
//
//	// 最大连接数
//	sqlDB.SetMaxIdleConns(25)
//	// 最大空闲连接数
//	sqlDB.SetMaxOpenConns(10)
//
//	//SetMaxOpenConns(n int)：
//	//功能：设置最大打开的连接数，包括正在使用的连接和连接池中的连接。
//	//默认值：0，表示不限制。
//	//示例：db.SetMaxOpenConns(100) 设置最大打开连接数为100。
//	//SetMaxIdleConns(n int)：
//	//功能：设置连接池中的最大空闲连接数。
//	//默认值：0，表示不保持空闲连接。
//	//示例：db.SetMaxIdleConns(10) 设置最大空闲连接数为10。
//	//SetConnMaxLifetime(d time.Duration)：
//	//功能：设置连接的最大生命周期，超过这个时间的连接将被关闭。
//	//示例：db.SetConnMaxLifetime(time.Hour) 设置连接的最大生命周期为1小时。
//	//SetConnMaxIdleTime(d time.Duration)：
//	//功能：设置连接的最大空闲时间，超过这个时间的空闲连接将被关闭。
//	//示例：db.SetConnMaxIdleTime(30 * time.Minute) 设置连接的最大空闲时间为30分钟。
//	//SetConnMaxIdleTime(d time.Duration)：
//	//功能：设置连接的最大空闲时间，超过这个时间的空闲连接将被关闭。
//	//示例：db.SetConnMaxIdleTime(30 * time.Minute) 设置连接的最大空闲时间为30分钟。
//	//Ping()：
//	//功能：验证数据库连接是否有效。通常在初始化时调用以确保连接池可用。
//	//示例：err := db.Ping() 如果连接失败，将返回错误。
//	//Conn()：
//	//功能：从连接池中获取一个连接。
//	//示例：conn, err := db.Conn(context.Background()) 获取一个连接。
//	//Close()：
//	//功能：关闭数据库连接池，释放所有资源。
//	//示例：db.Close() 关闭连接池
//
//}

//func GetDB() *gorm.DB {
//	return db
//}

// CRUD
// insert
func addRecord(db *gorm.DB) {

	// 结构体对象和表记录做映射

	// 添加老师对象
	teacher := entity.Teacher{
		Name: "Asashishi",
		Tno:  1072903224,
		Pwd:  "Asashishi107QS.",
	}
	fmt.Println(teacher)
	// 按照实例化对象的指针生成SQL语句,gorm会进行格式化并对数据进行回写 插入记录
	db.Debug().Create(&teacher)
	fmt.Println(teacher)

	// 班级
	class01 := entity.Class{
		Name:      "软件一班",
		Num:       78,
		TeacherID: 1,
	}
	class02 := entity.Class{
		Name:      "软件二班",
		Num:       70,
		TeacherID: 1,
	}
	class03 := entity.Class{
		Name:      "软件三班",
		Num:       45,
		TeacherID: 1,
	}
	// 批量插入记录
	classes := []entity.Class{class01, class02, class03}
	db.Create(&classes)

	//课程
	course01 := entity.Course{
		Name:      "计算机网络",
		Credit:    3,
		Period:    16,
		TeacherID: 1,
	}
	course02 := entity.Course{
		Name:      "C语言程序设计",
		Credit:    3,
		Period:    16,
		TeacherID: 1,
	}
	course03 := entity.Course{
		Name:      "Go语言程序设计",
		Credit:    3,
		Period:    16,
		TeacherID: 1,
	}
	// 批量插入记录
	courses := []entity.Course{course01, course02, course03}
	db.Create(&courses)

	// 多对多添加记录
}

// select
func selectRcord(db *gorm.DB) {

	// 查询全部记录
	var status []entity.Class
	db.Find(&status)
	// 查询后 返回值为结构体的数组对象
	fmt.Println(status)

	// 基于String的Where语句
	db.Where("teacher_id = ?", 1).Find(&status)
	fmt.Println(status)
	var total int64
	// Count()接受指针类型的int64
	db.Model(&entity.Teacher{}).Where("Name = ?", "Asashishi").Count(&total)
	fmt.Println(total)

	// 基于struct/map的where语句
	db.Where(entity.Class{Name: "软件一班", TeacherID: 1}).Find(&status)
	fmt.Println(status)
	db.Where(map[string]interface{}{"Name": "Asashishi", "TeacherID": 1}).Find(&status)
	fmt.Println(status)

	// 查询单条记录
	class := entity.Class{}
	db.Take(&class)
	fmt.Println(class)
	db.First(&class)
	fmt.Println(class)
	db.Last(&class)
	fmt.Println(class)

	// 其他查询
	// 查询指定记录的指定字段
	var course []entity.Course
	db.Select("name,credit").Where("teacher_id = ?", 1).Find(&course)
	fmt.Println(course)
	// 查询指定记录的并忽略指定字段
	db.Omit("teacher_id").Where("teacher_id = ?", 1).Find(&course)
	fmt.Println(course)
	// 分页查询略过前一条取两条
	db.Order("teacher_id desc").Limit(2).Offset(1).Find(&course)
	fmt.Println(course)
	// 分组查询, 使用结构体承接分组
	type GroupedCourse struct {
		TeacherID int
		Count     int
	}
	var grouped []GroupedCourse
	db.Model(&entity.Class{}).Select("teacher_id,Count(*) as count").Group("teacher_id").Having("Count > ?", 1).Scan(&grouped)
	fmt.Println(grouped)
	// ...
}

// delete
func deleteRecord(db *gorm.DB) {
	var course entity.Course

	// 删除一条记录
	db.Where("name = ?", "计算机网络").Take(&course)
	fmt.Println(course)
	db.Delete(&course)

	// 按条件删除
	db.Where("Credit < ?", 3).Delete(&entity.Course{})

	// 全部删除
	db.Where("1 = 1").Delete(&entity.Course{})
}

// update
func updateRecord(db *gorm.DB) {

	// 先查后删
	var course entity.Course
	db.Where("name = ?", "C语言程序设计").Take(&course)
	fmt.Println(course)
	course.Name = "R#程序设计"
	db.Save(&course)
	fmt.Println(course)

	// 更新所有字段
	db.Model(&entity.Course{}).Where("1 = 1").Update("Credit", 3)
	// 更新表达式
	db.Model(&entity.Course{}).Where("1 = 1").Update("Credit", gorm.Expr("Credit + ?", 1))

	// 按条件更新字段
	db.Model(&entity.Course{}).Where("teacher_id", 1).Update("Credit", 5)

	// 通过结构体更新多个字段
	db.Model(&entity.Course{}).Where("teacher_id", 1).Updates(&entity.Course{Credit: 5, Period: 20})

	// 通过Map字典更新多个字段
	db.Model(&entity.Course{}).Where("teacher_id", 1).Updates(map[string]interface{}{"Credit": 5, "Period": 2})
}

func main() {

	// 定义数据源
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_learn?charset=utf8&parseTime=True&loc=Local"

	// 定义log对象并输出至控制台
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//// 超时阈值
			//SlowThreshold: time.Second,
			LogLevel: logger.Info,
		})

	// gorm处理MySQL链接和操作
	db, err0 := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 日志配置
		Logger: newLogger,
	})
	//db, err0 := gorm.Open(sqlite.Open("gorm_learn.db"), &gorm.Config{
	//	// 日志配置
	//	Logger: newLogger,
	//})
	if err0 != nil {
		log.Println(err0.Error())
	}
	/*
		// 迁移表 如果表不存在则会自动创建 会按照结构体映射创建
		err1 := db.AutoMigrate(&entity.Teacher{})
		if err1 != nil {
			log.Println(err1.Error())
		}
		err2 := db.AutoMigrate(&entity.Class{})
		if err2 != nil {
			log.Println(err2.Error())
		}
		err3 := db.AutoMigrate(&entity.Course{})
		if err3 != nil {
			log.Println(err3.Error())
		}
		err4 := db.AutoMigrate(&entity.Student{})
		if err4 != nil {
			log.Println(err4.Error())
		}
	*/

	//addRecord(db)
	//selectRcord(db)
	//deleteRecord(db)
	updateRecord(db)
}