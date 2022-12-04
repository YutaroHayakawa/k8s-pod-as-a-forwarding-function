package main

import (
	"context"
	"fmt"
	"log"
	"os"

	ciliumv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/cilium/cilium/pkg/node/addressing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

const (
	groupName    = "cilium.io"
	groupVersion = "v2"
)

var (
	schemaGroupVersion = schema.GroupVersion{
		Group:   groupName,
		Version: groupVersion,
	}
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	addToScheme   = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		schemaGroupVersion,
		&ciliumv2.CiliumNode{},
		&ciliumv2.CiliumNodeList{},
	)

	metav1.AddToGroupVersion(scheme, schemaGroupVersion)
	return nil
}

func main() {
	var err error
	var config *rest.Config

	log.SetOutput(os.Stdout)

	nodeName := os.Getenv("NODE_NAME")
	if nodeName == "" {
		log.Fatal("NODE_NAME environment variable is not set")
	}

	config, err = rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to build client %s", err)
	}

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schemaGroupVersion
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.UnversionedRESTClientFor(&crdConfig)
	if err != nil {
		log.Fatalf("Failed to get REST client %s", err)
	}

	var node ciliumv2.CiliumNode
	err = client.Get().Resource("ciliumnodes").Name(nodeName).Do(context.TODO()).Into(&node)
	if err != nil {
		log.Fatalf("Failed to cilium nodes %s", err)
	}

	for _, ip := range node.Spec.Addresses {
		if ip.Type == addressing.NodeCiliumInternalIP {
			fmt.Println(ip.IP)
			return
		}
	}

	log.Fatalf("Failed to find Cilium node address")
}
