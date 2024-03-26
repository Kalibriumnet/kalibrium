package client

import (
	"context"
	"github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kalibriumwalletd server, and returns the client instance
func Connect(address string) (pb.KalibriumwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kalibriumwallet daemon is not running, start it with `kalibriumwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKalibriumwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
