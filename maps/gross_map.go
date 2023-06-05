package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creats new customer bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if unitVal, uExists := units[unit]; uExists {
		bill[item] += unitVal
		return true
	}
	return false
}

// RemoveItem removes item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitVal, uExists := units[unit]
	itemVal, iExists := bill[item]
	if uExists && iExists {
		itemVal -= unitVal
		switch {
		case itemVal < 0:
			return false
		case itemVal == 0:
			delete(bill, item)
			return true
		default:
			bill[item] = itemVal
			return true
		}
	}
	return false
}

// GetItem returns the quantitiy of an item that customer has in their bill
func GetItem(bill map[string]int, item string)(int, bool) {
	itemVal, exists := bill[item]
	return itemVal, exists
}
