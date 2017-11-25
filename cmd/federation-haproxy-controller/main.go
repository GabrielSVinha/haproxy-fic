package main

import(
  glog "github.com/golang/glog"
  controller "./pkg/controller"
  clientcmd "k8s.io/client-go/tools/clientcmd"
  parser "./pkg/parser"
)

func main(){
  fedClient, err := parser.ParseParameters()
  if err != nil{
    glog.Error("Failed to parse configuration: ", err)
  }
  //todo
  //hapController := controller.NewHAProxyFedIngressController(fedClient)
}

func CreateAPIClient (fedApiServerHost string, kubeConfig string) (*kubernetes.Clientset, error) {
  conf := clientcmd.BuildConfigFromFlags("", kubeConfig)
  cli, err := kubernetes.NewForConfig(conf)
  if err != nil {
		return nil, err
	}
  return client, nil
}

