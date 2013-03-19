package radio

import (
	"testing"
	"fmt"
)

func TestRadioAlpha( t *testing.T ){
	
	fmt.Println("TestRadioAlpha", "---------")
	
	fmt.Println("Spool Alpabet>", GetAjaxAlphabet() )
	
	fmt.Println("Test Callsigns", "-------------" )
	callsigns := []string{"unix09", "win98", "go380", "test123"}
	for idx, callsign := range callsigns{
		fmt.Println(">", idx, callsign, Callsign2Words(callsign))
	}
	
	
}
