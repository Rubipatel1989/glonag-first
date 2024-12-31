package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"goLangFirst/internal/types"
	"goLangFirst/internal/utils/response"
	"io"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
}

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a student")
		var student types.Student
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid request body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		//request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"message": "student created successfully"})

	}
}
