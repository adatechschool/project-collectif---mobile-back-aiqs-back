package handler

import (
	"encoding/json"
	"net/http"

	"wave/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllSurfSpots(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	surfSpots := []model.Surf{}
	db.Find(&surfSpots)
	respondJSON(w, http.StatusOK, surfSpots)
}

func CreateSurfSpots(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	surfSpot := model.Surf{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&surfSpot); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&surfSpot).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, surfSpot)
}

func GetSurfSpot(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["ID"]
	surfSpot := getSurfSpotOr404(db, id, w, r)
	if surfSpot == nil {
		return
	}
	respondJSON(w, http.StatusOK, surfSpot)
}

func UpdateSurfSpot(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["ID"]
	surfSpot := getSurfSpotOr404(db, id, w, r)
	if surfSpot == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&surfSpot); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&surfSpot).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, surfSpot)
}

func DeleteSurfSpot(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["ID"]
	surfSpot := getSurfSpotOr404(db, id, w, r)
	if surfSpot == nil {
		return
	}
	if err := db.Delete(&surfSpot).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	employee.Enable()
// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getSurfSpotOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Surf {
	surf := model.Surf{}
	if err := db.First(&surf, model.Surf{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &surf
}
