package service

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/mail"
	"github.com/cavelms/pkg/utils"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type authService interface {
	SignUp(ctx context.Context, fullName, email, password, role string) (*model.User, error)
	SignIn(ctx context.Context, email, password string) (*model.User, error)
	SignOut(ctx context.Context) error
	RefreshToken(ctx context.Context) (*model.User, error)
	VerifyEmail(ctx context.Context, id, code string, resend bool) (*model.User, error)
	ForgetPassword(email string) (*model.User, error)
	ResetPassword(email, password string) (*model.User, error)
	ChangePassword(email, token string) (*model.User, error)
	RequireAuth(ctx context.Context, obj interface{}, next graphql.Resolver, token *string) (res interface{}, err error)
}

type auth struct {
	claims *jwt.MapClaims
	*repository.Repository
}

func newAuthService(repo *repository.Repository) authService {
	return &auth{
		Repository: repo,
	}
}

func (a *auth) SignUp(ctx context.Context, fullName, email, password, role string) (*model.User, error) {
	hash, err := utils.EncryptPassword(password)
	if err != nil {
		return nil, err
	}

	var names = strings.Split(fullName, " ")
	user := &model.User{
		FirstName:    names[0],
		LastName:     names[1],
		FullName:     fullName,
		Email:        email,
		Role:         role,
		PasswordHash: hash,
	}

	err = a.DB.Create(user)
	if err != nil {
		return nil, err
	}

	code := utils.GenerateVerificationCode()
	err = a.RDBS.Set(user.ID, strings.TrimSpace(code), 600)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"code":     code,
		"fullname": fullName,
	}

	body, err := utils.ParseTemplate("signup", data)
	if err != nil {
		return nil, err
	}

	mail := mail.Mailer{
		ToAddrs: []string{email},
		Subject: "Account Activation",
		Body:    body,
	}

	err = a.Mail.Send(mail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *auth) SignIn(ctx context.Context, email, password string) (*model.User, error) {
	user := new(model.User)
	user.Email = email
	err := a.DB.FetchByEmail(user)
	if err != nil {
		return nil, err
	}

	isCorrect := utils.CheckPassword(user.PasswordHash, password)
	if !isCorrect {
		return nil, utils.ErrAuthenticationFailure
	}

	t, err := JWTAuthService().GenerateToken(*user, true)
	if err != nil {
		return nil, err
	}

	a.setCookie(ctx, t)

	user.Token = t.AccessToken
	user.TokenExpiredAt = t.AccessExpiresAt / int64(time.Second)
	return user, nil
}

func (a *auth) SignOut(ctx context.Context) error {
	c := ctx.Value(apiCtx("apiCtx")).(*gin.Context)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour),
	})
	return nil
}

func (a *auth) RefreshToken(ctx context.Context) (*model.User, error) {
	c := ctx.Value(apiCtx("apiCtx")).(*gin.Context)

	token, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}

	jwt := JWTAuthService()
	claims, err := jwt.ValidateRefreshToken(token)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:    claims["userId"].(string),
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}

	err = a.DB.FetchByID(&user)
	if err != nil {
		return nil, err
	}

	t, err := jwt.GenerateToken(user, false)
	if err != nil {
		return nil, err
	}

	user.Token = t.AccessToken
	user.TokenExpiredAt = t.AccessExpiresAt / int64(time.Second)
	return &user, nil
}

func (a *auth) VerifyEmail(ctx context.Context, id, code string, resend bool) (*model.User, error) {
	user := &model.User{}
	user.ID = id

	err := a.DB.FetchByID(user)
	if err != nil {
		return nil, err
	}

	if resend {
		c := utils.GenerateVerificationCode()
		err := a.sendCode(user, c)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	c, err := a.RDBS.Get(id)
	if err != nil {
		return nil, errors.New("error: Expired Code")
	}

	if c != code {
		return nil, errors.New("error: Invalid verification code")
	}

	user.IsVerified = true
	err = a.DB.UpdateOne(user)
	if err != nil {
		return nil, err
	}

	t, err := JWTAuthService().GenerateToken(*user, true)
	if err != nil {
		return nil, err
	}

	a.setCookie(ctx, t)

	user.Token = t.AccessToken
	user.TokenExpiredAt = t.AccessExpiresAt / int64(time.Second)
	return user, nil
}

func (a *auth) ResendCode(ctx context.Context, id string) (*model.User, error) { return nil, nil }
func (a *auth) ForgetPassword(email string) (*model.User, error)               { return nil, nil }
func (a *auth) ResetPassword(email, password string) (*model.User, error)      { return nil, nil }
func (a *auth) ChangePassword(email, token string) (*model.User, error)        { return nil, nil }

func (a *auth) setCookie(ctx context.Context, t *Token) {
	c := ctx.Value(apiCtx("apiCtx")).(*gin.Context)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    t.RefreshToken,
		HttpOnly: true,
		MaxAge:   int(t.RefreshExpiresAt / int64(time.Second)),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

}

func (a *auth) sendCode(user *model.User, code string) error {
	err := a.RDBS.Set(user.ID, strings.TrimSpace(code), 600)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"code":     code,
		"fullname": user.FullName,
	}

	body, err := utils.ParseTemplate("signup", data)
	if err != nil {
		return  err
	}

	mail := mail.Mailer{
		ToAddrs: []string{user.Email},
		Subject: "Account Activation",
		Body:    body,
	}

	err = a.Mail.Send(mail)
	if err != nil {
		return err
	}

	return nil
}

func (a *auth) RequireAuth(ctx context.Context, obj interface{}, next graphql.Resolver, token *string) (res interface{}, err error) {
	PREFIX := "Bearer "
	c := ctx.Value(apiCtx("apiCtx")).(*gin.Context)
	auth := c.GetHeader("Authorization")
	authToken := strings.TrimPrefix(auth, PREFIX)
	claims, err := JWTAuthService().ValidateAccessToken(authToken)
	if err != nil {
		a.claims = nil
		return nil, err
	}

	a.claims = &claims
	return next(ctx)
}
