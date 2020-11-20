package dao

import (
	"github.com/go-xorm/xorm"
	"mycasbin/dao/models"
	"log"
	"context"
)

type UserDao struct{}

func (dao *UserDao) AddUser(ctx context.Context, session *xorm.Session, username, pwdHash string) (int64, error) {
	user := models.User{}
	user.PwdHash = pwdHash
	user.Name = username
	user.Group = ""
	user.Rank = 1
	user.Email = ""
	_, err := session.InsertOne(&user)
	if err != nil {
		return 0, err
	}
	return user.Ysid, nil
}

func (dao *UserDao) AddGroup(ctx context.Context,session *xorm.Session, userID int64, groupName string) (error) {
	user := dao.QueryByUserID(ctx, session, userID)
	if user != nil{
		log.Fatal("用户不存在，未用户添加组失败！")
	}
	user.Group = groupName
	session.Where("ysid = ?", user.Ysid)
	_, err := session.Cols("group").Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) AddRank(ctx context.Context,session *xorm.Session, userID int64, rank int) (error) {
	user := dao.QueryByUserID(ctx, session, userID)
	if user != nil{
		log.Fatal("用户不存在，未用户添加组失败！")
	}
	user.Rank = rank
	session.Where("ysid = ?", user.Ysid)
	_, err := session.Cols("rank").Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) QueryByUserID(ctx context.Context,session *xorm.Session, userID int64) *models.User {
	user := &models.User{}
	user.Ysid = userID
	has, err := session.Get(user)
	if err != nil || !has {
		return nil
	}
	return user
}

func (dao *UserDao) QueryByUserName(ctx context.Context,session *xorm.Session, userName string) *models.User {
	user := &models.User{}
	user.Name = userName
	has, err := session.Get(user)
	if err != nil || !has {
		return nil
	}
	return user
}

func (dao *UserDao) DeleteByUserID(ctx context.Context,session *xorm.Session, userID int64) error {
	user := &models.User{}
	user.Ysid = userID
	_, err := session.Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) DeleteByUserName(ctx context.Context,session *xorm.Session, userName string) error {
	user := &models.User{}
	user.Name = userName
	_, err := session.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
