package nessielight

import (
	"fmt"
	"log"
	"os"

	"github.com/v2fly/v2ray-core/v4/common/uuid"
)

var authLog *log.Logger

type simpleTelegramAuthService struct {
	userManager *UserManager
	tokenDB     map[string]bool
}

func (r *simpleTelegramAuthService) GenToken() string {
	uid := uuid.New()
	token := uid.String()
	r.tokenDB[token] = true
	authLog.Printf("generate token %s", token)
	return token
}

func (r *simpleTelegramAuthService) Register(token string, id string) (User, error) {
	if !r.tokenDB[token] {
		return nil, fmt.Errorf("token %s invalid", token)
	}
	delete(r.tokenDB, token)
	user := (*r.userManager).NewUser(id)
	if err := (*r.userManager).SetUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func init() {
	authLog = log.New(os.Stderr, "[auth] ", log.LstdFlags|log.Lmsgprefix)
}
