package Http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"study/internal/core"
	"study/internal/repository"
	"study/pkg/postgres/models"
	"time"

	"github.com/gorilla/mux"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func WriteError(w http.ResponseWriter, err error, status int) {
	errDTO := ErrorDTO{Message: err.Error(), Time: time.Now()}
	http.Error(w, errDTO.ToString(), status)
}

type HTTPHandlers struct {
	Rep    *repository.Repository
	Logger *core.Logger
}

func NewHandlers(r *repository.Repository, logger *core.Logger) *HTTPHandlers {
	return &HTTPHandlers{Rep: r, Logger: logger}
}

// CREATE
/*
pattern: /tasks
method: POST
info: JSON

succeed:
-status code: 201
-response body: JSON

failed:
-status code: 400, 409, 500
-response body: JSON : error+time
*/
func (h *HTTPHandlers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var Create CreateDTO
	err := json.NewDecoder(r.Body).Decode(&Create)
	if err != nil {
		h.Logger.Error("error on decodong body: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	var CreateModel models.CreateModel
	CreateModel.Title = Create.Title
	CreateModel.Description = Create.Description

	err = h.Rep.InsertRow(CreateModel)
	if err != nil {
		h.Logger.Error("error on: inserting row", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(Create, "", "	")
	if err != nil {
		h.Logger.Error("error on creating new json: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		h.Logger.Error("error on writing response: ", core.Field{Key: "err: ", Value: err})
		fmt.Println("failed to write response: ", err.Error())
	}
	h.Logger.Debug("succesfully adding new row:", core.Field{Key: "CreateModel", Value: CreateModel})
}

// READ
/*
pattern: /tasks
method: GET
info: -

succeed:
-status code: 200
-response body: JSON slice

failed:
-status code: 500
-response body: JSON : error + time
*/

func (h *HTTPHandlers) ReadHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Rep.Read()
	if err != nil {
		h.Logger.Error("error on reading tasks: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	var Read ReadDTO
	Read.Tasks = tasks

	b, err := json.MarshalIndent(Read, "", "	")
	if err != nil {
		h.Logger.Error("error on writing json: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		h.Logger.Error("error on writing response: ", core.Field{Key: "err: ", Value: err})
		fmt.Println("failed to write response: ", err.Error())
	}
	h.Logger.Debug("succesfully reading tasks: ", core.Field{Key: "Tasks: ", Value: Read.Tasks})
}

// UPDATE
/*
pattern: /tasks/{id}
method: PATCH
info: pattern

succeed:
-status code: 200
-response body: JSON

failed:
-status code: 400, 404, 500
-response body: JSON: error + time
*/
func (h *HTTPHandlers) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("error on converting string into int: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	err = h.Rep.Update(id)
	if err != nil {
		h.Logger.Error("error on updating row: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	newRow := h.Rep.ReadOne(id)
	b, err := json.MarshalIndent(newRow, "", "	")
	if err != nil {
		h.Logger.Error("error on reading one row: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		h.Logger.Error("error on writing response: ", core.Field{Key: "err: ", Value: err})
		fmt.Println("failed to write response: ", err)
		return
	}
	h.Logger.Debug("succesfully updated row: ", core.Field{Key: "New: ", Value: newRow})
}

// DELETE
/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:
-status code: 204
-response body: -

failed:
-status code: 400, 404, 500
-response body: JSON: error + time
*/

func (h *HTTPHandlers) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("error on converting string into int: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	Info := h.Rep.ReadOne(id)

	err = h.Rep.Delete(id)
	if err != nil {
		h.Logger.Error("error on deleting row: ", core.Field{Key: "err: ", Value: err})
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.Logger.Debug("successfully deleted task: ", core.Field{Key: "Info: ", Value: Info})
}

/*
HEALTH
pattern: /health
method: GET
info: -

succeed:
status: 200
info: "ok"

*/

func (h *HTTPHandlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		h.Logger.Error("error on health: ", core.Field{Key: "err: ", Value: err})
		fmt.Println("error on writing ok in health: ", err)
		return
	}
}
