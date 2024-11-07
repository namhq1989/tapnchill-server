package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
)

func (s Service) GetIpCity(ctx *appcontext.AppContext, ip string) (*string, error) {
	ctx.Logger().Info("get ip city", appcontext.Fields{"ip": ip})

	ctx.Logger().Text("find in caching")
	city, err := s.cachingRepository.GetIpCity(ctx, ip)
	if city != nil {
		ctx.Logger().Text("found in caching, respond")
		return city, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("not found in caching, call api")
	city, err = s.externalApiRepository.GetIpCity(ctx, ip)
	if err != nil {
		ctx.Logger().Error("failed to call api", err, appcontext.Fields{})
		return nil, err
	}
	if city == nil {
		ctx.Logger().Text("ip city not found")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("persist in caching")
	if err = s.cachingRepository.SetIpCity(ctx, ip, *city); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done get ip city")
	return city, nil
}
