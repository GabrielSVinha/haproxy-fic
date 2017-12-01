package main

import(
  "os"
//  "io/ioutil"
  "bytes"
  "strconv"
  corev1 "k8s.io/api/core/v1"
)

const (
  fileName = "haproxy.cfg"
  tmpl = "haproxy.tmpl"
)

func WriteBackEnds(ips[]corev1.LoadBalancerIngress) (string, error){
  template, err := ioutil.ReadFile(tmpl)

  if err != nil {
    return "", err
  }

/*  cfg, err := os.Create(fileName)

  if err != nil{
    return "", err
  }

  defer cfg.Close()
*/
  var backendBuff bytes.Buffer

  for index, lb := range ips {
    backendBuff.WriteString("  server cluster")
    backendBuff.WriteString(strconv.Itoa(index))
    backendBuff.WriteString(" ")
    backendBuff.WriteString(lb.IP)
    backendBuff.WriteString(":80 check\n")

    _, err := template.WriteString(backendBuff.String())
    if err != nil{
      return "", err
    }
  }
/*
  buffer,err := ioutil.ReadFile(fileName)

  if err != nil {
    return "", err
  }
*/
  return backendBuff.String(), nil
}

