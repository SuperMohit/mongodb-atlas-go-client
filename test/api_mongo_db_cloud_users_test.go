/*
MongoDB Atlas Administration API

Testing MongoDBCloudUsersApiService

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

func Test_openapi_MongoDBCloudUsersApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MongoDBCloudUsersApiService CreateOneCloudUser", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MongoDBCloudUsersApi.CreateOneCloudUser(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MongoDBCloudUsersApiService ReturnOneCloudUserUsingItsId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userId string

        resp, httpRes, err := apiClient.MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsId(context.Background(), userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MongoDBCloudUsersApiService ReturnOneCloudUserUsingItsUsername", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var userName string

        resp, httpRes, err := apiClient.MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsUsername(context.Background(), userName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
