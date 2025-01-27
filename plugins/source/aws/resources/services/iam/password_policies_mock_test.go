package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIamPasswordPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.PasswordPolicy{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().GetAccountPasswordPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetAccountPasswordPolicyOutput{
			PasswordPolicy: &g,
		}, nil)
	return client.Services{
		Iam: m,
	}
}

func TestIamPasswordPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, PasswordPolicies(), buildIamPasswordPolicies, client.TestOptions{})
}
