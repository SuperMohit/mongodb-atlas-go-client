/*
MongoDB Atlas Administration API

Testing ProgrammaticAPIKeysApiService

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

func Test_openapi_ProgrammaticAPIKeysApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ProgrammaticAPIKeysApiService AssignOneOrganizationApiKeyToOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.AssignOneOrganizationApiKeyToOneProject(context.Background(), groupId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService CreateAccessListEntriesForOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.CreateAccessListEntriesForOneOrganizationApiKey(context.Background(), orgId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService CreateAndAssignOneOrganizationApiKeyToOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.CreateAndAssignOneOrganizationApiKeyToOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService CreateOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.CreateOneOrganizationApiKey(context.Background(), orgId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService RemoveOneAccessListEntryForOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var apiUserId string
        var ipAddress string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.RemoveOneAccessListEntryForOneOrganizationApiKey(context.Background(), orgId, apiUserId, ipAddress).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService RemoveOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.RemoveOneOrganizationApiKey(context.Background(), orgId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService ReturnAllAccessListEntriesForOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.ReturnAllAccessListEntriesForOneOrganizationApiKey(context.Background(), orgId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService ReturnAllOrganizationApiKeys", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeys(context.Background(), orgId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService ReturnAllOrganizationApiKeysAssignedToOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeysAssignedToOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService ReturnOneAccessListEntryForOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var ipAddress string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.ReturnOneAccessListEntryForOneOrganizationApiKey(context.Background(), orgId, ipAddress, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService ReturnOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.ReturnOneOrganizationApiKey(context.Background(), orgId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService SetApiUserGroupRoles", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.SetApiUserGroupRoles(context.Background(), groupId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService UnassignOneOrganizationApiKeyFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.UnassignOneOrganizationApiKeyFromOneProject(context.Background(), groupId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProgrammaticAPIKeysApiService UpdateOneOrganizationApiKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var orgId string
        var apiUserId string

        resp, httpRes, err := apiClient.ProgrammaticAPIKeysApi.UpdateOneOrganizationApiKey(context.Background(), orgId, apiUserId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
