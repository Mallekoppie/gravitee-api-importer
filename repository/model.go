package repository

type ImportApiRequest struct {
	Format            string   `json:"format"`
	Payload           string   `json:"payload"`
	Type              string   `json:"type"`
	WithDocumentation bool     `json:"with_documentation"`
	WithPathMapping   bool     `json:"with_path_mapping"`
	WithPolicies      []string `json:"with_policies"`
	WithPolicyPaths   bool     `json:"with_policy_paths"`
}

type ImportAPIResponse struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	Visibility  string        `json:"visibility"`
	State       string        `json:"state"`
	Tags        []interface{} `json:"tags"`
	ContextPath string        `json:"context_path"`
	Proxy       struct {
		VirtualHosts []struct {
			Path string `json:"path"`
		} `json:"virtual_hosts"`
		StripContextPath bool `json:"strip_context_path"`
		PreserveHost     bool `json:"preserve_host"`
		Groups           []struct {
			Name      string `json:"name"`
			Endpoints []struct {
				Name   string `json:"name"`
				Target string `json:"target"`
				Weight int    `json:"weight"`
				Backup bool   `json:"backup"`
				Type   string `json:"type"`
				HTTP   struct {
					ConnectTimeout           int  `json:"connectTimeout"`
					IdleTimeout              int  `json:"idleTimeout"`
					KeepAlive                bool `json:"keepAlive"`
					ReadTimeout              int  `json:"readTimeout"`
					Pipelining               bool `json:"pipelining"`
					MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
					UseCompression           bool `json:"useCompression"`
					FollowRedirects          bool `json:"followRedirects"`
				} `json:"http"`
			} `json:"endpoints"`
			LoadBalancing struct {
				Type string `json:"type"`
			} `json:"load_balancing"`
			HTTP struct {
				ConnectTimeout           int  `json:"connectTimeout"`
				IdleTimeout              int  `json:"idleTimeout"`
				KeepAlive                bool `json:"keepAlive"`
				ReadTimeout              int  `json:"readTimeout"`
				Pipelining               bool `json:"pipelining"`
				MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
				UseCompression           bool `json:"useCompression"`
				FollowRedirects          bool `json:"followRedirects"`
			} `json:"http"`
		} `json:"groups"`
	} `json:"proxy"`
	FlowMode string `json:"flow_mode"`
	Flows    []struct {
		Name         string `json:"name"`
		PathOperator struct {
			Operator string `json:"operator"`
			Path     string `json:"path"`
		} `json:"path-operator"`
		Condition string        `json:"condition"`
		Methods   []string      `json:"methods"`
		Pre       []interface{} `json:"pre"`
		Post      []interface{} `json:"post"`
		Enabled   bool          `json:"enabled"`
	} `json:"flows"`
	Plans     []interface{} `json:"plans"`
	Gravitee  string        `json:"gravitee"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
	Owner     struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Type        string `json:"type"`
	} `json:"owner"`
	Properties []interface{} `json:"properties"`
	Services   struct {
	} `json:"services"`
	Resources         []interface{} `json:"resources"`
	PathMappings      []string      `json:"path_mappings"`
	ResponseTemplates struct {
	} `json:"response_templates"`
	LifecycleState                 string `json:"lifecycle_state"`
	DisableMembershipNotifications bool   `json:"disable_membership_notifications"`
}

type GetAPIResponse struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	Visibility  string        `json:"visibility"`
	State       string        `json:"state"`
	Tags        []interface{} `json:"tags"`
	Labels      []interface{} `json:"labels"`
	Entrypoints []struct {
		Target string `json:"target"`
	} `json:"entrypoints"`
	ContextPath string `json:"context_path"`
	Proxy       struct {
		VirtualHosts []struct {
			Path string `json:"path"`
		} `json:"virtual_hosts"`
		StripContextPath bool `json:"strip_context_path"`
		PreserveHost     bool `json:"preserve_host"`
		Groups           []struct {
			Name      string `json:"name"`
			Endpoints []struct {
				Name   string `json:"name"`
				Target string `json:"target"`
				Weight int    `json:"weight"`
				Backup bool   `json:"backup"`
				Type   string `json:"type"`
				HTTP   struct {
					ConnectTimeout           int  `json:"connectTimeout"`
					IdleTimeout              int  `json:"idleTimeout"`
					KeepAlive                bool `json:"keepAlive"`
					ReadTimeout              int  `json:"readTimeout"`
					Pipelining               bool `json:"pipelining"`
					MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
					UseCompression           bool `json:"useCompression"`
					FollowRedirects          bool `json:"followRedirects"`
				} `json:"http"`
			} `json:"endpoints"`
			LoadBalancing struct {
				Type string `json:"type"`
			} `json:"load_balancing"`
			HTTP struct {
				ConnectTimeout           int  `json:"connectTimeout"`
				IdleTimeout              int  `json:"idleTimeout"`
				KeepAlive                bool `json:"keepAlive"`
				ReadTimeout              int  `json:"readTimeout"`
				Pipelining               bool `json:"pipelining"`
				MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
				UseCompression           bool `json:"useCompression"`
				FollowRedirects          bool `json:"followRedirects"`
			} `json:"http"`
		} `json:"groups"`
	} `json:"proxy"`
	FlowMode string `json:"flow_mode"`
	Flows    []struct {
		Name         string `json:"name"`
		PathOperator struct {
			Operator string `json:"operator"`
			Path     string `json:"path"`
		} `json:"path-operator"`
		Condition string        `json:"condition"`
		Methods   []string      `json:"methods"`
		Pre       []interface{} `json:"pre"`
		Post      []interface{} `json:"post"`
		Enabled   bool          `json:"enabled"`
	} `json:"flows"`
	Plans     []interface{} `json:"plans"`
	Gravitee  string        `json:"gravitee"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
	Owner     struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Type        string `json:"type"`
	} `json:"owner"`
	Properties []interface{} `json:"properties"`
	Services   struct {
	} `json:"services"`
	PictureURL        string        `json:"picture_url"`
	Resources         []interface{} `json:"resources"`
	PathMappings      []string      `json:"path_mappings"`
	ResponseTemplates struct {
	} `json:"response_templates"`
	LifecycleState                 string `json:"lifecycle_state"`
	DisableMembershipNotifications bool   `json:"disable_membership_notifications"`
	BackgroundURL                  string `json:"background_url"`
}

