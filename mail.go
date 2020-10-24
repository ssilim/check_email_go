// Author Sergey Silimyankin for Yandex

/*
Напишите функцию на golang, которая проверяет введенный пользователем e-mail на соответствие следующим правилам:
1). e-mail состоит из имени и доменной части, эти части разделяются символом "@";
2). доменная часть не короче 3 символов и не длиннее 256, является набором непустых строк, состоящих из символов a-z 0-9_- и разделенных точкой;
3). каждый компонент доменной части не может начинаться или заканчиваться символом "-";
4). имя (до @) не длиннее 128 символов, состоит из символов a-z0-9"._-;
5). в имени не допускаются две точки подряд;
6). если в имени есть двойные кавычки ", то они должны быть парными;
7). в имени могут встречаться символы "!,:", но только между парными двойными кавычками.
*/

package main

import (
	"fmt"
	"strings"
)

const symbols = "abcdefghijklmnopqrstuvwxyz_-0123456789\"!:."

func main() {

	mails := []string{"test@yandex.ru",
		"test123@yandex.ru",
		"tes._-t@yandex.ru",
		"t\"e\"st@yandex.ru",
		"test\"!\"@yandex.ru",
		"te\":\"st@yandex.ru",
	}

	for _, i := range mails {
		checkEmail(i)
	}
}

func allowChars(s string, symb string) bool {

	// Проверяем строку на допустимые символы

	for _, char := range s {
		if !strings.Contains(symb, string(char)) {
			return false
		}
	}
	return true
}

func checkSobaka(s string) bool {

	// e-mail состоит из имени и доменной части, эти части разделяются символом "@";

	if strings.Contains(s, "@") {
		return true
	}
	return false
}

func checkDomainSymbols(s string) bool {

	// Доменная часть не короче 3 символов и не длиннее 256, является набором непустых строк, состоящих из символов a-z 0-9_- и разделенных точкой;

	if len(s) > 3 && len(s) < 256 && len(s) > 0 && allowChars(s, symbols) {
		return true
	}
	return false
}

func checkBannedSymbols(s string, ss string) bool {

	// Каждый компонент доменной части не может начинаться или заканчиваться символом "-";

	if !strings.HasPrefix(s, "-") && !strings.HasSuffix(s, "-") && !strings.HasPrefix(ss, "-") && !strings.HasSuffix(ss, "-") {
		return true
	}
	return false
}

func checkNameSymbols(s string) bool {

	// Имя (до @) не длиннее 128 символов, состоит из символов a-z0-9"._-;

	if len(s) < 128 && allowChars(s, symbols) {
		return true
	}
	return false
}

func checkNameTwoDots(s string) bool {

	// В имени не допускаются две точки подряд;

	if s[strings.Index(s, ".")+1:strings.Index(s, ".")+2] != "." {
		return true
	}
	return false
}

func checkNameTwoQuotes(s string) bool {

	// Eсли в имени есть двойные кавычки ", то они должны быть парными;

	if strings.Count(s, "\"")%2 == 0 {
		return true
	}
	return false
}

func checkNameExclameColon(s string) bool {

	// В имени могут встречаться символы "!,:", но только между парными двойными кавычками.

	if strings.Contains(s, "!") || strings.Contains(s, ":") {
		if s[strings.Index(s, "!")+1:strings.Index(s, "!")+2] == "\"" && s[strings.Index(s, "!")-1:strings.Index(s, "!")] == "\"" || s[strings.Index(s, ":")+1:strings.Index(s, ":")+2] == "\"" && s[strings.Index(s, ":")-1:strings.Index(s, ":")] == "\"" {
			return true
		}
		return false
	}
	return true
}

func checkEmail(email string) {

	s := strings.Split(email, "@") // Разделяем по символу @

	if checkSobaka(email) {

		if checkDomainSymbols(s[1]) {

			if checkBannedSymbols(s[0], s[1]) {

				if checkNameSymbols(s[0]) {

					if checkNameTwoDots(s[0]) {

						if checkNameTwoQuotes(s[0]) {

							if checkNameExclameColon(s[0]) {
								fmt.Println("OK: ", email)

							} else {
								fmt.Println("7. Error: ", s[0], "В имени могут встречаться символы \"!,:\", но только между парными двойными кавычками.")
							}
						} else {
							fmt.Println("6. Error: ", s[0], "Если в имени есть двойные кавычки \", то они должны быть парными;")
						}
					} else {
						fmt.Println("5. Error: ", s[0], "В имени не допускаются две точки подряд;")
					}
				} else {
					fmt.Println("4. Error: ", s[0], "Имя (до @) не длиннее 128 символов, состоит из символов a-z0-9\"._-;")
				}
			} else {
				fmt.Println("3. Error: ", s[0], s[1], "Каждый компонент доменной части не может начинаться или заканчиваться символом '-'")
			}
		} else {
			fmt.Println("2. Error: ", s[1], "Доменная часть не короче 3 символов и не длиннее 256, является набором непустых строк, состоящих из символов a-z 0-9_- и разделенных точкой;")
		}
	} else {
		fmt.Println("1. Error: ", email, "В имени отсутствует символ @")
	}
}
