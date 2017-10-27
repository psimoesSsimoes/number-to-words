package ntw

import (
	"fmt"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
)

var portugueseMegasSingular = []string{"", "mil", "milhão", "mil milhões", "bilião"}
var portugueseMegasPlural = []string{"", "mil", "milhões", "mil milhões", "bilhões"}
<<<<<<< HEAD
var portugueseAdjectives = []string{"", " e ", " e ", " e ", " e ", " e ", " e ", " e ", " e ", " e ", " e "}
=======
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
var portugueseUnits = []string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove"}
var portugueseHundreds = []string{"", "cem", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos", "cento"}
var portugueseTens = []string{"", "dez", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
var portugueseTeens = []string{"dez", "onze", "doze", "treze", "catorze", "quinze", "dezasseis", "dezasete", "dezoito", "dezanove"}

<<<<<<< HEAD
// JOIN From strings.Join
// following robs pike advice from golang proverbs "A little copying is better than a little dependency."
func Join(a []string, sep string) string {

	switch len(a) {

	case 0:

		return ""

	case 1:

		return a[0]

	case 2:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1]

	case 3:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1] + sep + a[2]

	}

	n := len(sep) * (len(a) - 1)

	for i := 0; i < len(a); i++ {

		n += len(a[i])

	}

	b := make([]byte, n)

	bp := copy(b, a[0])

	for _, s := range a[1:] {

		bp += copy(b[bp:], sep)

		bp += copy(b[bp:], s)

	}

	return string(b)

}

func IntegerToPortuguese_PT(input int) string {
=======
func IntegerToPortuguesePT(input int) string {
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "menos")
		input *= -1
	}
<<<<<<< HEAD
	//fmt.Printf("%d  %f\n", input, math.Floor(math.Log10(math.Abs(float64(input))))+1)

	triplets := integerToTriplets(input)
=======

	triplets := integerToTriplets(input)
	switch {
	case len(triplets) == 0:
		return "zero"
	}
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		//log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		//log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
<<<<<<< HEAD
		if hundreds > 0 {
			word := fmt.Sprintf("%s e", portugueseHundreds[hundreds])
			words = append(words, word)
=======
		if hundreds > 0 && units == 0 && tens == 0 {
			var word string
			if idx == 0 && len(words) != 0 {
				word = fmt.Sprintf("e %s", portugueseHundreds[hundreds])
			} else {
				word = fmt.Sprintf("%s", portugueseHundreds[hundreds])
			}
			words = append(words, word)
		} else if hundreds > 0 {
			if hundreds == 1 {
				hundreds = 10
			}
			word := fmt.Sprintf("%s e", portugueseHundreds[hundreds])
			words = append(words, word)

>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, portugueseUnits[units])
		case 1:
<<<<<<< HEAD
			word := fmt.Sprintf("%s ", portugueseTeens[units])
=======
			word := fmt.Sprintf("%s", portugueseTeens[units])
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
			words = append(words, word)
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s e %s", portugueseTens[tens], portugueseUnits[units])
				words = append(words, word)
			} else {

<<<<<<< HEAD
				word := fmt.Sprintf("%s ", portugueseTens[tens])
=======
				word := fmt.Sprintf("%s", portugueseTens[tens])
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
				words = append(words, word)
			}
			break
		}
	tripletEnd:
		switch triplet {
<<<<<<< HEAD
		case 0:
			break
		case 1:
			if mega := portugueseMegasSingular[idx]; mega != "" {
				word := fmt.Sprintf("%s ", mega)
=======
		case 1:
			if mega := portugueseMegasSingular[idx]; mega != "" {
				if idx == 4 {
					sum := 0
					for i := 0; i < len(triplets)-1; i++ {
						sum += triplets[i]
					}
					if sum == 0 {

					} else {
						words = append(words, "um")
					}

				} else if idx == 1 && portugueseUnits[idx] == words[0] {
					words = words[1:]
				}
				word := fmt.Sprintf("%s", mega)
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
				words = append(words, word)
			}
			break
		default:
			if mega := portugueseMegasPlural[idx]; mega != "" {
<<<<<<< HEAD

				word := fmt.Sprintf("%s ", mega)
=======
				if idx == 1 && portugueseUnits[idx] == words[0] {
					words = words[1:]
				}
				word := fmt.Sprintf("%s", mega)
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949
				words = append(words, word)
			}
			break
		}
	}

<<<<<<< HEAD
	return Join(words, "")
=======
	return strings.Join(words, " ")
>>>>>>> 334e856a8e7d85389c88f7c73dbca878351d6949

}
