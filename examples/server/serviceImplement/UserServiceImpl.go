package serviceImplement

import "github.com/wangxingge/thrift_clientpool/examples/entity"

type UserServiceImpl struct {
}

func (srv *UserServiceImpl) GetUserBooks(userId string) (r []*entity.Book, err error) {
	return
}

func (srv *UserServiceImpl) GetUserInfo(userId string) (r *entity.User, err error) {
	return
}
func (srv *UserServiceImpl) GetAllUserInfo() (r []*entity.User, err error) {
	return
}

func (srv *UserServiceImpl) AddUser(userInfo *entity.User) (r bool, err error) {
	return
}

func (srv *UserServiceImpl) RemoveUser(userId string) (r bool, err error) {
	return
}

func (srv *UserServiceImpl) UpdateUserAvatar(userId string, avatar []byte) (r bool, err error) {
	return
}
