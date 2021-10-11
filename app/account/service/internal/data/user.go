package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/lalifeier/vgo/app/account/service/internal/biz"
	"github.com/lalifeier/vgo/app/account/service/internal/data/ent/user"
	"github.com/lalifeier/vgo/pkg/crypto/bcrypt"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user-service/data")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	hash, err := bcrypt.GeneratePassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetPasswordHash(string(hash)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	po, err := r.data.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(ctx)
	if err != nil {
		return false, err
	}
	return bcrypt.ValidatePassword(u.Password, po.PasswordHash)
}
