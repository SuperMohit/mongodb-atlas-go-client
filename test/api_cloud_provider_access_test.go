/*
MongoDB Atlas Administration API

Testing CloudProviderAccessApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_CloudProviderAccessApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CloudProviderAccessApiService AuthorizeOneCloudProviderAccessRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var roleId string

        resp, httpRes, err := apiClient.CloudProviderAccessApi.AuthorizeOneCloudProviderAccessRole(context.Background(), groupId, roleId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudProviderAccessApiService CreateOneCloudProviderAccessRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.CloudProviderAccessApi.CreateOneCloudProviderAccessRole(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudProviderAccessApiService DeauthorizeOneCloudProviderAccessRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var cloudProvider string
        var roleId string

        resp, httpRes, err := apiClient.CloudProviderAccessApi.DeauthorizeOneCloudProviderAccessRole(context.Background(), groupId, cloudProvider, roleId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudProviderAccessApiService ReturnAllCloudProviderAccessRoles", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.CloudProviderAccessApi.ReturnAllCloudProviderAccessRoles(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CloudProviderAccessApiService ReturnCloudProviderAccessRole", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var roleId string

        resp, httpRes, err := apiClient.CloudProviderAccessApi.ReturnCloudProviderAccessRole(context.Background(), groupId, roleId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
