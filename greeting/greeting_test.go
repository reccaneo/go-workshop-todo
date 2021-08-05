package greeting

import (
	"testing"
)

//	TODO: Test Driven Development
func TestGreetingYourName(t *testing.T) {
	//	AAA Pattern

	//	Arrange
	given := "Bob"
	want := "Hello ,Bob."

	//	Act
	get := Greet(given)

	//	Assert
	if want != get {
		t.Errorf("given a name %s want greeting %q, but got %q", given, want, get)
	}
}

func TestGreetingMyFriend(t *testing.T) {
	//	AAA Pattern

	//	Arrange
	given := ""
	want := "Hello ,my friend."

	//	Act
	get := Greet(given)

	//	Assert
	if want != get {
		t.Errorf("given a name %s want greeting %q, but got %q", given, want, get)
	}
}

func TestGreetingCapital(t *testing.T) {
	//	AAA Pattern

	//	Arrange
	given := "BOB"
	want := "HELLO ,BOB."

	//	Act
	get := Greet(given)

	//	Assert
	if want != get {
		t.Errorf("given a name %s want greeting %q, but got %q", given, want, get)
	}
}
