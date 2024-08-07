# High-Level Design and Implementation

This Go program implements a basic gRPC server for an Ad Service. The service provides functinality to record click events and retrieve the top clicked ads. 

## Key Components
- gRPC Server Setup
- Service Definitions
- Request and Response Structs
- Business Logic
- Main Function

1. gRPC Server Setup
```golang
// The program initializes a gRPC server that listens for client requetsts on port :50051
const (
    PORT = ":50051"
)

// main() setups up a listener on the specified TCP port and registers the AdServiceServer with the gRPC server
func main(){
    	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterAdServiceServer(s, NewAdServiceServer())

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
```

2. Service Definitions
```golang
type Empty struct{}
type UnimplementedAdServiceServer struct{}
// The AdServiceServer struct implements service methods that correspond to recording click events and fetching top ads.
// UnimplementedAdServiceServer is embedded as placeholder, generated by the gRPC framework when defining the service.
type AdServiceServer struct {
    UnimplementedAdServiceServer
}
```

3. Request and Response Structs
```golang
// ClickEvent represents a click on an ad
type ClickEvent struct {
    AdID string
    UserID string
    Timestamp int64
}

// TopAdsRequest represents structure for fetching the top clicked ads
type TopAdsRequest struct {
    AdCategory string
    Limit int
}

// TopAdsResponse represents structure containing the aggregated results
type TopAdsResponse struct {
        Ads []AggregatedResult
}

// AggregatedResult represents an individual ad's click data in the response
type AggregatedResult struct {
    AdID string 
    Clicks int 
    Timestamp int64
}
```

4. Business Logic
```golang
// RecordClick records a click event in a map that tracks counts by AdID
var clicksMap map[string]int

func (s *AdServiceServer) RecordClicks(ctx context.Context, event *ClickEvent) (*Empty, error) {
    clicksMap[event.AdID] += 1
    return &Empty{}, nil
}

// GetTopAds returns the top-clicked ads. Currently, it returns a hardcoded set of results for simplicity
func (s *AdServiceServer) GetTopAds(ctx context.Context, req *TopAdsRequest) (*TopAdsResponse) {
    //TODO implement logic to fetch top ads
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
5. Registering the Service
```golang
// finally, we register the AdServiceServer with the gRPC server
func RegisterAdServiceServer(s *grpc.Server, srv *AdServiceServer) {
    // this is where we hoop up our server methods with the gRPC server.
    // normally, this method is auto-generated by protc, but we'll simulate it here.
    log.Println("AdServiceServer registered.")
}
```