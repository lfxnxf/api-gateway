package service

import (
	"context"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
	"github.com/lfxnxf/school/api-gateway/model"
	"go.uber.org/zap"
)

// 获取站点列表
func (s *Service) GetSitesList(ctx context.Context, atom *school_http.Atom, req model.GetSitesListReq) (interface{}, error) {
	log := logging.For(ctx, "func", "GetDivers",
		zap.Int64("uid", atom.Uid),
		zap.Any("req", req),
	)

	resp := model.GetSitesListResp{
		List: make([]model.SiteInfo, 0),
	}

	where := make(map[string]interface{}, 0)
	if req.Name != "" {
		where["name"] = req.Name
	}

	sitesModelList, err := s.dao.GetAllSites(ctx, where)
	if err != nil {
		log.Errorw("s.dao.GetAllSites error", zap.Error(err))
		return resp, err
	}

	//组装数据
	for _, v := range sitesModelList {
		info := model.SiteInfo{
			Id:        v.Id,
			Name:      v.Name,
			Longitude: v.Longitude,
			Latitude:  v.Latitude,
		}
		resp.List = append(resp.List, info)
	}

	log.Infow("success!")
	return resp, nil
}
