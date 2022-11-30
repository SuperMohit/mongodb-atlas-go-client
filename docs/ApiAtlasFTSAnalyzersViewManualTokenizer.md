# ApiAtlasFTSAnalyzersViewManualTokenizer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxGram** | **int32** | Characters to include in the longest token that Atlas Search creates. | 
**MinGram** | **int32** | Characters to include in the shortest token that Atlas Search creates. | 
**Type** | **string** | Human-readable label that identifies this tokenizer type. | 
**Group** | **int32** | Index of the character group within the matching expression to extract into tokens. Use &#x60;0&#x60; to extract all character groups. | 
**Pattern** | **string** | Regular expression to match against. | 
**MaxTokenLength** | Pointer to **int32** | Maximum number of characters in a single token. Tokens greater than this length are split at this length into multiple tokens. | [optional] [default to 255]

## Methods

### NewApiAtlasFTSAnalyzersViewManualTokenizer

`func NewApiAtlasFTSAnalyzersViewManualTokenizer(maxGram int32, minGram int32, type_ string, group int32, pattern string, ) *ApiAtlasFTSAnalyzersViewManualTokenizer`

NewApiAtlasFTSAnalyzersViewManualTokenizer instantiates a new ApiAtlasFTSAnalyzersViewManualTokenizer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSAnalyzersViewManualTokenizerWithDefaults

`func NewApiAtlasFTSAnalyzersViewManualTokenizerWithDefaults() *ApiAtlasFTSAnalyzersViewManualTokenizer`

NewApiAtlasFTSAnalyzersViewManualTokenizerWithDefaults instantiates a new ApiAtlasFTSAnalyzersViewManualTokenizer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetMaxGram() int32`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetMaxGramOk() (*int32, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) SetMaxGram(v int32)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetMinGram() int32`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetMinGramOk() (*int32, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) SetMinGram(v int32)`

SetMinGram sets MinGram field to given value.


### GetType

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) SetType(v string)`

SetType sets Type field to given value.


### GetGroup

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetPattern

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMaxTokenLength

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetMaxTokenLength() int32`

GetMaxTokenLength returns the MaxTokenLength field if non-nil, zero value otherwise.

### GetMaxTokenLengthOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) GetMaxTokenLengthOk() (*int32, bool)`

GetMaxTokenLengthOk returns a tuple with the MaxTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenLength

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) SetMaxTokenLength(v int32)`

SetMaxTokenLength sets MaxTokenLength field to given value.

### HasMaxTokenLength

`func (o *ApiAtlasFTSAnalyzersViewManualTokenizer) HasMaxTokenLength() bool`

HasMaxTokenLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


