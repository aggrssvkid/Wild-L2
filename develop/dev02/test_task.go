package main

import "fmt"

func main() {
	// проверяем пустую строку
	base_str := ""
	expected := ""
	output_str := Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string: empty string   ", "[FALSE]")
	} else {
		fmt.Println("Base string: empty string   ", "[OK]")
	}

	// проверяем строку, в которой нет чисел, соответственно она не должна быть изменена
	base_str = "abcd"
	expected = "abcd"
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// проверяем на валидность, если строка не валидна, возвращаем пустую строку
	base_str = "1ad"
	expected = ""
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// проверяем на конвертацию только числа (без слэша)
	base_str = "a1b2c3d4"
	expected = "abbcccdddd"
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// проверяем слэш до цифры
	base_str = `abcd\4`
	expected = `abcd4`
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// проверяем слэш + число после цифры
	base_str = `a\410`
	expected = `a4444444444`
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// проверяем слэш после слэша
	base_str = `\\\\end`
	expected = `\\end`
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// проверяем на валидность. Строка не должна заканчиваться одиночным слэшэм
	base_str = `\\\\\`
	expected = ""
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}

	// суммируем разные варианты
	base_str = `\1ab7\2\\end4`
	expected = `1abbbbbbb2\endddd`
	output_str = Convert_string(base_str)
	if output_str != expected {
		fmt.Println("Base string:", base_str, "   [FALSE]")
	} else {
		fmt.Println("Base string:", base_str, "   [OK]")
	}
}
