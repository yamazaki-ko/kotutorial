package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	user := os.Getenv("MYSQL_ROOT_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := "godockerDB"
	db := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		db,
	)
	fmt.Println(dsn)

	return gorm.Open(mysql.Open("root:root_password@tcp(godockerDB)/dev_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
}

type User struct {
	ID        uint `gorm:"primarykey"`
	UserName  string
	LoginId   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AuthLogin(loginId string, password string) bool {
	db, err := GetDB()
	if err != nil {
		fmt.Println(err)
		return false
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer sqlDb.Close()

	var u User
	db.First(&u, "login_id = ?", loginId)

	return u.Password == password
}

/*

func Create() {
	// テーブル名
	tableName := "db_dev.users"

	u := &User{}

	//テーブル作成
	err := u.Migrate()
	if err != nil {
		fmt.Println(err)
		return
	}

	// パラメータ設定
	u.UserId = "00001"
	u.UserName = "myfuns"
	u.Password = "password"
	// データ作成
	err = u.Create(tableName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// データ読込
	user, err := u.Read(tableName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.UserId)
	fmt.Println(user.UserName)
	fmt.Println(user.Password)

	// データ更新
	u.UserName = "ourfuns"
	err = u.Update(tableName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// データ読込
	user, err = u.Read(tableName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.UserId)
	fmt.Println(user.UserName)
	fmt.Println(user.Password)

	// データ削除
	err = u.Delete(tableName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// データ読込
	user, err = u.Read(tableName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.UserId)
	fmt.Println(user.UserName)
	fmt.Println(user.Password)
}

func getDB() (*gorm.DB, error) {
	dsn := "root:root_password@tcp(godockerDB)/db_dev?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

type User struct {
	UserId   string
	UserName string
	Password string
}

func (u *User) Migrate() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDb.Close()
	err = db.AutoMigrate(u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *User) Create(tableName string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	err = db.Table(tableName).Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Read(tableName string) (User, error) {
	user := User{}
	db, err := getDB()
	if err != nil {
		return user, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return user, err
	}
	defer sqlDb.Close()

	result := []*User{}
	err = db.Table(tableName).Select("user_id,user_name,password").Where("user_id = ?", u.UserId).Scan(&result).Error
	if err != nil {
		return user, err
	}
	for _, r := range result {
		user.UserId = r.UserId
		user.UserName = r.UserName
		user.Password = r.Password
	}
	return user, nil
}
func (u *User) Update(tableName string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	err = db.Table(tableName).Where("user_id = ?", u.UserId).Update("user_name", u.UserName).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *User) Delete(tableName string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	err = db.Table(tableName).Where("user_id = ?", u.UserId).Delete(u).Error
	if err != nil {
		return err
	}
	return nil
}
*/
