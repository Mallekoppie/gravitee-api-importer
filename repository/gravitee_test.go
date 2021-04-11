package repository

import (
	"log"
	"testing"
)

func TestApiImport(t *testing.T) {
	response, err := ImportAPIContract(apiToImport)

	if err != nil {
		log.Println("API import failed: ", err.Error())
		t.Fail()
	}

	log.Println("Imported GUID: ", response.ID)
}

func TestApiImportPetstore(t *testing.T) {
	response, err := ImportAPIContract(apiToImportPetstoreExpanded)

	if err != nil {
		log.Println("API import failed: ", err.Error())
		t.Fail()
	}

	log.Println("Imported GUID: ", response.ID)
}

func TestApiImportUspTo(t *testing.T) {
	response, err := ImportAPIContract(apiToImportUspTo)

	if err != nil {
		log.Println("API import failed: ", err.Error())
		t.Fail()
	}

	log.Println("Imported GUID: ", response.ID)
}

func TestGetAPI(t *testing.T) {
	api, err := GetAPI("4bd314bf-1ebd-4ac0-9314-bf1ebd9ac09b")
	if err != nil {
		log.Println("Error getting API", err.Error())
		t.Fail()
	}

	log.Println("Returned API: ", api.ID)
}

func TestUpdatePath(t *testing.T) {
	err := UpdateProxyPath("1ef3455f-9ae2-4d8d-b345-5f9ae29d8da4", "/updatedfromunittest")
	if err != nil {
		log.Println("Error updating path", err.Error())
		t.Fail()
	}
}

func TestPublishApi(t *testing.T) {
	err := PublishApi("6a2b2a4b-2f60-44ab-ab2a-4b2f6014abed")
	if err != nil {
		log.Println("Error Publishing API", err.Error())
		t.Fail()
	}
}

func TestPublishApiPage(t *testing.T) {
	err := PublishPage("6a2b2a4b-2f60-44ab-ab2a-4b2f6014abed")
	if err != nil {
		log.Println("Error publishing page: ", err.Error())
		t.Fail()
	}
}

