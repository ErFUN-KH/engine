// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/user/proto/user.proto

package userpb

import github_com_fzerorubigd_balloon_pkg_postgres_model "github.com/fzerorubigd/balloon/pkg/postgres/model"
import time "time"
import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
import context "context"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/fzerorubigd/protobuf/extra"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const (
	UserSchema    = "aaa"
	UserTable     = "users"
	UserTableFull = UserSchema + "." + UserTable
)

type Manager struct {
	github_com_fzerorubigd_balloon_pkg_postgres_model.Manager
}

func NewManager() *Manager {
	return &Manager{}
}

func NewManagerFromTransaction(tx github_com_fzerorubigd_balloon_pkg_postgres_model.DBX) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)

	if err != nil {
		return nil, err
	}

	return m, nil
}

/*
main.modelData{
    table:     "users",
    schema:    "aaa",
    model:     "User",
    receiver:  "u",
    dbFields:  {"id", "email", "password", "status", "created_at", "updated_at", "last_login"},
    goFields:  {"Id", "Email", "Password", "Status", "CreatedAt", "UpdatedAt", "LastLogin"},
    types:     {"created_at":"Timestamp", "updated_at":"Timestamp", "last_login":"Timestamp"},
    createdAt: true,
    updatedAt: true,
    hasID:     true,
}
*/

func (m *Manager) CreateUser(ctx context.Context, u *User) error {
	var err error
	now := github_com_gogo_protobuf_types.TimestampNow()
	*u.CreatedAt = *now
	*u.UpdatedAt = *now
	func(in interface{}) {
		if o, ok := in.(interface{ PreInsert() }); ok {
			o.PreInsert()
		}
	}(u)
	CreatedAt, err := github_com_gogo_protobuf_types.TimestampFromProto(u.CreatedAt)
	if err != nil {
		return err
	}
	UpdatedAt, err := github_com_gogo_protobuf_types.TimestampFromProto(u.UpdatedAt)
	if err != nil {
		return err
	}
	LastLogin, err := github_com_gogo_protobuf_types.TimestampFromProto(u.LastLogin)
	if err != nil {
		return err
	}
	q := `INSERT INTO aaa.users(email, password, status, created_at, updated_at, last_login) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	row := m.GetDbMap().QueryRowxContext(ctx, q, u.Email, u.Password, u.Status, CreatedAt, UpdatedAt, LastLogin)
	return row.Scan(&u.Id)
}

func (m *Manager) UpdateUser(ctx context.Context, u *User) error {
	var err error
	now := github_com_gogo_protobuf_types.TimestampNow()
	*u.UpdatedAt = *now
	func(in interface{}) {
		if o, ok := in.(interface{ PreUpdate() }); ok {
			o.PreUpdate()
		}
	}(u)
	CreatedAt, err := github_com_gogo_protobuf_types.TimestampFromProto(u.CreatedAt)
	if err != nil {
		return err
	}
	UpdatedAt, err := github_com_gogo_protobuf_types.TimestampFromProto(u.UpdatedAt)
	if err != nil {
		return err
	}
	LastLogin, err := github_com_gogo_protobuf_types.TimestampFromProto(u.LastLogin)
	if err != nil {
		return err
	}
	q := `UPDATE aaa.users SET email = $1, password = $2, status = $3, created_at = $4, updated_at = $5, last_login = $6 WHERE id = $7`
	_, err = m.GetDbMap().ExecContext(ctx, q, u.Email, u.Password, u.Status, CreatedAt, UpdatedAt, LastLogin, u.Id)
	return err
}

func (m *Manager) GetUserByPrimary(ctx context.Context, id int64) (*User, error) {
	q := `SELECT id, email, password, status, created_at, updated_at, last_login FROM aaa.users WHERE id = $1`
	row := m.GetDbMap().QueryRowxContext(ctx, q, id)

	var u User
	var CreatedAt time.Time
	var UpdatedAt time.Time
	var LastLogin time.Time
	err := row.Scan(&u.Id, &u.Email, &u.Password, &u.Status, &CreatedAt, &UpdatedAt, &LastLogin)
	if err != nil {
		return nil, err
	}
	u.CreatedAt, _ = github_com_gogo_protobuf_types.TimestampProto(CreatedAt)
	u.UpdatedAt, _ = github_com_gogo_protobuf_types.TimestampProto(UpdatedAt)
	u.LastLogin, _ = github_com_gogo_protobuf_types.TimestampProto(LastLogin)
	return &u, nil
}
