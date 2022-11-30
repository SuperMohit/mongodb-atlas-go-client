# ApiAtlasFTSAnalyzersViewManualTokenFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalTokens** | Pointer to **string** | Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose &#x60;include&#x60; if you want to support queries on both the original tokens as well as the converted forms.   Choose &#x60;omit&#x60; if you want to query only on the converted forms of the original tokens. | [optional] [default to "include"]
**Type** | **string** | Human-readable label that identifies this token filter type. | 
**MaxGram** | **int32** | Value that specifies the maximum length of generated n-grams. This value must be greater than or equal to **minGram**. | 
**MinGram** | **int32** | Value that specifies the minimum length of generated n-grams. This value must be less than or equal to **maxGram**. | 
**TermNotInBounds** | Pointer to **string** | Value that indicates whether to index tokens shorter than **minGram** or longer than **maxGram**. | [optional] [default to "omit"]
**NormalizationForm** | Pointer to **string** | Normalization form to apply. | [optional] [default to "nfc"]
**Max** | Pointer to **int32** | Number that specifies the maximum length of a token. Value must be greater than or equal to **min**. | [optional] [default to 255]
**Min** | Pointer to **int32** | Number that specifies the minimum length of a token. This value must be less than or equal to **max**. | [optional] [default to 0]
**Matches** | **string** | Value that indicates whether to replace only the first matching pattern or all matching patterns. | 
**Pattern** | **string** | Regular expression pattern to apply to each token. | 
**Replacement** | **string** | Replacement string to substitute wherever a matching pattern occurs. | 
**MaxShingleSize** | **int32** | Value that specifies the maximum number of tokens per shingle. This value must be greater than or equal to **minShingleSize**. | 
**MinShingleSize** | **int32** | Value that specifies the minimum number of tokens per shingle. This value must be less than or equal to **maxShingleSize**. | 
**StemmerName** | **string** | Snowball-generated stemmer to use. | 
**IgnoreCase** | Pointer to **bool** | Flag that indicates whether to ignore the case of stop words when filtering the tokens to remove. | [optional] [default to true]
**Tokens** | **[]string** | The stop words that correspond to the tokens to remove. Value must be one or more stop words. | 

## Methods

### NewApiAtlasFTSAnalyzersViewManualTokenFiltersInner

`func NewApiAtlasFTSAnalyzersViewManualTokenFiltersInner(type_ string, maxGram int32, minGram int32, matches string, pattern string, replacement string, maxShingleSize int32, minShingleSize int32, stemmerName string, tokens []string, ) *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner`

NewApiAtlasFTSAnalyzersViewManualTokenFiltersInner instantiates a new ApiAtlasFTSAnalyzersViewManualTokenFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSAnalyzersViewManualTokenFiltersInnerWithDefaults

`func NewApiAtlasFTSAnalyzersViewManualTokenFiltersInnerWithDefaults() *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner`

NewApiAtlasFTSAnalyzersViewManualTokenFiltersInnerWithDefaults instantiates a new ApiAtlasFTSAnalyzersViewManualTokenFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalTokens

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetOriginalTokens() string`

GetOriginalTokens returns the OriginalTokens field if non-nil, zero value otherwise.

### GetOriginalTokensOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetOriginalTokensOk() (*string, bool)`

GetOriginalTokensOk returns a tuple with the OriginalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTokens

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetOriginalTokens(v string)`

SetOriginalTokens sets OriginalTokens field to given value.

### HasOriginalTokens

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) HasOriginalTokens() bool`

HasOriginalTokens returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetType(v string)`

SetType sets Type field to given value.


### GetMaxGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMaxGram() int32`

GetMaxGram returns the MaxGram field if non-nil, zero value otherwise.

### GetMaxGramOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMaxGramOk() (*int32, bool)`

GetMaxGramOk returns a tuple with the MaxGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMaxGram(v int32)`

SetMaxGram sets MaxGram field to given value.


### GetMinGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMinGram() int32`

GetMinGram returns the MinGram field if non-nil, zero value otherwise.

### GetMinGramOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMinGramOk() (*int32, bool)`

GetMinGramOk returns a tuple with the MinGram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGram

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMinGram(v int32)`

SetMinGram sets MinGram field to given value.


### GetTermNotInBounds

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetTermNotInBounds() string`

GetTermNotInBounds returns the TermNotInBounds field if non-nil, zero value otherwise.

### GetTermNotInBoundsOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetTermNotInBoundsOk() (*string, bool)`

GetTermNotInBoundsOk returns a tuple with the TermNotInBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermNotInBounds

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetTermNotInBounds(v string)`

SetTermNotInBounds sets TermNotInBounds field to given value.

### HasTermNotInBounds

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) HasTermNotInBounds() bool`

HasTermNotInBounds returns a boolean if a field has been set.

### GetNormalizationForm

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetNormalizationForm() string`

GetNormalizationForm returns the NormalizationForm field if non-nil, zero value otherwise.

### GetNormalizationFormOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetNormalizationFormOk() (*string, bool)`

GetNormalizationFormOk returns a tuple with the NormalizationForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationForm

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetNormalizationForm(v string)`

SetNormalizationForm sets NormalizationForm field to given value.

### HasNormalizationForm

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) HasNormalizationForm() bool`

HasNormalizationForm returns a boolean if a field has been set.

### GetMax

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMatches

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMatches() string`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMatchesOk() (*string, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMatches(v string)`

SetMatches sets Matches field to given value.


### GetPattern

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetReplacement

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.


### GetMaxShingleSize

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMaxShingleSize() int32`

GetMaxShingleSize returns the MaxShingleSize field if non-nil, zero value otherwise.

### GetMaxShingleSizeOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMaxShingleSizeOk() (*int32, bool)`

GetMaxShingleSizeOk returns a tuple with the MaxShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxShingleSize

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMaxShingleSize(v int32)`

SetMaxShingleSize sets MaxShingleSize field to given value.


### GetMinShingleSize

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMinShingleSize() int32`

GetMinShingleSize returns the MinShingleSize field if non-nil, zero value otherwise.

### GetMinShingleSizeOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetMinShingleSizeOk() (*int32, bool)`

GetMinShingleSizeOk returns a tuple with the MinShingleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinShingleSize

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetMinShingleSize(v int32)`

SetMinShingleSize sets MinShingleSize field to given value.


### GetStemmerName

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetStemmerName() string`

GetStemmerName returns the StemmerName field if non-nil, zero value otherwise.

### GetStemmerNameOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetStemmerNameOk() (*string, bool)`

GetStemmerNameOk returns a tuple with the StemmerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStemmerName

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetStemmerName(v string)`

SetStemmerName sets StemmerName field to given value.


### GetIgnoreCase

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetTokens

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ApiAtlasFTSAnalyzersViewManualTokenFiltersInner) SetTokens(v []string)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