var apiToImport = `---
swagger: "2.0"
info:
  description: "Used to deply API's to API GW"
  version: "1.0.0"
  title: "Deployment API"
  contact: {}
consumes:
- "application/json"
produces:
- "application/json"
paths:
  /deployments/{environment}/apis:
    get:
      summary: "Get all Apis"
      description: "returns all the API's deployed in that environment"
      operationId: "getAllApis"
      parameters:
      - name: "environment"
        in: "path"
        required: true
        type: "string"
        default: "DEV"
        enum:
        - "DEV"
        - "INT"
        - "QA"
        - "PRD"
        - "DR"
        x-exportParamName: "Environment"
      - name: "apiName"
        in: "query"
        description: "search for API's"
        required: false
        type: "string"
        x-example: "accounts"
        x-exportParamName: "ApiName"
        x-optionalDataType: "String"
      responses:
        200:
          description: "Status 200"
    post:
      summary: "deployContract"
      description: "Deploys a contract"
      operationId: "deployApi"
      consumes: []
      parameters:
      - name: "environment"
        in: "path"
        required: true
        type: "string"
        default: "DEV"
        enum:
        - "DEV"
        - "INT"
        - "QA"
        - "PRD"
        - "DR"
        x-exportParamName: "Environment"
      - in: "body"
        name: "body"
        required: true
        schema:
          type: "object"
        x-exportParamName: "Body"
      responses:
        201:
          description: "Status 201"
        412:
          description: "Status 412"
  /deployments/{environment}/apis/{apiId}:
    delete:
      summary: "Undeply an API"
      description: "Removes plan, stops the API and deletes it"
      parameters:
      - name: "environment"
        in: "path"
        required: true
        type: "string"
        x-exportParamName: "Environment"
      - name: "apiId"
        in: "path"
        required: true
        type: "string"
        x-exportParamName: "ApiId"
      responses:
        204:
          description: "Delete successfull"
        412:
          description: "Delete failed"
  /deployments/{environment}/apis/{apiId}/tags:
    put:
      summary: "tags an API"
      description: "Adds the tag to the existing tags"
      operationId: "tagApi"
      consumes: []
      parameters:
      - name: "environment"
        in: "path"
        required: true
        type: "string"
        x-exportParamName: "Environment"
      - name: "apiId"
        in: "path"
        required: true
        type: "string"
        x-exportParamName: "ApiId"
      - in: "body"
        name: "body"
        required: true
        schema:
          type: "array"
          items:
            $ref: "#/definitions/Tag"
        x-exportParamName: "Body"
      responses:
        202:
          description: "Tagging completed"
        409:
          description: "Tag does not exist"
  /promotions/{environment}:
    post:
      summary: "Promotes an API to the next environment"
      consumes: []
      parameters:
      - name: "environment"
        in: "path"
        required: true
        type: "string"
        default: "INT"
        enum:
        - "INT"
        - "QA"
        - "PRD"
        - "DR"
        x-exportParamName: "Environment"
      - in: "body"
        name: "body"
        required: true
        schema:
          $ref: "#/definitions/Promotion"
        x-exportParamName: "Body"
      responses:
        201:
          description: "Promotion successfull"
        412:
          description: "Promotion Failed"
definitions:
  Deployment:
    type: "object"
    required:
    - "contract"
    properties:
      contract:
        type: "string"
        description: "the contract in json format"
      ImportParams:
        $ref: "#/definitions/ImportParameters"
    description: "describes a deployment"
  ImportParameters:
    type: "object"
    required:
    - "ApprovedRequestHeaders"
    - "ApprovedResponseCodeRewrite"
    - "ApprovedResponseCodes"
    - "ApprovedResponseHeaders"
    - "BackendEndpoint"
    - "ConnectTimeout"
    - "Description"
    - "MaxConcurrentConnections"
    - "Product"
    - "ReadTimeout"
    - "RequestValidationEnabled"
    - "Service"
    - "UseCompression"
    - "Version"
    properties:
      Product:
        type: "string"
        minLength: 3
      Service:
        type: "string"
        minLength: 3
      Version:
        type: "number"
        minimum: 1
      BackendEndpoint:
        type: "string"
        minLength: 10
      Description:
        type: "string"
      ApprovedRequestHeaders:
        type: "string"
        description: "Authorization"
      ApprovedResponseHeaders:
        type: "string"
        description: "Authorization"
      ApprovedResponseCodes:
        type: "string"
        description: "200,201,400"
      ContractType:
        type: "string"
        enum:
        - "REST"
        - "SOAP"
      ApprovedResponseCodeRewrite:
        type: "number"
        description: "205"
      ConnectTimeout:
        type: "number"
        example: 10000
        description: "10000"
      ReadTimeout:
        type: "number"
        example: 30000
        minimum: 1
      MaxConcurrentConnections:
        type: "number"
        minimum: 10
      UseCompression:
        type: "boolean"
        default: true
      RequestValidationEnabled:
        type: "boolean"
        default: true
    description: "Describes how the contract will be imported"
  Tag:
    type: "string"
    description: "used to tag an API\n\ntest - The test Gateway\ngw-int - To internal\
      \ segments\ngw-ext - To the public internet\ngw-pub - Out to the public internet\n\
      gw-third - Out to third party"
    enum:
    - "test"
    - "gw-int"
    - "gw-ext"
    - "gw-third"
    - "gw-pub"
    default: "test"
  Promotion:
    type: "object"
    required:
    - "Product"
    - "Service"
    - "Tags"
    - "Version"
    properties:
      Product:
        type: "string"
      Service:
        type: "string"
      Version:
        type: "number"
      Tags:
        type: "array"
        items:
          $ref: "#/definitions/Tag"
    description: "parameters for a promotion"
`

