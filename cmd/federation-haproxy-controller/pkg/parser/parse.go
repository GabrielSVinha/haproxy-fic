package parser

import(
  clientcmd "k8s.io/client-go/tools/clientcmd"
  "os"
  util "k8s.io/federation/pkg/kubefed/util"
)

func ParseParameters () (federationclientset.Interface){
  config := util.NewAdminConfig(clientcmd.NewDefaultPathOptions())
  return config.FederationClientset("federation-controller-manager@kfed", os.Getenv("KUBECONFIG"))
}

