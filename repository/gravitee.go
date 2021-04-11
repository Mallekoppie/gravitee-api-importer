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

	//	TODO: Must be moved to config
	graviteeManagementAPIHost string = "http://localhost:8083"
	graviteeUsername          string = "admin"
	graviteePassword          string = "admin"
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

	serverUrl := fmt.Sprintf("%s/management/organizations/DEFAULT/environments/DEFAULT/apis/import/swagger?definitionVersion=2.0.0", graviteeManagementAPIHost)

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
