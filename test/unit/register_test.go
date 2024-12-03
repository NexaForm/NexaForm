package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func TestRegisterUser(t *testing.T) {
	url := "https://localhost:8080/register"

	t.Run("successful register ", func(t *testing.T) {
		newMockUser := struct {
			FirstName string `json:"firstname"`
			LastName  string `json:"lastname"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}{
			FirstName: "amir",
			LastName:  "amiri",
			Email:     "amirreza.jabbari2023@gmail.com",
			Password:  "123werhd",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("failed to marshal requests: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("failed to make post request: %v", err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response: %v", err)
		}

		res := new(Response)
		err = json.Unmarshal(body, res)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}
		assert.Equal(t, http.StatusCreated, resp.StatusCode, "status code should be 201")
		assert.True(t, res.Success, "success field should be true")
		assert.Equal(t, "user successfully registered", res.Message, "message should be 'user successfully registered'")
		assert.Nil(t, res.Data, "data field should be nil for this test case")
		assert.Empty(t, res.Error, "error field should be empty")
	})

	t.Run("email not provided", func(t *testing.T) {
		newMockUser := MockUser{
			FirstName: "john",
			LastName:  "johnny",
			Password:  "12@Amir###90",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}

		res := new(Response)
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "status code should be %s", fiber.StatusBadRequest)
		assert.Contains(t, res.Error, "Key: 'UserRegisterReq.Email' Error:Field validation for 'Email' failed on the 'required' tag", "error message should contain 'Key: 'UserRegisterReq.Email' Error:Field validation for 'Email' failed on the 'required' tag'")
	})

	t.Run("password not provided", func(t *testing.T) {
		newMockUser := MockUser{
			FirstName: "john",
			LastName:  "johnny",
			Email:     "john@gmail.com",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}

		res := new(Response)
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "status code should be %s", fiber.StatusBadRequest)
		assert.Contains(t, res.Error, "Key: 'UserRegisterReq.Password' Error:Field validation for 'Password' failed on the 'required' tag", "error message should contain 'Key: 'UserRegisterReq.Password' Error:Field validation for 'Password' failed on the 'required' tag'")
	})

	t.Run("invalid email format", func(t *testing.T) {
		newMockUser := MockUser{
			FirstName: "john",
			LastName:  "johnny",
			Email:     "john@invalid",
			Password:  "12@Amir###90",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}

		res := new(Response)
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "status code should be %s", fiber.StatusBadRequest)
		assert.Contains(t, res.Error, "invalid email format")
	})

	t.Run("email already registered", func(t *testing.T) {
		// Assuming the email is already registered in the system
		newMockUser := MockUser{
			FirstName: "john",
			LastName:  "johnny",
			Email:     "john@gmail.com",
			Password:  "12@Amir###90",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}

		res := new(Response)
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, fiber.StatusConflict, resp.StatusCode, "status code should be %s", fiber.StatusConflict)
		assert.Contains(t, res.Error, "email already exists", "error message should contain 'email already exists'")
	})

	t.Run("missing first name", func(t *testing.T) {
		newMockUser := MockUser{
			LastName: "johnny",
			Email:    "john@gmail.com",
			Password: "12@Amir###90",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}
		res := new(Response)
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "status code should be %s", fiber.StatusBadRequest)
		assert.Contains(t, res.Error, "Key: 'UserRegisterReq.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag", "error message should contain 'Key: 'UserRegisterReq.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag'")
	})
	t.Run("missing last name", func(t *testing.T) {
		newMockUser := MockUser{
			FirstName: "john",
			Email:     "john@gmail.com",
			Password:  "12@Amir###90",
		}

		reqBody, err := json.Marshal(newMockUser)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response: %v", err)
		}

		res := new(Response)
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "status code should be %s", fiber.StatusBadRequest)
		assert.Contains(t, res.Error, "Key: 'UserRegisterReq.LastName' Error:Field validation for 'LastName' failed on the 'required' tag", "error message should contain 'Key: 'UserRegisterReq.LastName' Error:Field validation for 'LastName' failed on the 'required' tag'")
	})
}
