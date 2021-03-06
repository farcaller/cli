package controller

import (
	"time"

	tcs "github.com/k8sdb/apimachinery/client/clientset"
	clientset "k8s.io/client-go/kubernetes"
)

type Controller struct {
	// Kubernetes client
	Client clientset.Interface
	// ThirdPartyExtension client
	ExtClient tcs.ExtensionInterface
}

const (
	DatabaseNamePrefix = "kubedb"
	LabelDatabaseKind  = "kubedb.com/kind"
	LabelDatabaseName  = "kubedb.com/name"
	sleepDuration      = time.Second * 10
)
