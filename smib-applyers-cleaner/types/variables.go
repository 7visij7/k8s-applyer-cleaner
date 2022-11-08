package types

import (
	"time"
	"os"
)

var (
	// PROJECTS = map[string]string { "ds1-genr01-smib-dso-smbi" : "TOKEN",
	// 			 "ds1-genr01-smib-dso-smbn" : "TOKEN",
	// 			 "ds1-genr01-smib-dso-smbu" : "TOKEN",
	// 			 "ds1-genr01-smib-dso-smib" : "TOKEN",
	// 			 "ds1-genr01-smib-st-smbi" : "TOKEN",
	// 			 "ds1-genr01-smib-st-smbn" : "TOKEN",
	// 			 "ds1-genr01-smib-st-smbu" : "TOKEN",
	// 			 "ds1-genr01-smib-st-smib" : "TOKEN" }
	OPENSHIFT_URL = os.Getenv("OPENSHIFT_SERVER")
	DEPLOYMENT_URL= "/apis/apps/v1/namespaces/%s/deployments"
	CURRENT_TIMESTUMP = time.Now().Unix()
	ENCRYPT_KEY = os.Getenv("ENCRYPT_KEY")
)

type  Projects struct {
	Project map[string]string
}

type Scale struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata MetadataScale `json:"metadata"`
	Spec SpecScale `json:"spec"`
}

type MetadataScale struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

type SpecScale struct {
	Replicas int `json:"replicas"`
}

type Deployments struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Items []Item `json:"items"`
}

type Item struct{
	Metadata Metadata `json:"metadata"`
	Status Status `json:"status"`
} 

type Metadata struct{
	Name string `json:"name"`
	Namespace string `json:"namespace"`
} 

type Status struct{
	Replicas int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	Conditions []Condition `json:"conditions"`
} 

type Condition struct{
	Type string `json:"type"`
	Status string `json:"status"`
	LastUpdateTime string `json:"lastUpdateTime"`
} 