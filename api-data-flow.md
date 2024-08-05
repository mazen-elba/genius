# Data Flow

The data flow in this gRPC Ad Service involves below stages:

1. Client Sends Requests
```json
/**
A client initiates a request to either record a click event or fetch top ads.

These requests are encapsulated in specific messages such as `ClickEvent` for recording clicks and `TopAdsRequest` for fetching top ads.
*/

// Click Event Request
{
    "AcID": "ad_123",
    "UserID": "user_456",
    "Timestamp": 123156848
}

// Top Ads Request
{
    "AdCategory": "electronics",
    "Limit": 100
}
```

2. Server Receives Request
```golang
// Upon receving the request, the gRPC server forwards the request to the appropriate service method implemented in the `AdServiceServer` struct

// RecordClick handles incoming click events
func (s *AdServiceServer) RecordClick(ctx context.Context, event *ClickEvent) (*Empty, err) {
    clicksMap[event.AdID] += 1
    return &Empty{}, nil
}

// GetTopAds returns the top 100 most clicked ads
func (s *AdServiceServer) GetTopAds(ctx context.Context, req *TopAdsRequest) (*TopAdsResponse) {
	ads := []AggregatedResult{
		{AdID: "ad_123", Clicks: 100, Timestamp: time.Now().Unix()},
		{AdID: "ad_456", Clicks: 200, Timestamp: time.Now().Add(-24 * time.Hour).Unix()},
		{AdID: "ad_789", Clicks: 150, Timestamp: time.Now().Add(-48 * time.Hour).Unix()},
	}

    // sort ads in-place
	sort.Slice(ads, func(i, j int) bool { return ads[i].Clicks > ads[j].Clicks })

	return &TopAdsResponse{Ads: ads}, nil
}
```

3. Server Processes Business Logic
```golang
// The service methods implement busienss logic to handle the requests.

// RecordingClick events updates an in-memory map `clicksMap` with the count of the clicks per ad
var clicksMap map[string]int

func (s *AdServiceServer) RecordClicks(ctx context.Context, event *ClickEvent) (*Empty, error) {
    clicksMap[event.AdID] += 1
    return &Empty{}, nil
}
```

4. Server Sends Response