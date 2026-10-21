package models

import "testing"

func TestUser_IsEligible(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		expected bool
	}{
		{
			name:     "Valid adult user",
			user:     User{Name: "Anushka", Email: "anushka@test.com", Age: 22},
			expected: true,
		},
		{
			name:     "Underage user",
			user:     User{Name: "Kid", Email: "kid@test.com", Age: 16},
			expected: false,
		},
		{
			name:     "Empty email",
			user:     User{Name: "NoEmail", Email: "", Age: 25},
			expected: false,
		},
		{
			name:     "Boundary: exactly 18",
			user:     User{Name: "Eighteen", Email: "eighteen@test.com", Age: 18},
			expected: true,
		},
		{
			name:     "Anushka 150 age",
			user:     User{Name: "OneFifty", Email: "bot@test.com", Age: 150},
			expected: false,
		},
		{
			name:     "User 1",
			user:     User{Name: "BadEmail", Email: "invalidemail.com", Age: 25},
			expected: false,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.user.IsEligible()
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
