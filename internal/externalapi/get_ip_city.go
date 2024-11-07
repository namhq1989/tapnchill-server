package externalapi

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
)

type GetIpCityResult struct {
	City string
}

type getCityByIpApiResult struct {
	City string `json:"city"`
}

func (ea ExternalApi) GetIpCity(ctx *appcontext.AppContext, ip string) (*GetIpCityResult, error) {
	var apiResult = &getCityByIpApiResult{}

	_, err := ea.locationClient.R().
		SetQueryParams(map[string]string{
			"token": ea.ipInfoToken,
		}).
		SetResult(&apiResult).
		Get(fmt.Sprintf("/%s", ip))

	if err != nil {
		ctx.Logger().Error("[externalapi] error when get ip city", err, appcontext.Fields{})
		return nil, err
	}

	if apiResult == nil || apiResult.City == "" {
		return nil, nil
	}

	return &GetIpCityResult{
		City: apiResult.City,
	}, nil
}
