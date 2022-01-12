package main

func Repeat(character string, repeatTime int) string {
	var repeated string 
	if repeatTime < 1 {
		return character
	}

	for i := 0 ; i < repeatTime; i++ {
		repeated += character
	}

	return repeated
}