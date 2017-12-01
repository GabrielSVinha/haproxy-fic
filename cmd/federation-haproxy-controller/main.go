package main

import(
  //controller "k8s.io/federation/cmd/federation-haproxy-controller/app/controller"
  "fmt"
  kubernetes "k8s.io/client-go/kubernetes"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

func main(){
  config, err := ParseParameters()

  if err != nil{
    fmt.Println("Failed to parse configuration: ", err)
  }

  fedClientSet, err := kubernetes.NewForConfig(config)

  if err != nil{
    panic(err.Error())
  }

  deploys, err := fedClientSet.CoreV1().Services("demo").List(metav1.ListOptions{})// TO DO: list services attached to ingress res
  for index, svc := range deploys.Items{
    _, err := WriteBackEnds(svc.Status.LoadBalancer.Ingress)
    if err != nil {
      fmt.Println("Erro no index: ", index)
      panic(err.Error())
    }
  }
  for{
    fmt.Println("Do something")
  }
}

