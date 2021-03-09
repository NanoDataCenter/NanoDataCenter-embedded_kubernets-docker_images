import json
def get_site_data(file=None):
    if file == None:
         file_handle = open("/data/redis_server.json",'r')
    else:
         file_handle = open(file,'r')
         
    data = file_handle.read()
    file_handle.close()
    return json.loads(data)

