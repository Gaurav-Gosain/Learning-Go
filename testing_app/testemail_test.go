package testing_app

import("testing")

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("Hello")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	_, err = IsEmail("itsgauravgosain@gmail.com")
	if err != nil {
		t.Error("itsgauravgosain@gmail.com is an email")
	}
}