package utils

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func GetArgValue(args []string, arg_name string) (string, error) {
	if slices.Contains(args, arg_name) {
		port_arg_idx := IndexOf(arg_name, args)

		if len(args) > port_arg_idx-1 {

			return args[port_arg_idx+1], nil

		} else {
			str_error := fmt.Sprintf("Couldn't find arg value for %s in args %v", arg_name, args)
			return "", errors.New(str_error)

		}

	} else {
		str_error := fmt.Sprintf("Couldn't find arg %s in args %v", arg_name, args)
		return "", errors.New(str_error)
	}

}
