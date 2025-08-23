package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsMap := map[string]int{}
    unitsMap["quarter_of_a_dozen"] = 3
    unitsMap["half_of_a_dozen"] = 6
    unitsMap["dozen"] = 12
    unitsMap["small_gross"] = 120
    unitsMap["gross"] = 144
    unitsMap["great_gross"] = 1728
    return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int, 0)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
    if !exists {
        return false
    }
    if exists {
        _, existIt := bill[item]
        if existIt {
            bill[item] += units[unit]
        } else {
        	bill[item] = units[unit]
        }
        return true
    }
    return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, existItem := bill[item]
    if !existItem {
        return false
    }
    _, existUnit := units[unit]
    if !existUnit {
        return false
    }
    quantity := bill[item]
    if quantity - units[unit] < 0 {
        return false
    } else if quantity - units[unit] == 0 {
		delete(bill, item)
        return true
    } else {
		bill[item] -= units[unit]
        return true
    }
    return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, existItem := bill[item]
    if !existItem {
        return 0, false
    }
    return bill[item], true
}
