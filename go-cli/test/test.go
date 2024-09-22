package main

import (
	"strings"
	"fmt"
)

type MapFormat map[string]string

func main() {
    my_in_f := "{date}.{time}_{string}"
    my_out_f := "{date}|{time}/{string}"
    my_string := "2024-12-31.12:34:56_foo_bar"
    var names, separators []string
    var f_names, f_separators []string
    formatInput(my_in_f, &names, &separators)
    formatInput(my_out_f, &f_names, &f_separators)
    LineMap := getInputNames(my_string, &names, &separators)
    f_line := formatLine(&LineMap, &f_names, &f_separators)
    fmt.Println(f_line)
}

func formatLine(terms *MapFormat, f_names *[]string, f_separators *[]string) string {
    var line string = ""
    sep_len := len(*f_separators)
    for i, term := range *f_names {
	line += (*terms)[term]
	if i < sep_len {
	    line += (*f_separators)[i]
	}

    }
    return line
}

func getInputNames(line string, names *[]string, separators *[]string) MapFormat {
    var words = make(MapFormat)
    var curr []string
    var remainder string = line
    var idx int = 0
    for i, separator := range *separators {
	curr = strings.SplitN(remainder, separator, 2)
	words[(*names)[i]] = curr[0]
	remainder = curr[1]
	idx = i + 1
    }
    words[(*names)[idx]] = remainder
    return words
}

func formatInput(line string, names *[]string, separators *[]string) {
    words := strings.Split(line,"}")
    var name_sep []string
    for _, word := range words {
	name_sep = strings.Split(word,"{")
	if len(name_sep) < 2 {
	    *names = append(*names, name_sep[0])
	}else {
	    if name_sep[0] == "" {
		*names = append(*names, name_sep[1])
		continue
	    }
	    *separators = append(*separators, name_sep[0])
	    *names = append(*names, name_sep[1])
	}
    }
}
