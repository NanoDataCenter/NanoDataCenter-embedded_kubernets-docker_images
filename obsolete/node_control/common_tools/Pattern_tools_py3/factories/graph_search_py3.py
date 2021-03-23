

def common_qs_search(site_data,qs,search_list): # generalized graph search
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    for i in range(0,len(search_list)-1):
        if type(search_list[i]) == list:
           query_list = qs.add_match_relationship( query_list,relationship = search_list[i][0],label = search_list[i][1] )
        else:
           query_list = qs.add_match_relationship( query_list,relationship = search_list[i] )
           
    if type(search_list[-1]) == list:
       query_list = qs.add_match_terminal( query_list,relationship = search_list[-1][0],label = search_list[-1][1] )
    else:
       query_list = qs.add_match_terminal( query_list,relationship = search_list[-1] )
       
    node_sets, node_sources = qs.match_list(query_list)        
    return node_sources
    
def common_package_search(site_data,qs,search_list): # generalized graph search
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    for i in range(0,len(search_list)-1):
        if type(search_list[i]) == list:
           query_list = qs.add_match_relationship( query_list,relationship = search_list[i][0],label = search_list[i][1] )
        else:
           query_list =qs.add_match_relationship( query_list,relationship = search_list[i] )
           
   
    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", 
                                        property_mask={"name":search_list[-1]} )
       
    package_sets, package_sources = qs.match_list(query_list)
    
    return package_sources
    