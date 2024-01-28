package service

import (
	"context"

	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth/repo"
	cargomarketbid "github.com/andhikasamudra/fiber-starter-pack/protobuf"
	"github.com/google/uuid"
)

func (s *AuthService) GetProfileRPC(request *cargomarketbid.GetUserProfileRequest) (*cargomarketbid.GetUserProfileResponse, error) {
	ctxBg := context.TODO()

	userData, err := s.AuthModel.ReadUser(ctxBg, s.DbAdapter, repo.GetUserRequest{
		GUID: uuid.MustParse(request.UserGuid),
	})
	if err != nil {
		return nil, err
	}

	var roles []string
	for _, i := range userData.Roles {
		roles = append(roles, i.RoleKey)
	}

	return &cargomarketbid.GetUserProfileResponse{
		Email:       userData.Email,
		IsActive:    userData.IsActive,
		Roles:       roles,
		PicName:     userData.Profile.Name,
		PhoneNumber: userData.Profile.PhoneNumber,
		CompanyName: userData.Profile.CompanyName,
	}, nil
}
