package handler

import (
	"crud-toy/internal/app/model"
	"crud-toy/internal/app/service/execution"
	"encoding/json"
	"net/http"
	"time"
)

var (
	timeout time.Duration = 2 * time.Second
	prefix  string        = "key"
)

type Handler interface {
	CreateProc(w http.ResponseWriter, r *http.Request)
	ReadProcByID(w http.ResponseWriter, r *http.Request)
	ReadAllProc(w http.ResponseWriter, r *http.Request)
	UpdateProc(w http.ResponseWriter, r *http.Request)
	DeleteProc(w http.ResponseWriter, r *http.Request)
}

type ProcHandler struct {
	procExec execution.Execution
}

func NewProcHandler(procExecutor execution.Execution) Handler {
	return &ProcHandler{
		procExec: procExecutor,
	}
}

func (p *ProcHandler) CreateProc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	proc := model.Proc{}
	json.NewDecoder(r.Body).Decode(&proc)
	result, err := p.procExec.CreateProc(&proc)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(result)
}

func (p *ProcHandler) ReadProcByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var proc model.Proc
	json.NewDecoder(r.Body).Decode(&proc)
	result, err := p.procExec.ReadProcByID(&proc)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(result)
}

func (p *ProcHandler) ReadAllProc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	procs, err := p.procExec.ReadAllProc()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(procs)
}

func (p *ProcHandler) UpdateProc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var proc model.Proc
	json.NewDecoder(r.Body).Decode(&proc)
	result, err := p.procExec.UpdateProc(&proc)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&result)

}

func (p *ProcHandler) DeleteProc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var proc model.Proc
	json.NewDecoder(r.Body).Decode(&proc)
	err := p.procExec.DeleteProc(&proc)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode("value deleted")
}
