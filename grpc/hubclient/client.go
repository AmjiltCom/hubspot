package hubclient

import (
	"context"
	"fmt"
	hubSpotProto "github.com/AmjiltCom/hubspot/grpc/hubSpotProto"
	"google.golang.org/grpc"

	"time"
)

func GegHubConfig(request *hubSpotProto.ConfigRequest, servicePath string) (*hubSpotProto.ConfigResponse, error) {
	fmt.Println(servicePath)
	conn, err := grpc.Dial(servicePath, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := hubSpotProto.NewHubSpotClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := c.GetConfig(ctx, request)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func GegHubContact(request *hubSpotProto.PartnerRequest, servicePath string) (*hubSpotProto.PartnerResponse, error) {
	fmt.Println(servicePath)
	conn, err := grpc.Dial(servicePath, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := hubSpotProto.NewHubSpotClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r, err := c.GetPartner(ctx, request)

	if err != nil {
		return nil, err
	}
	return r, nil
}
