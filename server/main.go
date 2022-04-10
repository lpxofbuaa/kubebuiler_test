package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	crdtryVersioned "kubebuilder.test/crdtry/generated/crdtry/clientset/versioned"
	request "kubebuilder.test/crdtry/server/request"
)

var handler *request.Handler

func CreateLpxPod(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	lpxpodReq := request.LpxpodRequest{}
	err = json.Unmarshal(body, &lpxpodReq)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	msg, err := handler.HandleCreateLpxpodReq(&lpxpodReq)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	rmsg, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(200)
	w.Write(rmsg)
}

func ListLpxPod(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
}

func GetLpxPod(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	lpxpodReq := request.LpxpodRequest{}
	err = json.Unmarshal(body, &lpxpodReq)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	msg, err := handler.HandleGetLpxpodReq(&lpxpodReq)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	rmsg, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(200)
	w.Write(rmsg)
}

func DeleteLpxPod(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		return
	}
}

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := crdtryVersioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	handler = &request.Handler{Clientset: clientset}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/lpxpod/create", CreateLpxPod)
	mux.HandleFunc("/api/lpxpod/list", ListLpxPod)
	mux.HandleFunc("/api/lpxpod/get", GetLpxPod)
	mux.HandleFunc("/api/lpxpod/delete", DeleteLpxPod)
	http.ListenAndServe(":8002", mux)
}
