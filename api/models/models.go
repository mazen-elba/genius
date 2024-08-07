package models

import (
	"context"
	"database/sql"
)

// gRPC server structure
type Server struct{}

type AdClick struct {
	AdID      string
	UserID    string
	UniqueID  string
	Timestamp int64
}

type AdData struct {
	AdID      string
	Clicks    int32
	Timestamp int64
}

// TopAdsRequest represents structure for fetching the top clicked ads
type TopAdsRequest struct {
	AdCategory string
	Limit      int
}

// TopAdsResponse represents structure containing the aggregated results
type TopAdsResponse struct {
	Ads []AdData
}

// gRPC Query Service definition
type QueryService struct {
	// forward compatibility
	// GetTopAds(context.Context, *TopAdsRequest)(*TopAdsResponse, err)
}

// GetTopAds returns a list of top ads by click count
func (s *QueryService) GetTopAds(ctx context.Context, req *TopAdsRequest) (*TopAdsResponse, error) {
	var db *sql.DB
	rows, err := db.Query("SELECT ad_id, clicks FROM ad_stats ORDER BY clicks DESC LIMIT $1", req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []AdData
	for rows.Next() {
		var ad AdData
		err = rows.Scan(&ad.AdID, &ad.Clicks)
		if err != nil {
			return nil, err
		}
	}

	return &TopAdsResponse{Ads: ads}, nil
}
