package controllers

import (
	"encoding/json"
	"net/http"
	"policy-search-api/services"

	"github.com/gorilla/mux"
)

// RegisterPolicyRoutes sets up the API routes for policy operations
func RegisterPolicyRoutes(r *mux.Router, policyService *services.PolicyService) {
	r.HandleFunc("/api/v1/policies/search", func(w http.ResponseWriter, r *http.Request) {
		searchPolicies(w, r, policyService)
	}).Methods("GET")
}

// searchPolicies handles the search request
// @Summary Search policies
// @Description Search policies based on various criteria
// @Tags policies
// @Accept  json
// @Produce  json
// @Param type query string false "Policy Type"
// @Param company query string false "Company Name"
// @Param id query string false "Policy Id"
// @Param name query string false "Policy Name"
// @Param years query string false "Number of Years"
// @Success 200 {array} domain.Policy
// @Failure 404 {string} string "No policies found"
// @Router /api/v1/policies/search [get]
func searchPolicies(w http.ResponseWriter, r *http.Request, policyService *services.PolicyService) {
	criteria := map[string]string{
		"type":    r.URL.Query().Get("type"),
		"company": r.URL.Query().Get("company"),
		"id":      r.URL.Query().Get("id"),
		"name":    r.URL.Query().Get("name"),
		"years":   r.URL.Query().Get("years"),
	}

	policies, err := policyService.SearchPolicies(criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if policies == nil {
		http.Error(w, "No policies found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(policies)
}
