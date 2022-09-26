package auth

import (
	"github.com/kataras/iris/v12"
	"github.com/kmilodenisglez/github.template-srv.restapi.iris.go/lib"
	"github.com/kmilodenisglez/github.template-srv.restapi.iris.go/repo/db"
	"github.com/kmilodenisglez/github.template-srv.restapi.iris.go/schema"
	"github.com/kmilodenisglez/github.template-srv.restapi.iris.go/schema/dto"
)

type Provider interface {
	GrantIntent(userCredential *dto.UserCredIn, data interface{}) (*dto.GrantIntentResponse, *dto.Problem)
}

// region ======== EVOTE AUTHENTICATION PROVIDER =========================================

type ProviderDrone struct {
	// walletLocations string
	repo *db.RepoDrones
}

func (p *ProviderDrone) GrantIntent(uCred *dto.UserCredIn, options interface{}) (*dto.GrantIntentResponse, *dto.Problem) {
	// getting the users
	user, err := (*p.repo).GetUser(uCred.Username, true)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}
	checksum, _ := lib.Checksum("SHA256", []byte(uCred.Password))
	if user.Passphrase == checksum {
		return &dto.GrantIntentResponse{Identifier: user.Username, DID: user.Username}, nil
	}

	return nil, dto.NewProblem(iris.StatusUnauthorized, schema.ErrFile, schema.ErrCredsNotFound)
}

// endregion =============================================================================