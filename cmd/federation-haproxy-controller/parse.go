package main

import(
  rest "k8s.io/client-go/rest"
  clientcmd "k8s.io/client-go/tools/clientcmd"
  "os"
  "fmt"
)

func ParseParameters () (*rest.Config, error){
  fmt.Println("Time to get things done ", os.Getenv("KUBECONFIG"))

  //config := util.NewAdminConfig(clientcmd.NewDefaultPathOptions())
  //fmt.Println("Vamo ve se o admin config ta ", &config)
  //return config.FederationClientset("federation-controller-manager@kfed", os.Getenv("KUBECONFIG"))
  config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
  fmt.Println("Checkpoint de erro 0")
  return config, err
}

