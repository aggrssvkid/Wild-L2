package main

import (
	"strconv"
)

func Convert_string(input_str string) string {
	runes := []rune(input_str)
	output_str := make([]rune, 0)

	var start int
	var i int
	len_str := len(runes)

	if is_valid(runes) == false {
		return ""
	}

	for i < len_str {
		if runes[i] >= '0' && runes[i] <= '9' {
			output_str, i = convert_nbr(runes, output_str, i, start)
			if i == len_str {
				return string(output_str)
			}
			start = i
		} else if runes[i] == '\\' {
			output_str, i = convert_slash(runes, output_str, i, start)
			start = i
		} else {
			i++
		}
	}
	if start != i {
		output_str = append(output_str, runes[start:]...)
	}
	return string(output_str)
}

func convert_nbr(runes []rune, output_str []rune, i int, start int) ([]rune, int) {
	output_str = append(output_str, runes[start:i-1]...)
	end := i
	for runes[end] >= '0' && runes[end] <= '9' {
		end++
		if end == len(runes) {
			break
		}
	}
	string_nbr := runes[i:end]
	nbr, _ := strconv.Atoi(string(string_nbr))
	symbol_slice := make([]rune, nbr)
	for j, _ := range symbol_slice {
		symbol_slice[j] = runes[i-1]
	}
	output_str = append(output_str, symbol_slice...)
	return output_str, end
}

func convert_slash(runes []rune, output_str []rune, i int, start int) ([]rune, int) {
	output_str = append(output_str, runes[start:i]...)
	if i+2 < len(runes) && (runes[i+2] >= '0' && runes[i+2] <= '9') {
		output_str, i = convert_nbr(runes, output_str, i+2, i+1)
	} else {
		output_str = append(output_str, runes[i+1])
		i += 2
	}
	return output_str, i
}

func is_valid(runes []rune) bool {
	// дальше будем обращаться к элементам по индексу, поэтому проверяем на пустую строку
	if len(runes) == 0 {
		return true
	}

	// строка не должна начинаться с цифры, так как нечего конвертировать
	if runes[0] >= '0' && runes[0] <= '9' {
		return false
	}

	// проверяем на одиночный слэш на конце (после слэша должен быть символ)
	if runes[len(runes)-1] == '\\' {
		j := len(runes) - 1
		k := 0
		for j >= 0 {
			if runes[j] == '\\' {
				k++
			} else {
				break
			}
			j--
		}
		if k%2 != 0 {
			return false
		}
	}
	return true
}
