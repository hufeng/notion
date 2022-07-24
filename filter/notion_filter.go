package filter

type Filter map[string]any

func NewFilter(f Filter) Filter {
	return Filter{"filter": f}
}

// compose condition
func And(filters ...Filter) Filter {
	return Filter{"and": filters}
}

func Or(filters ...Filter) Filter {
	return Filter{"or": filters}
}

// text prop filter
func withTextPropFilter(prop string, cond string, val any) Filter {
	return Filter{"property": prop, "text": Filter{cond: val}}
}

func TextProp__equals(prop, val string) Filter {
	return withTextPropFilter(prop, "equals", val)
}

func TextProp__does_not_equal(prop, val string) Filter {
	return withTextPropFilter(prop, "do_not_equlas", val)
}

func TextProp__contains(prop, val string) Filter {
	return withTextPropFilter(prop, "contains", val)
}

func TextProp__does_not_contain(prop, val string) Filter {
	return withTextPropFilter(prop, "does_not_contain", val)
}

func TextProp__starts_with(prop, val string) Filter {
	return withTextPropFilter(prop, "starts_with", val)
}

func TextProp__ends_with(prop, val string) Filter {
	return withTextPropFilter(prop, "ends_with", val)
}

func TextProp__is_empty(prop string) Filter {
	return withTextPropFilter(prop, "is_empty", true)

}

func TextProp__is_not_empty(prop string) Filter {
	return withTextPropFilter(prop, "is_not_empty", true)
}

// number filter
func withNumPropFilter(prop, cond string, val any) Filter {
	return Filter{"property": prop, "number": Filter{cond: val}}
}

func NumProp__equals(prop string, val int) Filter {
	return withNumPropFilter(prop, "equals", val)
}

func NumProp__does_not_equal(prop string, val int) Filter {
	return withNumPropFilter(prop, "does_not_equal", val)
}

func NumProp__greater_than(prop string, val int) Filter {
	return withNumPropFilter(prop, "greater_than", val)
}

func NumProp__less_than(prop string, val int) Filter {
	return withNumPropFilter(prop, "less_than", val)
}

func NumProp__greater_than_or_equal_to(prop string, val int) Filter {
	return withNumPropFilter(prop, "greater_than_or_equla_to", val)
}

func NumProp__less_than_or_equal_to(prop string, val int) Filter {
	return withNumPropFilter(prop, "less_than_or_equla_to", val)
}

func NumProp__isEmpty(prop string) Filter {
	return withNumPropFilter(prop, "is_empty", true)
}

func NumProp__is_not_empty(prop string) Filter {
	return withNumPropFilter(prop, "is_not_empty", true)
}

// Checkbox filter
func withCheckboxFilter(prop string, cond string, val bool) Filter {
	return Filter{"property": prop, "checkbox": Filter{cond: val}}
}

func CheckboxProp__equals(prop string, val bool) Filter {
	return withCheckboxFilter(prop, "equals", val)
}

func CheckboxProp__does_not_equal(prop string, val bool) Filter {
	return withCheckboxFilter(prop, "does_not_equal", val)
}

// Select filter
func withSelectFilter(prop string, cond string, val any) Filter {
	return Filter{"property": prop, "select": Filter{cond: val}}
}

func SelectProp__equal(prop, val string) Filter {
	return withSelectFilter(prop, "equals", val)
}

func SelectProp__does_not_equal(prop, val string) Filter {
	return withSelectFilter(prop, "does_not_equals", val)
}

func SelectProp__is_empty(prop string) Filter {
	return withSelectFilter(prop, "is_empty", true)
}

func SelectProp__is_not_empty(prop string, val bool) Filter {
	return withSelectFilter(prop, "is_not_empty", true)
}

// multiple select
func withMulSelectFilter(prop string, cond string, val any) Filter {
	return Filter{"property": prop, "multi_select": Filter{cond: val}}
}

func MulSelectProp__contains(prop, val string) Filter {
	return withMulSelectFilter(prop, "contains", val)
}

func MulSelectProp__does_not_contain(prop, val string) Filter {
	return withMulSelectFilter(prop, "does_not_contain", val)
}

func MulSelectProp__is_empty(prop string) Filter {
	return withMulSelectFilter(prop, "is_empty", true)
}

func MulSelectProp__is_not_empty(prop string) Filter {
	return withMulSelectFilter(prop, "is_not_empty", true)
}

// date filter
func withDateFilter(prop string, cond string, val any) Filter {
	return Filter{"property": prop, "date": Filter{cond: val}}
}

func DateProp__equals(prop string, val string) Filter {
	return withDateFilter(prop, "equals", val)
}

func DateProp__before(prop string, val string) Filter {
	return withDateFilter(prop, "before", val)
}

func DateProp__after(prop string, val string) Filter {
	return withDateFilter(prop, "after", val)
}

func DateProp__on_or_before(prop string, val string) Filter {
	return withDateFilter(prop, "on_or_before", val)
}

func DateProp__is_empty(prop string) Filter {
	return withDateFilter(prop, "is_empty", true)
}

func DateProp__is_not_empty(prop string) Filter {
	return withDateFilter(prop, "is_not_empty", prop)
}

func DateProp__on_or_after(prop string, val string) Filter {
	return withDateFilter(prop, "on_or_after", val)
}

func DateProp__past_week(prop string) Filter {
	return withDateFilter(prop, "past_week", Filter{})
}

func DateProp__past_month(prop string) Filter {
	return withDateFilter(prop, "past_month", Filter{})
}

func DateProp__past_year(prop string) Filter {
	return withDateFilter(prop, "past_year", Filter{})
}

func DateProp__next_week(prop string) Filter {
	return withDateFilter(prop, "next_week", Filter{})
}

func DateProp__next_month(prop string) Filter {
	return withDateFilter(prop, "next_month", Filter{})
}

func DateProp__next_year(prop string) Filter {
	return withDateFilter(prop, "past_year", Filter{})
}

// people filter
func withPeopleFilter(prop string, cond string, val any) Filter {
	return Filter{"property": "prop", "people": Filter{cond: val}}
}

func PeopleProp__contains(prop string, val string) Filter {
	return withPeopleFilter(prop, "contains", val)
}

func PeopleProp__does_not_contain(prop string, val string) Filter {
	return withPeopleFilter(prop, "does_not_contain", val)
}

func PeopleProp__is_empty(prop string) Filter {
	return withPeopleFilter(prop, "is_empty", true)
}

func PeopleProp__is_not_empty(prop string) Filter {
	return withPeopleFilter(prop, "is_not_empty", true)
}

// file filter
func withFileFilter(prop string, cond string, val any) Filter {
	return Filter{"property": prop, "file": Filter{cond: val}}
}

func FileProp__is_empty(prop string) Filter {
	return withFileFilter(prop, "is_empty", true)
}

func FileProp__is_not_empty(prop string) Filter {
	return withFileFilter(prop, "is_not_empty", true)
}
