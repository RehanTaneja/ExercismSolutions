package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var num int = 10
    switch colors[0]{
        case "black":
        	num *= 0
        case "brown":
        	num *= 1
        case "red":
        	num *= 2
        case "orange":
        	num *= 3
        case "yellow":
        	num *= 4
        case "green":
        	num *= 5
        case "blue":
        	num *= 6
        case "violet":
        	num *= 7
        case "grey":
        	num *= 8
        case "white":
        	num *= 9
        default:
        	return -1
    }
    switch colors[1]{
        case "black":
        	num += 0
        case "brown":
        	num += 1
        case "red":
        	num += 2
        case "orange":
        	num += 3
        case "yellow":
        	num += 4
        case "green":
        	num += 5
        case "blue":
        	num += 6
        case "violet":
        	num += 7
        case "grey":
        	num += 8
        case "white":
        	num += 9
        default:
        	return -1
    }
    return num
}
