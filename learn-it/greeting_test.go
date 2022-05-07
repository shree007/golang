package main

import "testing"

func TestGreeting( t *testing.T ){

	// test for empty argument
	emptyResult := greeting("")
    
	if emptyResult != "Hello Dude" {
		t.Error("greeting(\"\") failed" + " expected ","Hello Dude",  " and got ", emptyResult)
	} else{
		t.Logf("greeting(\"\") success, expected %v and got %v ", "Hello Dude", emptyResult)
	}

	//  test for valid argument
	result := greeting("Polo")

	if result != "Hello Polo!" {
		t.Error("greeting(\"\") failed" + " expected ","Hello Polo!",  " and got ", result)
	} else{
		t.Logf("greeting(\"\") success, expected %v and got %v ", "Hello Polo!",  result)
	}

}