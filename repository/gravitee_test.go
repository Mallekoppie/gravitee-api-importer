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
