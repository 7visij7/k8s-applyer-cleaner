package cleaner

import (
	"fmt"
	"time"
	"strings"
	"io/ioutil"
	"encoding/json"
	"smib-applyers-cleaner/types"
	"smib-applyers-cleaner/pkg/aes"
	"smib-applyers-cleaner/pkg/http"
	"smib-applyers-cleaner/pkg/config"
	"smib-applyers-cleaner/pkg/errors"
)

func ConvertDateToTimestump(date string) (int64) {
    timestamp, err := time.Parse(time.RFC3339,date);
	errors.CheckError(err)

	return timestamp.Unix()
}

func GetDeployments(namespace string, token string) (types.Deployments) {
	url := fmt.Sprintf("%s%s", types.OPENSHIFT_URL, fmt.Sprintf(types.DEPLOYMENT_URL, namespace))
	resp := http.HttpRequest("GET", namespace, token, url, "")

	if resp.StatusCode != 200 {
		fmt.Printf(" ERROR, can not get information about Deployments in Namespace - %s. Respond code -  %d.\n", namespace, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	errors.CheckError(err)

	var data types.Deployments
	_ = json.Unmarshal([]byte(body), &data)

	return data
}

func ReduceReplica(name string, namespace string, token string)  {
	url := fmt.Sprintf("%s%s/%s/scale", types.OPENSHIFT_URL, fmt.Sprintf(types.DEPLOYMENT_URL, namespace),  name)
	payload := fmt.Sprintf(`{"kind": "Scale", "apiVersion": "autoscaling/v1", "metadata": {"name": "%s", "namespace": "%s"},"spec": { "replicas": 0}}`, name, namespace)
	resp := http.HttpRequest("PUT", namespace, token, url, payload)
	
	if resp.StatusCode == 200 {
		fmt.Printf("Successfully scaled replica to 0 for Deployment - %s in Namespace - %s.\n", name, namespace)
		return
	}
	fmt.Printf(" ERROR, can not scale replica to 0 for Deployment - %s in Namespace - %s. Respond code -  %d.\n", name, namespace, resp.StatusCode)
}

func FilterDeloyment(deployments types.Deployments, namespace string, token string) {
	for _, j := range deployments.Items {
		if IsApplyer(j.Metadata.Name) && j.Status.Replicas > 0 {
			if IsOldApplyer(ConvertDateToTimestump(j.Status.Conditions[1].LastUpdateTime)) {
				ReduceReplica(j.Metadata.Name, namespace, token)
			}
		}
	}
}

func IsApplyer(name string) (bool){
	if name == "smib-marketplace-applyer"{ return false }
	return strings.Contains(name, "applyer")
}

func IsOldApplyer(timestamp int64) (bool) {
	if (types.CURRENT_TIMESTUMP - 60*30) > timestamp {
		return true
	}
	return false
}

func CheckApplyer() {
	creds := config.GetCreds()
	for namespace, token := range creds {
		FilterDeloyment(GetDeployments(namespace, aes.Decrypt(token)), namespace, aes.Decrypt(token))
	}
}