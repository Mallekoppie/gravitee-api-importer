package repository

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mallekoppie/goslow/platform"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient       *http.Client
	DefaultTLSConfig = &tls.Config{InsecureSkipVerify: true}

	ErrGraviteeIncorrectResponseCode = errors.New("Incorrect response code")
	ErrGraviteeAPIPathAlreadyInUse   = errors.New("API Path already used by another API")
	ErrGraviteeProxyPathInvalid      = errors.New("Proxy path invalid")

	//	TODO: Must be moved to config
	graviteeManagementAPIHost string = "http://localhost:8083"
	graviteeUsername          string = "admin"
	graviteePassword          string = "admin"
	graviteeOrganizationId    string = "DEFAULT"
	graviteeEnvironmentId     string = "DEFAULT"
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 30
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
			TLSClientConfig:     DefaultTLSConfig,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func ImportAPIContract(contract string) (importApiResponse ImportAPIResponse, err error) {
	requestBody := ImportApiRequest{
		Format:            "API",
		Payload:           contract,
		Type:              "INLINE",
		WithDocumentation: true,
		WithPathMapping:   true,
		WithPolicyPaths:   false,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		platform.Logger.Error("Marshalling request for contract import", zap.Error(err))
		return importApiResponse, err
	}
	bodyData := bytes.NewBuffer(data)

	serverUrl := fmt.Sprintf("%s/management/organizations/%s/environments/%s/apis/import/swagger?definitionVersion=2.0.0", graviteeManagementAPIHost, graviteeOrganizationId, graviteeEnvironmentId)

	req, err := http.NewRequest(http.MethodPost, serverUrl, bodyData)

	req.SetBasicAuth(graviteeUsername, graviteePassword)
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("accept", "application/json")

	response, err := httpClient.Do(req)
	defer response.Body.Close()
	if err != nil {
		platform.Logger.Error("Calling Gravitee Management API", zap.Error(err))
		return importApiResponse, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		platform.Logger.Error("Reading create API response body", zap.Error(err))
		return importApiResponse, err
	}

	platform.Logger.Info("Import API Response Body", zap.String("response_body", string(responseData)))

	if response.StatusCode != http.StatusCreated && response.StatusCode == http.StatusBadRequest {
		responseString := string(responseData)
		responseString = strings.ToLower(responseString)

		if strings.Contains(responseString, "is already covered by an other api") {
			platform.Logger.Error("API Import. API path already in use", zap.String("response_body", responseString))
			return importApiResponse, ErrGraviteeAPIPathAlreadyInUse
		} else {
			platform.Logger.Error("API Import. Unknown reason for 400 response", zap.String("response_body", responseString))
			return importApiResponse, err
		}
	} else if response.StatusCode != http.StatusCreated {
		platform.Logger.Error("Incorrect response code for API Import", zap.Int("status_code_expected", http.StatusCreated), zap.Int("status_code_actual", response.StatusCode))

		return importApiResponse, ErrGraviteeIncorrectResponseCode
	}

	platform.Logger.Info("API Imported successfully")

	importApiResponse = ImportAPIResponse{}

	err = json.Unmarshal(responseData, &importApiResponse)
	if err != nil {
		platform.Logger.Error("Unmarshalling API response", zap.Error(err))

		return importApiResponse, err
	}

	return importApiResponse, nil
}

func GetAPI(apiId string) (api GetAPIResponse, err error) {
	serverUrl := fmt.Sprintf("%s/management/organizations/%s/environments/%s/apis/%s", graviteeManagementAPIHost, graviteeOrganizationId, graviteeEnvironmentId, apiId)

	request, err := http.NewRequest(http.MethodGet, serverUrl, nil)
	if err != nil {
		platform.Logger.Info("Creating GetAPI Request", zap.Error(err))
		return
	}
	request.SetBasicAuth(graviteeUsername, graviteePassword)

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		platform.Logger.Error("Calling Get API", zap.Error(err))
		return
	}

	if response.StatusCode != http.StatusOK {
		platform.Logger.Error("Get API returned incorrect response code", zap.Int("status_code_expected", http.StatusOK),
			zap.Int("status_code_actual", response.StatusCode))
		return
	}
	api = GetAPIResponse{}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		platform.Logger.Error("Reading Get API Response", zap.Error(err))
		return
	}

	api = GetAPIResponse{}

	err = json.Unmarshal(responseData, &api)
	if err != nil {
		platform.Logger.Error("Unmarshalling Get API Response", zap.Error(err))
		return
	}

	return api, nil
}

