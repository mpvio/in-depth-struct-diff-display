package main

import (
	"fmt"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/r3labs/diff/v3"
)

func DiffStrings(a, b string) string {
	aRunes := strings.Split(a, "")
	bRunes := strings.Split(b, "")

	matcher := difflib.NewMatcher(aRunes, bRunes)
	opcodes := matcher.GetOpCodes()

	var buf strings.Builder
	for _, op := range opcodes {
		switch op.Tag {
		case 'e': // equal
			buf.WriteString(strings.Join(aRunes[op.I1:op.I2], ""))
		case 'r': // replace
			buf.WriteString("<")
			buf.WriteString(strings.Join(aRunes[op.I1:op.I2], ""))
			buf.WriteString("→")
			buf.WriteString(strings.Join(bRunes[op.J1:op.J2], ""))
			buf.WriteString(">")
		case 'd': // delete
			buf.WriteString("-<")
			buf.WriteString(strings.Join(aRunes[op.I1:op.I2], ""))
			buf.WriteString(">")
		case 'i': // insert
			buf.WriteString("+<")
			buf.WriteString(strings.Join(bRunes[op.J1:op.J2], ""))
			buf.WriteString(">")
		}
	}
	return buf.String()
}

func CompareStructs(a, b any) {
	changes, _ := diff.Diff(a, b)
	for _, c := range changes { // c = change
		fmt.Printf("%s: ", strings.Join(c.Path, "."))
		switch c.Type {
		case "update":
			if strOld, ok := c.From.(string); ok {
				strNew := c.To.(string)
				fmt.Printf("{%v}\n", DiffStrings(strOld, strNew))
			} else {
				fmt.Printf("{%v → %v}\n", c.From, c.To)
			}
		case "create":
			fmt.Printf("+{%v}\n", c.To)
		case "delete":
			fmt.Printf("-{%v}\n", c.From)
		}
	}
}
