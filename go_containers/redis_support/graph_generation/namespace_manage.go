package graph_generation


func (v *Build_Configuration)push_namespace( name [2]string){

   if v.namespace_len < v.namespace_max_len {
       v.namespace[v.namespace_len] = name
	   v.namespace_len += 1
   }else{
       v.namespcae = append(v.namespace,name)
	   v.namespace_len +=1
	   v.namespace_max_len +=1
	}


}


func (v *Build_Configuration)pop_namespace() [2]string {

   return_value := v.namespace[v.namespace_len-1]
   v.namespace_len -= 1
   return return_value
}


func (v *Build_Configuration)get_last() [2]string {
   
    return v.namespace[v.namespace_len-1]
}
