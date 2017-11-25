package parser

import(
  "k8s.io/client-go/tools/clientcmd"
  "os"
  "k8s.io/federation/client/clientset_generated/federation_clientset"
)

func ParseParameters () (federationclientset.Interface){
  config := util.NewAdminConfig(clientcmd.NewDefaultPathOptions())
  return config.FederationClientset("federation-controller-manager@kfed", os.Getenv("KUBECONFIG"))
}

