// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateUFileTokenRequest is request schema for CreateUFileToken action
type CreateUFileTokenRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"false"`

	// 令牌允许操作的bucket，默认*表示全部
	AllowedBuckets []string `required:"false"`

	// 令牌允许执行的操作，[ TOKEN_ALLOW_NONE , TOKEN_ALLOW_READ , TOKEN_ALLOW_WRITE , TOKEN_ALLOW_DELETE , TOKEN_ALLOW_LIST, TOKEN_ALLOW_IOP , TOKEN_ALLOW_DP  ]。默认TOKEN_ALLOW_NONE
	AllowedOps []string `required:"false"`

	// 令牌允许操作的key前缀，默认*表示全部
	AllowedPrefixes []string `required:"false"`

	// 令牌的超时时间点（时间戳），默认一天；注意：过期时间不能超过 4102416000
	ExpireTime *int `required:"false"`

	// 令牌名称
	TokenName *string `required:"true"`
}

// CreateUFileTokenResponse is response schema for CreateUFileToken action
type CreateUFileTokenResponse struct {
	response.CommonBase

	// 创建令牌的token_id
	TokenId string
}

// NewCreateUFileTokenRequest will create request of CreateUFileToken action.
func (c *UFileClient) NewCreateUFileTokenRequest() *CreateUFileTokenRequest {
	req := &CreateUFileTokenRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateUFileToken - 创建ufile令牌
func (c *UFileClient) CreateUFileToken(req *CreateUFileTokenRequest) (*CreateUFileTokenResponse, error) {
	var err error
	var res CreateUFileTokenResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateUFileToken", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}