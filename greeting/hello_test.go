package greeting

import "testing"

/*func TestHelloEmptyArgs(t *testing.T) {
	// test for empty argument
	emptyArgResult := hello("")
	if emptyArgResult != "Hello Dude!" {
		t.Errorf("hello failed, expected %v, got %v", "Hello Dude!", emptyArgResult)
	} else {
		t.Logf("hello sucess, expected %v, got %v", "Hello Dude!", emptyArgResult)
	}
}
*/
func TestHelloValidArgs(t *testing.T) {
	// test for valid argument
	result := hello("Avadhut")
	if result != "Hello Avadhut!" {
		t.Errorf("hello(\"Avadhut\") failed, expected %v, got %v", "Hello Avadhut!", result)
	} else {
		t.Logf("hello(\"Avadhut\") sucess, expected %v, got %v", "Hello Dude!", result)
	}
}
