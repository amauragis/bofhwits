package bofhwitsbot

import "testing"

func Test_FormatTextOneLine(t *testing.T) {
	input1 := `
20:23 <Malkar> though it is really quiet in here
20:23 <Malkar> what the hell are you people doing on a friday night
20:23 <Malkar> you /should/ be keeping me entertained while I work`

	output1 := formatTextOneLine(input1)
	expOutput1 := `20:23 <Malkar> though it is really quiet in here | 20:23 <Malkar> what the hell are you people doing on a friday night | 20:23 <Malkar> you /should/ be keeping me entertained while I work`

	if output1 != expOutput1 {
		t.Errorf("One-line format error!\nExpected:\n%s\nGot: \n%s\n", expOutput1, output1)
	}
}

func Test_FormatTextMultiLine(t *testing.T) {
	t.Skip("Test not implemented.")
	// input1 :=
	// 	`20:23 <Malkar> though it is really quiet in here
	// 20:23 <Malkar> what the hell are you people doing on a friday night
	// 20:23 <Malkar> you /should/ be keeping me entertained while I work`

}
