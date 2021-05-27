package gormtest

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pkgerr "github.com/pkg/errors"
	"time"
)

// https://www.jianshu.com/p/fcef8b425ce9
// https://blog.csdn.net/qq_28053177/article/details/82018708

func Gorm_test() {
	//查询全部
	//var (
	//	products []EsbProduct
	//	users    []EsbUser
	//	total    int
	//)
	//db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/rlgw_soft?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4")
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/code001?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	defer db.Close()
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 全局禁用表名复数
	db.SingularTable(true)
	// 自动检查 User 结构是否变化
	// db.AutoMigrate(&EsbProduct{})
	// db.AutoMigrate(&EsbProductStock{})
	// db.AutoMigrate(&EsbProductFile{})
	// db.AutoMigrate(&EsbFile{})
	//db.AutoMigrate(&RlgwDevice{})

	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)

	// // 创建
	//product := EsbProduct{Name: "商品8", Price: 800}
	//db.Create(&product)
	//fmt.Println(product.ID)

	user := EsbUser{Phone: "15210768111"}
	// db.Create(&user)

	cachePool := RlgwDevice{DeviceCode: "jnrl0001", DeviceState: 0, LastOnline: time.Now()}
	db.Create(&cachePool)

	// // 查询
	//product := EsbProduct{StockID: 1}
	//db.Model(&product).Find(&product)

	//deviceParams := &RlgwDeviceParams{ParamId: 1}
	//db.Model(&deviceParams).Find(&deviceParams)
	//db.Find(&deviceParams).Count(&total)
	// 查一个
	//db.First(&deviceParams, "param_id = ?", 1)

	//product := EsbProduct{}
	//db.Model(&product).Find(&product)
	//
	//db.Find(&products).Count(&total)
	// // 一对一
	//db.Model(&product).Related(&product.Stock, "Stock")
	// // 一对多
	// 查询一条记录
	//db.Model(&product).Preload("File").Related(&product.Files, "Files")

	//db.Preload("Roles.Resources").Preload("Roles").Find(&users)
	//db.Model(&users).Count(&total)

	// gDB := db.Model(&EsbProduct{})

	// gDB.Count(&total)

	// db.Preload("Stock").Find(&products)
	// db.Preload("Files").Preload("File").Find(&products)
	// db.Preload("Stock").Preload("Files.File").Preload("Files").Order("create_time ASC").Limit(2).Offset(0).Find(&products)
	// if err = pkgerr.WithStack(db.Preload("Stock").Preload("Files").Preload("File").Order("create_time DESC").Limit(2).Offset(3).Find(&products).Error); err != nil {
	// 	return
	// }
	// 修改
	//var dev RlgwDevice = RlgwDevice{DeviceCode: "jnrl0001"}
	//db.Where("device_code = ?", dev.DeviceCode).Find(&dev)
	//db.Preload("TrainDevice").Find(&dev)
	//db.Model(&dev).Related(&dev.TrainDevice, "DeviceCode")

	//var dev RlgwTrainDevice = RlgwTrainDevice{DeviceCode: "jnrl0001"}
	//db.Raw("SELECT train_code from rlgw_train_device WHERE rlgw_train_device.device_code = ?", dev.DeviceCode).Scan(&dev)

	pkgerr.WithStack(db.Model(&RlgwDevice{}).Where("device_code = ?", cachePool.DeviceCode).Updates(cachePool).Error)

	db.Model(&user).Update("ip", "127.0.0.1")
	// 删除

}
