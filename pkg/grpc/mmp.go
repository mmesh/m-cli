package grpc

import (
	"context"
	"io"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

// Control method implementation of NetworkAPI gRPC Service
func Control(client rpc.NetworkAPIClient, wg *sync.WaitGroup, waitc chan struct{}) error {
	// ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	// defer cancel()
	stream, err := client.Control(context.TODO())
	if err != nil {
		msg.Errorf("%v.Exec(_) = _, %v", client, err)
		wg.Done()
		return err
	}

	go func() {
		for {
			payload, err := stream.Recv()
			// if err == io.EOF
			if errors.Is(err, io.EOF) {
				// read done.
				waitc <- struct{}{}
				return
			}
			if err != nil {
				msg.Tracef("Failed to receive mmp payload: %v", err)
				waitc <- struct{}{}
				return
			}
			// msg.Debugf("Processing cmd payload from %s", payload.SrcID)

			if err := mmp.Processor(context.TODO(), payload); err != nil {
				// msg.Tracef("Unable to process mmp payload: %v", err)
				msg.Error(errors.Cause(err))
				waitc <- struct{}{}
				return
			}
		}
	}()

	go func() {
		for {
			payload := <-mmp.TxControlQueue
			if err := stream.Send(payload); err != nil {
				msg.Tracef("Failed to send mmp payload: %v", errors.Cause(err))
				waitc <- struct{}{}
				return
			}
		}
	}()

	mmp.TxControlQueue <- &mmsp.Payload{
		SrcID:       viper.GetString("mm.id"),
		PayloadType: mmsp.PayloadType_NODE_INIT,
		Node:        nil,
	}

	<-waitc
	_ = stream.CloseSend()
	wg.Done()

	return nil
}
