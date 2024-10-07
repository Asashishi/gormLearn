**打印调试日志**
`db.Debug().Create(&struct{})`
**迁移表,表不存在于数据库中则会创建表**
`db.AutoMigrate(&struct{}):`
通过set设置附加参数 如:
`db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(struct[])`

**创建表**
`db.Migrator().CreateTable(&struct{})`
检测表结构体对应的表是否存在
`db.Migrator().HasTable(&struct{})`
依据表名检测是否存在
`db.Migrator().HasTable("TableName")`

**表字段操作**

**索引**
为字段添加索引
`db.Migrator().CreateIndex(&struct{},"Name")`
修改索引名称
`db.Migrator().RenameIndex(&struct{},"Before","After")`
为字段删除索引
`db.Migrator().DropIndex(&struct{},"Name")`
检查索引是否存在
`db.Migrator().HasIndex(&struct{},"Name"))`

**添加表单记录**
`db.Create(&struct)
Arr := []struct{struct1,struct2,struct3}`
指针传递结构体对象进行回写
`db.Create(&Arr)`
指定批量大小
`db.CreateInBatches(&Arr,100)`

**单表查询**
查询全部记录, 返回数组对象
`Arr := []struct{struct1,struct2,struct3}
result := db.Find(&Arr)`
查询完成后, 所有的数据会被映射到结构体并返回给Arr切片
`fmt.Println(result[0].Name)`
**单条记录用一个实例接收**
查询单条记录
``db.Take(&struct) [SELECT * FROM `table` LIMIT 1]``
查询根据主键id排序后的的升序/倒序第一条记录
``db.First/Last(&struct) [SELECT * FROM `table` ORDER BY `id` ASC/DESC LIMIT 1]``
**基于SQLString的条件查询字符串型**
`db.Where("id = example and num = example").Take(&struct)`
**基于struct和map的条件查询语句**
`db.Where(struct{attribute: value, attribute: value}).Find(&Arr)`
`db.Where(map[string]interface{}{"attribute": value, "attribute": value}).Find(&Arr)`

**其他查询**
Select语句查询指定字段
`db.Select("sqlAttribute","sqlAttribute",).Where("sqlAttribute = ?",value).Take(&struct)`
Omit语句忽略指定字段
`db.Omit("sqlAttribute","sqlAttribute",).Where("sqlAttribute = ?",value).Take(&struct)`
Not语句类似SQLString形式的Where语句,进行指定项忽略
`db.Not("sqlAttribute = ?",attribute).Find(&Arr)`
Or语句与SQLString形式的Where语句一起使用进行模糊查询
`db.Where("sqlAttribute = ?",attribute).Or("sqlAttribute like ?",attribute).Find(&Arr)`
Order语句排序, SQLString形式
`db.Where("sqlAttribute = ?",attribute).Order(sqlAttribute ASC,sqlAttribute ASC).Find(&Arr)`
Count语句必须使用Model()传递接收映射的结构体对象
`var total int64
db.Model(&struct).Where("Name = ?",example).Count(&total)`
Limit, Offset语句进行分页查询,取十条跳过前一条
`db.Order(sqlAttribute DESC).Limit(10).Offset(1).Find(&Arr)`
Group Having分组查询, Group必须与Select进行连用, SQLString形式
使用新定义的结构体承接分组
`type Grouped struct{
    sqlAttribute type,
    ...
}
var grouped []Grouped
db.Model(&struct{}).Select("SQLString").Group("sqlAttribute").Having("sqlAttribute > ?“, 1).Scan(&Arr)`

**删除**
`db.Migrator().DropTable(&struct{})
db.Migrator().DropTable("TableName")`
删除字段
`db.Migrator().DropColumn(&struct{},"Name")`
删除一条记录
`struct := struct{sqlAttribute: value}
db.Take(&struct)
db.Delete(&struct)`
条件删除
`db.Where(SQLString).Delete(struct{})`
删除所有记录
`db.Where("1 = 1").Delete(&struct)`

**查询**
Save语句更新某条记录的所有字段
`variable := struct{}
db.First(&variable)
variable.attribute = "value"
ab.Save(&variable)`
Update语句基于主键更新某个字段
`db.Model(&variable).Update("sqlAttribute","value")`
Update语句更新所有字段值
`db.Model(&variable).Where("1 = 1").Update("sqlAttribute","value")`
Update Where条件更新
`db.Model(&variable).Where("sqlAttribute = ?","value").Update("sqlAttribute","value")`
**通过struct更新多个字段值, 零值字段不会被更新**
`db.Model(&struct{}).Where("sqlAttribute = ?","value").Updates(struct{attribute: value,attribute: value})`
**通过map更新多个字段, 零值字段也会更新**
`db.Model(&struct{}).Where("sqlAttribute = ?","value").Updates(map[string]interface{}{"attribute": value, "attribute": value})`
**更新表达式**
字段值自增
`Update("attribute",gorm.Expr("attribute+1"))`
eg:
`db.Model(&struct{}).Update("attribute", gormExpr("attribute+1"))`