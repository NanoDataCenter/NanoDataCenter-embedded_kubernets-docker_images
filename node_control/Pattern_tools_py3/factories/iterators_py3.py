  



def pattern_iter_find_lowest(input_data,sort_key,data_key):
    value = input_data[data_key]
    ref_value = input_data[sort_key]
    for i,item in input_data.items():
       if data[sort_key] < ref_value:
          ref_value = item[sort_key]
          value = item[data_key]
    return value

def pattern_iter_strip_dict_dict( input_data,filter_key ): # strip dictionary of items where dict
    return_value = {}
    for key,item in input_data.items():
       return_value[key] = item[filter_key]
    return return_value   
