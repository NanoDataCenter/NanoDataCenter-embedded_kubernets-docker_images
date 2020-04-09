

import redis
import json
import psycopg2
from psycopg2.extensions import ISOLATION_LEVEL_AUTOCOMMIT


class Postgres_Utilities(object):

   def __init__(self,logical_user):
       self.logical_user = logical_user
       file_handle = open("/data/redis_server.json",'r')
       data = file_handle.read()
       file_handle.close()
       self.redis_site = json.loads(data)
       self.fetch_username_password()
       self.setup_postgres_connection()


   def fetch_username_password(self):
       redis_site = self.redis_site
       self.redis_password_handle  = redis.StrictRedis( host = redis_site["host"], port = redis_site["port"], db= redis_site["redis_password_db"] , decode_responses=True)
       self.username = self.redis_password_handle.hget(self.logical_user,"username")
       self.password = self.redis_password_handle.hget(self.logical_user,"password")
       self.database = self.redis_password_handle.hget(self.logical_user,"database")
       self.host     = self.redis_password_handle.hget(self.logical_user,"host")
       
       
   def setup_postgres_connection(self):
       self.conn = psycopg2.connect(host=self.host,database=self.database, user=self.username, password=self.password)
       self.conn.set_isolation_level(ISOLATION_LEVEL_AUTOCOMMIT);
       print("conn",self.conn)
       self.cursor = self.conn.cursor()
       
   def close(self):
       self.cursor.close()