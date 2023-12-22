package internal_connector

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	auth "snapser/cloudecode/snapserpb/auth"
	inventorypb "snapser/cloudecode/snapserpb/inventory"
	profilepb "snapser/cloudecode/snapserpb/profiles"
	statspb "snapser/cloudecode/snapserpb/statistics"
	storagepb "snapser/cloudecode/snapserpb/storage"
)

type SnapserServiceConnector struct {
	InventoryClient  inventorypb.InventoryServiceClient
	StatisticsClient statspb.StatisticsServiceClient
	StorageClient    storagepb.StorageServiceClient
	ProfileClient    profilepb.ProfilesServiceClient
	Auth             auth.AuthServiceClient
}

func New() (*SnapserServiceConnector, error) {
	pgs := &SnapserServiceConnector{}
	inventoryClient, err := createInventoryClient()
	if err != nil {
		return nil, err
	}
	pgs.InventoryClient = inventoryClient
	statisticsClient, err := createStatisticsClient()
	if err != nil {
		return nil, err
	}
	pgs.StatisticsClient = statisticsClient
	storageClient, err := createStorageClient()
	if err != nil {
		return nil, err
	}
	pgs.StorageClient = storageClient
	pgs.ProfileClient, err = createProfiledClient()
	if err != nil {
		return nil, err
	}
	pgs.StorageClient = storageClient
	authenticateSvc, err := createAuthenticate()
	if err != nil {
		return nil, err
	}
	pgs.Auth = authenticateSvc
	return pgs, nil
}

// We use GRPC internally to provide fast requests between services.  Users are authenticated at our API Gateway and then trusted once internal;
// The API Gateway will add/override a header called "User-Id" that is the authenticated trusted user and can be relied on internally.
func createInventoryClient() (inventorypb.InventoryServiceClient, error) {
	conn, err := grpc.Dial("service-inventory:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := inventorypb.NewInventoryServiceClient(conn)
	return client, nil
}

func createStatisticsClient() (statspb.StatisticsServiceClient, error) {
	conn, err := grpc.Dial("service-statistics:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := statspb.NewStatisticsServiceClient(conn)
	return client, nil
}
func createStorageClient() (storagepb.StorageServiceClient, error) {
	conn, err := grpc.Dial("service-storage:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := storagepb.NewStorageServiceClient(conn)
	return client, nil
}
func createProfiledClient() (profilepb.ProfilesServiceClient, error) {
	conn, err := grpc.Dial("service-profiles:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := profilepb.NewProfilesServiceClient(conn)
	return client, nil
}
func createAuthenticate() (auth.AuthServiceClient, error) {
	conn, err := grpc.Dial("service-auth:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := auth.NewAuthServiceClient(conn)
	return client, nil
}
