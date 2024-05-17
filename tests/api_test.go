package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/woozie-10/api-clean-arch/domain"
)
const baseURL = "http://localhost:9090"
func TestApi(t *testing.T) {
	t.Run("AddStudent", func(t *testing.T) {
		student := domain.Student{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "username",
			Group:     "Group1",
			Course:    "Course1",
			Specialty: "Specialty1",
		}

		studentJSON, err := json.Marshal(student)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", baseURL+"/students", bytes.NewBuffer(studentJSON))
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Error testing Add handler: expected status code %v, got %v", http.StatusCreated, resp.StatusCode)
		}
	})
	
	t.Run("GetStudents", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/students", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing Get handler: expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("GetByUsername", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/students/username", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing GetByUsername handler: expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("GetByGroup", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/students/group/Group1", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing GetByGroup handler: expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("GetByCourse", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/students/course/Course1", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing GetByCourse handler: expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("UpdateStudent", func(t *testing.T) {
		student := domain.Student{
			FirstName: "Tyler",
			LastName:  "Doe",
			Username:  "username",
			Group:     "Group1",
			Course:    "Course1",
			Specialty: "Specialty1",
		}

		studentJSON, err := json.Marshal(student)
		if err != nil {
			t.Fatal(err)
		}
		
		req, err := http.NewRequest("PUT", baseURL+"/students/username", bytes.NewBuffer(studentJSON))
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing Update handler: expected status code %v, got %v", http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("DeleteStudent", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", baseURL+"/students/username", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing Delete handler: expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})

	t.Run("AddAssessment", func(t *testing.T) {
		assessment := domain.Assessment{
			Username: "username",
			Marks: map[string][]int{
				"math": {1, 2, 3, 4, 5},
			},
		}
		assessmentJSON, err := json.Marshal(assessment)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", baseURL+"/assessments", bytes.NewBuffer(assessmentJSON))
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated{
			t.Errorf("Error testing AddAssessment handler: expected status code %v, got %v", http.StatusCreated, resp.StatusCode)
		}
	})

	t.Run("GetAssessment", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/assessments/username", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp)
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error testing GetAssessment handler: expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})





}
