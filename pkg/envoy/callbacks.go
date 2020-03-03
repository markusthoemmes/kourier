/*
 Copyright 2020 The Knative Authors
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

*/

package envoy

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

//// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
//OnStreamOpen(context.Context, int64, string) error
//// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
//OnStreamClosed(int64)
//// OnStreamRequest is called once a request is received on a stream.
//// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
//OnStreamRequest(int64, *v2.DiscoveryRequest) error
//// OnStreamResponse is called immediately prior to sending a response on a stream.
//OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse)
//// OnFetchRequest is called for each Fetch request. Returning an error will end processing of the
//// request and respond with an error.
//OnFetchRequest(context.Context, *v2.DiscoveryRequest) error
//// OnFetchResponse is called immediately prior to sending a response.
//OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse)

func (cb *Callbacks) Report() {
}
func (cb *Callbacks) OnStreamOpen(ctx context.Context, id int64, typ string) error {
	cb.Logger.Infof("OnStreamOpen %d open for %s", id, typ)
	return nil
}
func (cb *Callbacks) OnStreamClosed(id int64) {
	cb.Logger.Infof("OnStreamClosed %d closed", id)
}

func (cb *Callbacks) OnStreamRequest(streamid int64, req *v2.DiscoveryRequest) error {
	if req.ErrorDetail != nil {
		if cb.OnError != nil {
			cb.OnError()
		}

		cb.Logger.Infof("OnStreamRequest Error Node : %v <---> StreamId : %d <---> Error : Code -> %v"+
			" <------> Message"+
			" -> %v <---------> Details -> %v", req.Node.Id, streamid, req.ErrorDetail.Code, req.ErrorDetail.Message, req.ErrorDetail.Details)
		return fmt.Errorf("OnStreamRequest Error Node : %v <---> StreamId : %d <---> Error : %v", req.Node.Id, streamid,
			req.ErrorDetail)
	}
	return nil
}
func (cb *Callbacks) OnStreamResponse(i int64, request *v2.DiscoveryRequest, response *v2.DiscoveryResponse) {
}
func (cb *Callbacks) OnFetchRequest(ctx context.Context, req *v2.DiscoveryRequest) error {
	return nil
}
func (cb *Callbacks) OnFetchResponse(req *v2.DiscoveryRequest, resp *v2.DiscoveryResponse) {
}

type Callbacks struct {
	Logger  *zap.SugaredLogger
	OnError func()
}
