package tutorial

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Tutorial2() {
	// テーブル名
	tableName := "db_dev.users"

	u := &User{}
	// パラメータ設定
	u.UserId = "00001"
	u.UserName = "myfuns"
	u.Password = "password"
	// データ作成
	err := u.Create(tableName)
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
	dsn := "root:root_password@tcp(localhost:3306)/db_dev?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

type User struct {
	UserId   string
	UserName string
	Password string
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
