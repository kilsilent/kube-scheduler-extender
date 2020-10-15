package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"io"
	"kube-scheduler-extender/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
	schedulerapi "k8s.io/kube-scheduler/extender/v1"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()
	Router.GET("/", Index)
	Router.GET("/healthcheck", HealthCheck)
	Router.POST("/filter", Filter)
	Router.POST("/prioritize", Prioritize)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to kube-scheduler-extender!\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OK\n")

}

func Filter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var buf bytes.Buffer
	body := io.TeeReader(r.Body, &buf)
	var extenderArgs schedulerapi.ExtenderArgs
	var extenderFilterResult *schedulerapi.ExtenderFilterResult
	if err := json.NewDecoder(body).Decode(&extenderArgs); err != nil {
		log.Errorln("解析参数错误:", err)
		extenderFilterResult = &schedulerapi.ExtenderFilterResult{
			Error: err.Error(),
		}
	} else {
		//b, _ := json.Marshal(extenderArgs)
		//fmt.Println(string(b))
		extenderFilterResult = controller.Filter(extenderArgs)
	}

	if response, err := json.Marshal(extenderFilterResult); err != nil {
		log.Errorln("json 格式化 extenderFilterResult:", err)
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func Prioritize(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var buf bytes.Buffer
	body := io.TeeReader(r.Body, &buf)
	var extenderArgs schedulerapi.ExtenderArgs
	var hostPriorityList *schedulerapi.HostPriorityList
	if err := json.NewDecoder(body).Decode(&extenderArgs); err != nil {
		log.Errorln("解析参数错误:", err)
		hostPriorityList = &schedulerapi.HostPriorityList{}
	} else {
		hostPriorityList = controller.Prioritize(extenderArgs)

	}

	if response, err := json.Marshal(hostPriorityList); err != nil {
		log.Errorln("json 格式化 hostPriorityList:", err)
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// func Bind(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
