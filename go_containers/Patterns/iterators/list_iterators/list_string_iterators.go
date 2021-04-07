package list_interators_string


func Is_value_in_list(value string, list []string) bool {
    for _, v := range list {
        if v == value {
            return true
        }
    }
    return false
}
