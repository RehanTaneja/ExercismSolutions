package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units map[string]int = make(map[string]int)
    units["quarter_of_a_dozen"]=3
    units["half_of_a_dozen"]=6
    units["dozen"]=12
    units["small_gross"]=120
    units["gross"]=144
    units["great_gross"]=1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var bill map[string]int = make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	var inUnits bool = false
    for i := range units{
        if i==unit{
            inUnits = true
        }
    }
    if inUnits == false{
        return false
    }else{
        var inBill bool = false
        for i := range bill{
            if i == item{
                inBill = true
            }
        }
        if inBill == false{
            bill[item] = units[unit]
        }else{
            bill[item] += units[unit]
        }
        return true        
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	var inBill bool = false
    for i := range bill{
        if i == item{
            inBill = true
        }
    }
    if inBill == false{
        return false
    }
    var inUnits bool = false
    for i := range units{
        if unit == i{
            inUnits = true
        }
    }
    if inUnits == false{
        return false
    }
    var newQuantity = bill[item] - units[unit]
    if newQuantity < 0{
        return false
    }else if newQuantity == 0{
        delete(bill, item)
        return true
    }else{
        bill[item] = newQuantity
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	var inBill bool = false
    for i := range bill{
        if i == item{
            inBill = true
        }
    }
    if inBill == false{
        return 0,false
    }else{
        return bill[item],true
    }
}
