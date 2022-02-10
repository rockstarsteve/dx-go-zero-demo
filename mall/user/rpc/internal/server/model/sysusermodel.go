package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserFieldNames          = builder.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysUserIdPrefix   = "cache:sysUser:id:"
	cacheSysUserNamePrefix = "cache:sysUser:name:"
)

type (
	SysUserModel interface {
		Insert(data *SysUser) (sql.Result, error)
		FindOne(id int64) (*SysUser, error)
		FindOneByName(name string) (*SysUser, error)
		Update(data *SysUser) error
		Delete(id int64) error
	}

	defaultSysUserModel struct {
		sqlc.CachedConn
		table string
	}

	SysUser struct {
		Id             int64          `db:"id"`               // 编号
		Name           string         `db:"name"`             // 用户名
		NickName       sql.NullString `db:"nick_name"`        // 昵称
		Avatar         sql.NullString `db:"avatar"`           // 头像
		Password       sql.NullString `db:"password"`         // 密码
		Salt           sql.NullString `db:"salt"`             // 加密盐
		Email          sql.NullString `db:"email"`            // 邮箱
		Mobile         sql.NullString `db:"mobile"`           // 手机号
		Status         sql.NullInt64  `db:"status"`           // 状态  0：禁用   1：正常
		DeptId         sql.NullInt64  `db:"dept_id"`          // 机构ID
		CreateBy       sql.NullString `db:"create_by"`        // 创建人
		CreateTime     time.Time      `db:"create_time"`      // 创建时间
		LastUpdateBy   sql.NullString `db:"last_update_by"`   // 更新人
		LastUpdateTime sql.NullTime   `db:"last_update_time"` // 更新时间
		DelFlag        int64          `db:"del_flag"`         // 是否删除  -1：已删除  0：正常
		JobId          sql.NullInt64  `db:"job_id"`           // 岗位Id
	}
)

func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &defaultSysUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_user`",
	}
}

func (m *defaultSysUserModel) Insert(data *SysUser) (sql.Result, error) {
	sysUserIdKey := fmt.Sprintf("%s%v", cacheSysUserIdPrefix, data.Id)
	sysUserNameKey := fmt.Sprintf("%s%v", cacheSysUserNamePrefix, data.Name)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
		return conn.Exec(query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.JobId)
	}, sysUserIdKey, sysUserNameKey)
	return ret, err
}

func (m *defaultSysUserModel) FindOne(id int64) (*SysUser, error) {
	sysUserIdKey := fmt.Sprintf("%s%v", cacheSysUserIdPrefix, id)
	var resp SysUser
	err := m.QueryRow(&resp, sysUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) FindOneByName(name string) (*SysUser, error) {
	sysUserNameKey := fmt.Sprintf("%s%v", cacheSysUserNamePrefix, name)
	var resp SysUser
	err := m.QueryRowIndex(&resp, sysUserNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", sysUserRows, m.table)
		if err := conn.QueryRow(&resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) Update(data *SysUser) error {
	sysUserIdKey := fmt.Sprintf("%s%v", cacheSysUserIdPrefix, data.Id)
	sysUserNameKey := fmt.Sprintf("%s%v", cacheSysUserNamePrefix, data.Name)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.DeptId, data.CreateBy, data.LastUpdateBy, data.LastUpdateTime, data.DelFlag, data.JobId, data.Id)
	}, sysUserIdKey, sysUserNameKey)
	return err
}

func (m *defaultSysUserModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	sysUserIdKey := fmt.Sprintf("%s%v", cacheSysUserIdPrefix, id)
	sysUserNameKey := fmt.Sprintf("%s%v", cacheSysUserNamePrefix, data.Name)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysUserIdKey, sysUserNameKey)
	return err
}

func (m *defaultSysUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysUserIdPrefix, primary)
}

func (m *defaultSysUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
