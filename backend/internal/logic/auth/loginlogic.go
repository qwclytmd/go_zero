package auth

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"next.com/next/backend/internal/svc"
	"next.com/next/backend/internal/types"
	"next.com/next/models"
	"next.com/next/package/jwt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	var currentManager models.Managers
	if err = l.svcCtx.BackendDB.Model(&models.Managers{}).First(&currentManager).Error; err != nil {
		logx.Error(err)
		return nil, errors.New("用户名或密码错误")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(currentManager.Password), []byte(req.Password)); err != nil {
		logx.Error(err)
		return nil, errors.New("用户名或密码错误")
	}

	//token
	payload := jwt.Payload{UserId: currentManager.Id, UserName: currentManager.Username}
	token, err := jwt.CreateToken(payload, l.svcCtx.Config.Auth.AccessSecret, 24)
	if err != nil {
		return nil, err
	}

	//menus
	menus, err := l.GetPermissions(currentManager.RoleId)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{
		Token: token,
		Menus: menus,
	}

	return
}

// 获取权限
func (l *LoginLogic) GetPermissions(roleId int64) ([]*models.Menus, error) {
	var role models.Roles
	if err := l.svcCtx.BackendDB.Model(&models.Roles{}).Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}

	var menuInfo []models.Menus
	if err := l.svcCtx.BackendDB.Model(&models.Menus{}).Where("id in ? AND status = 1", strings.Split(role.Permission, ",")).Order("pid").Find(&menuInfo).Error; err != nil {
		return nil, err
	}

	menuTree := buildMenuTree(menuInfo, 0)

	return menuTree, nil
}

func buildMenuTree(menus []models.Menus, parentId int64) []*models.Menus {
	var tree []*models.Menus
	for _, menu := range menus {
		if menu.Pid == parentId {
			child := &menu
			child.Children = buildMenuTree(menus, menu.Id)
			tree = append(tree, child)
		}
	}
	return tree
}
