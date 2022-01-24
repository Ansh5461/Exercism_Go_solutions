package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    m := map[string]int{}
    m["quarter_of_a_dozen"] = 3
    m["half_of_a_dozen"] = 6
    m["dozen"] = 12
    m["small_gross"] = 120
    m["gross"] = 144
    m["great_gross"] = 1728
    return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_,err := units[unit]
    if err == false{
        return false
    } else {
    bill[item] = bill[item] + units[unit]
    }
return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exist := bill[item]
    if exist == false{
        return false
    }
    value, exist1 := units[unit]
    if exist1 == false{
        return false
    }
	if bill[item] >= value {
		bill[item] -= value
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, exists := bill[item]
    if exists == false {
        return 0, false
    } else {
    return val, true
    }
}
