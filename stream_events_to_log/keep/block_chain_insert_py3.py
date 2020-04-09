from   ethereum_block_chain.web3_top_class import Web_Class_IPC
from   redis_support_py3.graph_query_support_py3 import  Query_Support
import redis
import datetime
import json

class Save_Block_Chain_Data(object):

   def __init__(self,redis_handle):
       signing_key = '/key_store/UTC--2019-12-08T20-29-05.205871190Z--75dca28623f88b105b8d0c718b4bfde0f1568688'
       ipc_socket = "/ipc/geth.ipc"
       self.w3 = Web_Class_IPC(ipc_socket,redis_handle,signing_key)
       print(self.w3.get_block_number())
       
       
   def append_data(self,contract_name,method,*data) :
        contract_object = self.w3.get_contract(contract_name)
        receipt = self.w3.transact_contract_data(contract_object,method, *data)
        return receipt
        
if __name__ == "__main__":
   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
   redis_site = json.loads(data)
   qs = Query_Support( redis_site )
   redis_data_handle = qs.get_redis_data_handle()
   redis_contract_handle = redis.StrictRedis( host = redis_site["host"] , port=redis_site["port"], db=redis_site["redis_contract_db"] )
   save_block_chain_data =  Save_Block_Chain_Data(redis_contract_handle)
   receipt = save_block_chain_data.append_data("EventHandler","transmit_event",["event_name","event_sub_id","data"])
   print(receipt)