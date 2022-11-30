/*
MongoDB Atlas Administration API

Testing DatabaseUsersApiService

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

func Test_openapi_DatabaseUsersApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DatabaseUsersApiService CreateOneDatabaseUserInOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.DatabaseUsersApi.CreateOneDatabaseUserInOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseUsersApiService RemoveOneDatabaseUserFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var databaseName string
        var username string

        resp, httpRes, err := apiClient.DatabaseUsersApi.RemoveOneDatabaseUserFromOneProject(context.Background(), groupId, databaseName, username).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseUsersApiService ReturnAllDatabaseUsersFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.DatabaseUsersApi.ReturnAllDatabaseUsersFromOneProject(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseUsersApiService ReturnOneDatabaseUserFromOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var databaseName string
        var username string

        resp, httpRes, err := apiClient.DatabaseUsersApi.ReturnOneDatabaseUserFromOneProject(context.Background(), groupId, databaseName, username).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseUsersApiService UpdateOneDatabaseUserInOneProject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var databaseName string
        var username string

        resp, httpRes, err := apiClient.DatabaseUsersApi.UpdateOneDatabaseUserInOneProject(context.Background(), groupId, databaseName, username).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}