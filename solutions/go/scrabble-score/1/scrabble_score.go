package scrabble

import (
    "strings"
)

func Score(word string) int {
	word = strings.ToUpper(word)
    var points int = 0
    for _,x := range word{
        switch x{
            case 'A','E','I','O','U','L','N','R','S','T':
            	points++
            case 'D','G':
            	points+=2
        	case 'B','C','M','P':
            	points += 3
        	case 'F','H','V','W','Y':
            	points += 4
            case 'K':
            	points += 5
            case 'J','X':
            	points += 8
            case 'Q','Z':
            	points += 10
        }
    }
    return points
}