func UpdateProxyPath(apiId string, proxyPath string) (err error) {
	startLetter := proxyPath[:1]
	if startLetter != "/" {
		platform.Logger.Error("Invalid Proxy path for update", zap.String("api_id", apiId), zap.String("proxy_path", proxyPath))

	}

	api, err := GetAPI(apiId)
	if err != nil {
		platform.Logger.Error("Retrieving API before updating the proxy path")
		return err
	}

	updateRequest := api.MapToUpdateDeploymentPathRequest()
	updateRequest.Proxy.VirtualHosts[0].Path = proxyPath
	updateRequestData, err := json.Marshal(updateRequest)
	if err != nil {
		platform.Logger.Error("Error marshalling update deployment request", zap.Error(err))
		return
	}

	buffer := bytes.NewBuffer(updateRequestData)

	serverUrl := fmt.Sprintf("%s/management/organizations/%s/environments/%s/apis/%s", graviteeManagementAPIHost, graviteeOrganizationId, graviteeEnvironmentId, apiId)

	request, err := http.NewRequest(http.MethodPut, serverUrl, buffer)
	if err != nil {
		platform.Logger.Error("Creating new update path request", zap.Error(err))
		return
	}

	request.SetBasicAuth(graviteeUsername, graviteePassword)
	request.Header.Add("content-type", "application/json;charset=UTF-8")
	request.Header.Add("accept", "application/json")

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		platform.Logger.Error("Error calling API to update path", zap.String("api_id", apiId), zap.Error(err))
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		platform.Logger.Error("Update path read response body", zap.Error(err))
		return
	}

	if response.StatusCode != http.StatusOK {
		platform.Logger.Error("Update Path incorrect response", zap.Int("status_code_expected", http.StatusOK), zap.Int("status_code_actual", response.StatusCode))
		return ErrGraviteeIncorrectResponseCode
	}

	platform.Logger.Info("Path update successfull", zap.String("api_id", apiId), zap.String("proxy_path", proxyPath), zap.String("response_body", string(responseData)))

	return
}

func PublishApi(apiId string) (err error) {
	api, err := GetAPI(apiId)
	if err != nil {
		platform.Logger.Error("Retrieving API before Publishing API")
		return err
	}

	updateRequest := api.MapToUpdateDeploymentPathRequest()
	updateRequest.LifecycleState = "published"
	updateRequestData, err := json.Marshal(updateRequest)
	if err != nil {
		platform.Logger.Error("Error marshalling Publish API request", zap.Error(err))
		return
	}

	buffer := bytes.NewBuffer(updateRequestData)

	serverUrl := fmt.Sprintf("%s/management/organizations/%s/environments/%s/apis/%s", graviteeManagementAPIHost, graviteeOrganizationId, graviteeEnvironmentId, apiId)

	request, err := http.NewRequest(http.MethodPut, serverUrl, buffer)
	if err != nil {
		platform.Logger.Error("Creating new Publish API request", zap.Error(err))
		return
	}

	request.SetBasicAuth(graviteeUsername, graviteePassword)
	request.Header.Add("content-type", "application/json;charset=UTF-8")
	request.Header.Add("accept", "application/json")

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		platform.Logger.Error("Error calling API to Publish API", zap.String("api_id", apiId), zap.Error(err))
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		platform.Logger.Error("Publish API read response body", zap.Error(err))
		return
	}

	if response.StatusCode != http.StatusOK {
		platform.Logger.Error("Publish API incorrect response", zap.Int("status_code_expected", http.StatusOK), zap.Int("status_code_actual", response.StatusCode))
		return ErrGraviteeIncorrectResponseCode
	}

	platform.Logger.Info("Publish successfull", zap.String("api_id", apiId), zap.String("response_body", string(responseData)))

	return
}

func PublishPage(apiId string) (err error) {
	serverUrl := fmt.Sprintf("%s/management/organizations/%s/environments/%s/apis/%s/pages/?root=true", graviteeManagementAPIHost, graviteeOrganizationId, graviteeEnvironmentId, apiId)

	request, err := http.NewRequest(http.MethodGet, serverUrl, nil)
	if err != nil {
		platform.Logger.Info("Creating GetAPI Request", zap.Error(err))
		return
	}
	request.SetBasicAuth(graviteeUsername, graviteePassword)
	request.Header.Add("accept", "application/json")

	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		platform.Logger.Error("Calling Get API", zap.Error(err))
		return
	}

	if response.StatusCode != http.StatusOK {
		platform.Logger.Error("Get API returned incorrect response code", zap.Int("status_code_expected", http.StatusOK),
			zap.Int("status_code_actual", response.StatusCode))
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		platform.Logger.Error("Reading Get API Response", zap.Error(err))
		return
	}

	api := GetPagesResponse{}
	err = json.Unmarshal(responseData, &api)
	if err != nil {
		platform.Logger.Error("Unmarshalling Get API Response", zap.Error(err))
		return
	}

	var pageId string
	for _, v := range api {
		if v.Type == "SWAGGER" {
			pageId = v.ID
		}
	}

	updateRequestBody := PublishPageRequest{Published: true}

	updateRequestData, err := json.Marshal(updateRequestBody)
	if err != nil {
		platform.Logger.Error("Marchalling for page publish", zap.Error(err))
		return
	}

	buffer := bytes.NewBuffer(updateRequestData)

	pageUpdateUrl := fmt.Sprintf("%s/management/organizations/%s/environments/%s/apis/%s/pages/%s", graviteeManagementAPIHost, graviteeOrganizationId, graviteeEnvironmentId, apiId, pageId)

	updateRequest, err := http.NewRequest(http.MethodPatch, pageUpdateUrl, buffer)
	if err != nil {
		platform.Logger.Error("Creating new request for page update", zap.Error(err))
		return
	}

	updateRequest.SetBasicAuth(graviteeUsername, graviteePassword)
	updateRequest.Header.Add("content-type", "application/json;charset=UTF-8")
	updateRequest.Header.Add("accept", "application/json")

	updateResponse, err := httpClient.Do(updateRequest)
	if err != nil {
		platform.Logger.Error("Error calling publish page", zap.Error(err))
		return
	}

	if updateResponse.StatusCode != http.StatusOK {
		platform.Logger.Error("Incorrect response code for publish page")
		return
	}

	return
}
