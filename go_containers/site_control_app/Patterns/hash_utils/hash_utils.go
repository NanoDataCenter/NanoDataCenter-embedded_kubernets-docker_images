package hash_utils


func  Has_key( table *map[string]interface{}, key *string ) bool {

  _,flag := (*table)[*key]
  return flag


}