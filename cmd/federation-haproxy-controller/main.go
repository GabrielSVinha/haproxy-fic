package main

import(
  glog "github.com/golang/glog"
  controller "./pkg/controller"
  kubernetes "k8s.io/client-go/kubernetes"
  clientcmd "k8s.io/client-go/tools/clientcmd"
  parser "./pkg/parser"
)

func main(){
  cfg := parser.ParseParameters()
  // create config based on flags (next: annotations)
  kubernetesClient := CreateAPIClient(cfg.FedAPIHost, conf.KubeConfig)
}

func CreateAPIClient (fedApiServerHost string, kubeConfig string) (*kubernetes.Clientset, error) {
  conf := clientcmd.BuildConfigFromFlags("", kubeConfig)
  cli, err := kubernetes.NewForConfig(conf)
  if err != nil {
		return nil, err
	}
  return client, nil
}