var apiToImportPetstoreExpanded = `openapi: "3.0.0"
info:
  title: Simple API overview
  version: 2.0.0
paths:
  /:
    get:
      operationId: listVersionsv2
      summary: List API versions
      responses:
        '200':
          description: |-
            200 response
          content:
            application/json:
              examples: 
                foo:
                  value:
                    {
                      "versions": [
                        {
                            "status": "CURRENT",
                            "updated": "2011-01-21T11:33:21Z",
                            "id": "v2.0",
                            "links": [
                                {
                                    "href": "http://127.0.0.1:8774/v2/",
                                    "rel": "self"
                                }
                            ]
                        },
                        {
                            "status": "EXPERIMENTAL",
                            "updated": "2013-07-23T11:33:21Z",
                            "id": "v3.0",
                            "links": [
                                {
                                    "href": "http://127.0.0.1:8774/v3/",
                                    "rel": "self"
                                }
                            ]
                        }
                      ]
                    }
        '300':
          description: |-
            300 response
          content:
            application/json: 
              examples: 
                foo:
                  value: |
                   {
                    "versions": [
                          {
                            "status": "CURRENT",
                            "updated": "2011-01-21T11:33:21Z",
                            "id": "v2.0",
                            "links": [
                                {
                                    "href": "http://127.0.0.1:8774/v2/",
                                    "rel": "self"
                                }
                            ]
                        },
                        {
                            "status": "EXPERIMENTAL",
                            "updated": "2013-07-23T11:33:21Z",
                            "id": "v3.0",
                            "links": [
                                {
                                    "href": "http://127.0.0.1:8774/v3/",
                                    "rel": "self"
                                }
                            ]
                        }
                    ]
                   }
  /v2:
    get:
      operationId: getVersionDetailsv2
      summary: Show API version details
      responses:
        '200':
          description: |-
            200 response
          content:
            application/json: 
              examples:
                foo:
                  value:
                    {
                      "version": {
                        "status": "CURRENT",
                        "updated": "2011-01-21T11:33:21Z",
                        "media-types": [
                          {
                              "base": "application/xml",
                              "type": "application/vnd.openstack.compute+xml;version=2"
                          },
                          {
                              "base": "application/json",
                              "type": "application/vnd.openstack.compute+json;version=2"
                          }
                        ],
                        "id": "v2.0",
                        "links": [
                          {
                              "href": "http://127.0.0.1:8774/v2/",
                              "rel": "self"
                          },
                          {
                              "href": "http://docs.openstack.org/api/openstack-compute/2/os-compute-devguide-2.pdf",
                              "type": "application/pdf",
                              "rel": "describedby"
                          },
                          {
                              "href": "http://docs.openstack.org/api/openstack-compute/2/wadl/os-compute-2.wadl",
                              "type": "application/vnd.sun.wadl+xml",
                              "rel": "describedby"
                          },
                          {
                            "href": "http://docs.openstack.org/api/openstack-compute/2/wadl/os-compute-2.wadl",
                            "type": "application/vnd.sun.wadl+xml",
                            "rel": "describedby"
                          }
                        ]
                      }
                    }
        '203':
          description: |-
            203 response
          content:
            application/json: 
              examples:
                foo:
                  value:
                    {
                      "version": {
                        "status": "CURRENT",
                        "updated": "2011-01-21T11:33:21Z",
                        "media-types": [
                          {
                              "base": "application/xml",
                              "type": "application/vnd.openstack.compute+xml;version=2"
                          },
                          {
                              "base": "application/json",
                              "type": "application/vnd.openstack.compute+json;version=2"
                          }
                        ],
                        "id": "v2.0",
                        "links": [
                          {
                              "href": "http://23.253.228.211:8774/v2/",
                              "rel": "self"
                          },
                          {
                              "href": "http://docs.openstack.org/api/openstack-compute/2/os-compute-devguide-2.pdf",
                              "type": "application/pdf",
                              "rel": "describedby"
                          },
                          {
                              "href": "http://docs.openstack.org/api/openstack-compute/2/wadl/os-compute-2.wadl",
                              "type": "application/vnd.sun.wadl+xml",
                              "rel": "describedby"
                          }
                        ]
                      }
                    }`

