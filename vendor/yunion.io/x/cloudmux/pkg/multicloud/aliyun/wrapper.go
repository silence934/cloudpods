// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"math/rand"
	"runtime/debug"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
)

func processCommonRequest(client *sdk.Client, req *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("client.ProcessCommonRequest error: %s", r)
			debug.PrintStack()
			response = nil
			jsonError := jsonutils.NewDict()
			jsonError.Add(jsonutils.NewString("SignatureNonceUsed"), "Code")
			err = errors.Error(jsonError.String())
		}
	}()
	if req.ApiName == "DescribeDBInstances" {
		i := rand.Intn(100)
		req.SetReadTimeout(1 * time.Minute)
		req.SetConnectTimeout(1 * time.Minute)
		log.Infof("bug fix log %d Domain:%v", i, req.Domain)
		log.Infof("bug fix log %d GetReadTimeout:%v", i, req.GetReadTimeout())
		log.Infof("bug fix log %d GetConnectTimeout:%v", i, req.GetConnectTimeout())
		response, err = client.ProcessCommonRequest(req)
		if err != nil {
			log.Infof("bug fix log %d err:%v", i, err)
		}
		return
	}

	return client.ProcessCommonRequest(req)
}
