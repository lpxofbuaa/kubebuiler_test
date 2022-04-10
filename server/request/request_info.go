package request

type LpxpodRequest struct {
	Name           string `json:"name,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	Podport        int32  `json:"podport,omitempty"`
	Deploymentname string `json:"deploymentname,omitempty"`
	Schedule       string `json:"schedule,omitempty"`
}
