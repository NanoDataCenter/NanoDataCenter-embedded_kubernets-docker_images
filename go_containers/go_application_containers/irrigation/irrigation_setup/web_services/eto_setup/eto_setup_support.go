package eto_setup

import (
   
	"strings"
)

func  create_change_recharge_level(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_recharge_html(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")
}

func create_recharge_html( div string )string{
    var config   ETO_Parameter_Setup
   
    config.Save_flag      = true
    config.Div_name       = div
    config.Title          = "Change Recharge Level"
    config.Field          = "recharge_level"
    config.Parm_name      = "recharge_level"
   
    config.Helper1        = "ETO Recharge Level:"
    config.Helper2        = "1"
    config.Helper3        = " inch"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
    
    
    
    





func  create_crop_utilization(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_crop_utilization_html(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

func create_crop_utilization_html( div string )string{
    var config   ETO_Parameter_Setup
   
    config.Save_flag      = true
    config.Div_name       = div
    config.Title          = "Change Crop Utilization"
    config.Field          = "crop_utilization"
    config.Parm_name      = "crop_utilization"
    
    config.Helper1        = "Crop Utilization: "
    config.Helper2        = "100"
    config.Helper3        = " %"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
  
func  create_change_salt_flush(div_name string)string{
     size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_salt_flush_html(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")
}

func create_salt_flush_html( div string )string{
    var config   ETO_Parameter_Setup
   
    config.Save_flag      = true
    config.Div_name       = div
    config.Title          = "Change Salt Flush"
    config.Field          = "salt_flush"
    config.Parm_name      = "salt_flush"
    
    config.Helper1        = "Salt Flush:"
    config.Helper2        = "100"
    config.Helper3        = " %"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
 


func  create_change_sprayer_efficiency(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_sprayer_efficiency_html(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

func create_sprayer_efficiency_html( div string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = true
    config.Div_name       = div
    config.Title          = "Change Sprayer Efficiency"
    config.Field          = "sprayer_effiency"
    config.Parm_name      = "sprayer_effiency"
   
    config.Helper1        = "Sprayer Efficiency: "
    config.Helper2        = "100"
    config.Helper3        = " %"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
 


func  create_change_sprayer_rate(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_sprayer_rate_html(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

func create_sprayer_rate_html( div string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = true
    config.Div_name       = div
    config.Title          = "Change Sprayer Rate "
    config.Field          = "sprayer_rate"
    config.Parm_name      = "sprayer_rate"
    
    config.Helper1        = "Sprayer Rate: "
    config.Helper2        = "1"
    config.Helper3        = " GPM"
    config.Max_Range      = "50"
    config.Step           = ".1"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
 
 
func create_change_tree_radius( div_name string )string{
   size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_tree_radius_html(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

    
func create_tree_radius_html( div  string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = true
    config.Div_name       = div
    config.Title          = "Change Tree Radius"
    config.Field          = "tree_radius"
    config.Parm_name      = "tree_radius"
    
    config.Helper1        = "Tree Radius: "
    config.Helper2        = "1"
    config.Helper3        = " feet"
    config.Max_Range      = "15"
    config.Step           = ".1"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}    
    