type UpdateDeploymentPathRequest struct {
	Name        string        `json:"name"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	Visibility  string        `json:"visibility"`
	Tags        []interface{} `json:"tags"`
	Labels      []interface{} `json:"labels"`
	Proxy       struct {
		VirtualHosts []struct {
			Path string `json:"path"`
		} `json:"virtual_hosts"`
		StripContextPath bool `json:"strip_context_path"`
		PreserveHost     bool `json:"preserve_host"`
		Groups           []struct {
			Name      string `json:"name"`
			Endpoints []struct {
				Name   string `json:"name"`
				Target string `json:"target"`
				Weight int    `json:"weight"`
				Backup bool   `json:"backup"`
				Type   string `json:"type"`
				HTTP   struct {
					ConnectTimeout           int  `json:"connectTimeout"`
					IdleTimeout              int  `json:"idleTimeout"`
					KeepAlive                bool `json:"keepAlive"`
					ReadTimeout              int  `json:"readTimeout"`
					Pipelining               bool `json:"pipelining"`
					MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
					UseCompression           bool `json:"useCompression"`
					FollowRedirects          bool `json:"followRedirects"`
				} `json:"http"`
			} `json:"endpoints"`
			LoadBalancing struct {
				Type string `json:"type"`
			} `json:"load_balancing"`
			HTTP struct {
				ConnectTimeout           int  `json:"connectTimeout"`
				IdleTimeout              int  `json:"idleTimeout"`
				KeepAlive                bool `json:"keepAlive"`
				ReadTimeout              int  `json:"readTimeout"`
				Pipelining               bool `json:"pipelining"`
				MaxConcurrentConnections int  `json:"maxConcurrentConnections"`
				UseCompression           bool `json:"useCompression"`
				FollowRedirects          bool `json:"followRedirects"`
			} `json:"http"`
		} `json:"groups"`
	} `json:"proxy"`
	FlowMode string `json:"flow_mode"`
	Flows    []struct {
		Name         string `json:"name"`
		PathOperator struct {
			Operator string `json:"operator"`
			Path     string `json:"path"`
		} `json:"path-operator"`
		Condition string        `json:"condition"`
		Methods   []string      `json:"methods"`
		Pre       []interface{} `json:"pre"`
		Post      []interface{} `json:"post"`
		Enabled   bool          `json:"enabled"`
	} `json:"flows"`
	Plans      []interface{} `json:"plans"`
	Properties []interface{} `json:"properties"`
	Services   struct {
	} `json:"services"`
	PictureURL        string        `json:"picture_url"`
	Resources         []interface{} `json:"resources"`
	PathMappings      []string      `json:"path_mappings"`
	ResponseTemplates struct {
	} `json:"response_templates"`
	LifecycleState                 string `json:"lifecycle_state"`
	DisableMembershipNotifications bool   `json:"disable_membership_notifications"`
}

func (in *GetAPIResponse) MapToUpdateDeploymentPathRequest() UpdateDeploymentPathRequest {
	out := UpdateDeploymentPathRequest{
		Name:                           in.Name,
		Version:                        in.Version,
		Description:                    in.Description,
		Visibility:                     in.Visibility,
		Tags:                           in.Tags,
		Labels:                         in.Labels,
		Proxy:                          in.Proxy,
		FlowMode:                       in.FlowMode,
		Flows:                          in.Flows,
		Plans:                          in.Plans,
		Properties:                     in.Properties,
		Services:                       in.Services,
		PictureURL:                     in.PictureURL,
		Resources:                      in.Resources,
		PathMappings:                   in.PathMappings,
		ResponseTemplates:              in.ResponseTemplates,
		LifecycleState:                 in.LifecycleState,
		DisableMembershipNotifications: in.DisableMembershipNotifications,
	}

	return out
}
