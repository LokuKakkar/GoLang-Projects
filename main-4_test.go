package main

// Run using go test
import "testing"

func TestGetUser(t *testing.T) {
	md := MockDataStore{
		Users: map[int]User{
			2: {ID: 2, FirstName: "Johnn"},
		},
	}

	s := &Service{ds: md}

	u, err := s.GetUser(2)
	if err != nil {
		t.Errorf("Error getting user %v", err)
	}

	if u.FirstName != "Johnn" {
		t.Errorf("got: %v, expected: %v", u.FirstName, "Johnn")
	}
}
