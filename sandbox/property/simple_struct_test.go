package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)


func TestUserProfile(t *testing.T) {
	t.Run("Generate valid user profiles", func(t *testing.T){
		exampleCount := 10
		for i := range exampleCount {
			randomProfile := rapid.Make[*UserProfile]()
			fmt.Printf("random profile: %+v", randomProfile.Example(i))
		}
	})
}
// Traditional example-based tests
func TestUserProfileValidation(t *testing.T) {
	testCases := []struct {
		name string
		user UserProfile 
		expectError bool
		errorContains string
	}{
		{
			name: "Empty profile",
		 	user: UserProfile{},
			expectError: true,
			errorContains: "username must contain",
		},
		{
			name: "Valid profile",
		 	user: UserProfile{
				ID: 1, 
				Name: Name{FirstName: "John", LastName: "Smith"},
				UserName: "johnsmith",
				Password: "test123",
			},
			expectError: false,
			errorContains: "",
		},
		{
			name: "Valid profile",
		 	user: UserProfile{
				ID: 1, 
				Name: Name{FirstName: "John", LastName: "Smith"},
				UserName: "john.smith",
				Password: "test123",
			},
			expectError: false,
			errorContains: "",
		},
		{
			name: "Invalid username",
		 	user: UserProfile{
				ID: 1, 
				Name: Name{FirstName: "John", LastName: "Smith"},
				UserName: "john-smith",
				Password: "test123",
			},
			expectError: true,
			errorContains: "username must contain",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			err := tc.user.Validate()

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorContains != "" {
					assert.Contains(t, err.Error(), tc.errorContains)
				}
			} else {
				assert.NoError(t, err)
			}

		})
	}
}