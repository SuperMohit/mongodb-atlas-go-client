/*
MongoDB Atlas Administration API

Testing AtlasSearchApiService

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

func Test_openapi_AtlasSearchApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AtlasSearchApiService CreateOneAtlasSearchIndex", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.AtlasSearchApi.CreateOneAtlasSearchIndex(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AtlasSearchApiService RemoveOneAtlasSearchIndex", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var indexId string

        resp, httpRes, err := apiClient.AtlasSearchApi.RemoveOneAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AtlasSearchApiService ReturnAllAtlasSearchIndexesForOneCollection", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var collectionName string
        var databaseName string

        resp, httpRes, err := apiClient.AtlasSearchApi.ReturnAllAtlasSearchIndexesForOneCollection(context.Background(), groupId, clusterName, collectionName, databaseName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AtlasSearchApiService ReturnAllUserDefinedAnalyzersForOneCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.AtlasSearchApi.ReturnAllUserDefinedAnalyzersForOneCluster(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AtlasSearchApiService ReturnOneAtlasSearchIndex", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var indexId string

        resp, httpRes, err := apiClient.AtlasSearchApi.ReturnOneAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AtlasSearchApiService UpdateAllUserDefinedAnalyzersForOneCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string

        resp, httpRes, err := apiClient.AtlasSearchApi.UpdateAllUserDefinedAnalyzersForOneCluster(context.Background(), groupId, clusterName).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AtlasSearchApiService UpdateOneAtlasSearchIndex", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var clusterName string
        var indexId string

        resp, httpRes, err := apiClient.AtlasSearchApi.UpdateOneAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}