package greetings

import(
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T){
	name := "Adamu"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Adamu")

	if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloEmpty(t *testing.T){
	name := ""
	msg, err := Hello(name)

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}