var apiToImportUspTo = `openapi: 3.0.1
servers:
  - url: '{scheme}://developer.uspto.gov/ds-api'
    variables:
      scheme:
        description: 'The Data Set API is accessible via https and http'
        enum:
          - 'https'
          - 'http'
        default: 'https'
info:
  description: >-
    The Data Set API (DSAPI) allows the public users to discover and search
    USPTO exported data sets. This is a generic API that allows USPTO users to
    make any CSV based data files searchable through API. With the help of GET
    call, it returns the list of data fields that are searchable. With the help
    of POST call, data can be fetched based on the filters on the field names.
    Please note that POST call is used to search the actual data. The reason for
    the POST call is that it allows users to specify any complex search criteria
    without worry about the GET size limitations as well as encoding of the
    input parameters.
  version: 1.0.0
  title: USPTO Data Set API
  contact:
    name: Open Data Portal
    url: 'https://developer.uspto.gov'
    email: developer@uspto.gov
tags:
  - name: metadata
    description: Find out about the data sets
  - name: search
    description: Search a data set
paths:
  /:
    get:
      tags:
        - metadata
      operationId: list-data-sets
      summary: List available data sets
      responses:
        '200':
          description: Returns a list of data sets
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/dataSetList'
              example:
                {
                  "total": 2,
                  "apis": [
                    {
                      "apiKey": "oa_citations",
                      "apiVersionNumber": "v1",
                      "apiUrl": "https://developer.uspto.gov/ds-api/oa_citations/v1/fields",
                      "apiDocumentationUrl": "https://developer.uspto.gov/ds-api-docs/index.html?url=https://developer.uspto.gov/ds-api/swagger/docs/oa_citations.json"
                    },
                    {
                      "apiKey": "cancer_moonshot",
                      "apiVersionNumber": "v1",
                      "apiUrl": "https://developer.uspto.gov/ds-api/cancer_moonshot/v1/fields",
                      "apiDocumentationUrl": "https://developer.uspto.gov/ds-api-docs/index.html?url=https://developer.uspto.gov/ds-api/swagger/docs/cancer_moonshot.json"
                    }
                  ]
                }
  /{dataset}/{version}/fields:
    get:
      tags:
        - metadata
      summary: >-
        Provides the general information about the API and the list of fields
        that can be used to query the dataset.
      description: >-
        This GET API returns the list of all the searchable field names that are
        in the oa_citations. Please see the 'fields' attribute which returns an
        array of field names. Each field or a combination of fields can be
        searched using the syntax options shown below.
      operationId: list-searchable-fields
      parameters:
        - name: dataset
          in: path
          description: 'Name of the dataset.'
          required: true
          example: "oa_citations"
          schema:
            type: string
        - name: version
          in: path
          description: Version of the dataset.
          required: true
          example: "v1"
          schema:
            type: string
      responses:
        '200':
          description: >-
            The dataset API for the given version is found and it is accessible
            to consume.
          content:
            application/json:
              schema:
                type: string
        '404':
          description: >-
            The combination of dataset name and version is not found in the
            system or it is not published yet to be consumed by public.
          content:
            application/json:
              schema:
                type: string
  /{dataset}/{version}/records:
    post:
      tags:
        - search
      summary: >-
        Provides search capability for the data set with the given search
        criteria.
      description: >-
        This API is based on Solr/Lucene Search. The data is indexed using
        SOLR. This GET API returns the list of all the searchable field names
        that are in the Solr Index. Please see the 'fields' attribute which
        returns an array of field names. Each field or a combination of fields
        can be searched using the Solr/Lucene Syntax. Please refer
        https://lucene.apache.org/core/3_6_2/queryparsersyntax.html#Overview for
        the query syntax. List of field names that are searchable can be
        determined using above GET api.
      operationId: perform-search
      parameters:
        - name: version
          in: path
          description: Version of the dataset.
          required: true
          schema:
            type: string
            default: v1
        - name: dataset
          in: path
          description: 'Name of the dataset. In this case, the default value is oa_citations'
          required: true
          schema:
            type: string
            default: oa_citations
      responses:
        '200':
          description: successful operation
          content:
            application/json:
              schema:
                type: array
                items:
                  type: object
                  additionalProperties:
                    type: object
        '404':
          description: No matching record found for the given criteria.
      requestBody:
        content:
          application/x-www-form-urlencoded:
            schema:
              type: object
              properties:
                criteria:
                  description: >-
                    Uses Lucene Query Syntax in the format of
                    propertyName:value, propertyName:[num1 TO num2] and date
                    range format: propertyName:[yyyyMMdd TO yyyyMMdd]. In the
                    response please see the 'docs' element which has the list of
                    record objects. Each record structure would consist of all
                    the fields and their corresponding values.
                  type: string
                  default: '*:*'
                start:
                  description: Starting record number. Default value is 0.
                  type: integer
                  default: 0
                rows:
                  description: >-
                    Specify number of rows to be returned. If you run the search
                    with default values, in the response you will see 'numFound'
                    attribute which will tell the number of records available in
                    the dataset.
                  type: integer
                  default: 100
              required:
                - criteria
components:
  schemas:
    dataSetList:
      type: object
      properties:
        total:
          type: integer
        apis:
          type: array
          items:
            type: object
            properties:
              apiKey:
                type: string
                description: To be used as a dataset parameter value
              apiVersionNumber:
                type: string
                description: To be used as a version parameter value
              apiUrl:
                type: string
                format: uriref
                description: "The URL describing the dataset's fields"
              apiDocumentationUrl:
                type: string
                format: uriref
                description: A URL to the API console for each API`
