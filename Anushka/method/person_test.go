package method

import "testing"

func TestIsAdult(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{"Adult (age 23)", 23, true},
		{"Child (age 15)", 15, false},
		{"Boundary (age 18)", 18, true},
		{"Boundary (age 17)", 17, false},
		{"Baby (age 0)", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{Name: "TestUser", Age: tt.age}
			result := p.IsAdult()

			if result != tt.expected {
				t.Errorf("For age %d: expected %v, but got %v", tt.age, tt.expected, result)
			}
		})
	}
}

func TestHaveBirthday(t *testing.T){
	p := Person{Name: "Anushka", Age: 23}
	p.HaveBirthday()
	if p.Age != 24{
		t.Errorf("Expected age 24 but got %d",p.Age)
	}
}