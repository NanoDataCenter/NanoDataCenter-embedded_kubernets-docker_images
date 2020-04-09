
import json
import sys

contract_names = []
for i in  range(1,len(sys.argv),1): 
   contract_names.append(sys.argv[i])
   
contract_string = json.dumps(contract_names)
File_object = open("contracts_to_load.json","w")
File_object.write(contract_string)
File_object.close()