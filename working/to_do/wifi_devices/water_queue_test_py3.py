
import redis
import msgpack
import time

    
    
redis_handle = redis.StrictRedis(  db=11 )

payload = {}
payload["command"] = "SEND_ALERT"
payload["message"] = "Water On"
payload["duration"] =1000

payload_msgpack = msgpack.packb(payload)
for i in range(0,1000):
   time.sleep(10)
   print(payload_msgpack)
   redis_handle.rpush("OP_1/INPUT",payload_msgpack)
