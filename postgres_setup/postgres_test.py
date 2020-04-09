
import psycopg2
from psycopg2.extensions import ISOLATION_LEVEL_AUTOCOMMIT

conn = psycopg2.connect(host="127.0.0.1",database="logger", user="logger", password="ready2go")
conn.set_isolation_level(ISOLATION_LEVEL_AUTOCOMMIT);
print("conn",conn)
cursor = conn.cursor()
print("cursor",cursor)
try:
    # use the execute() method to make a SQL request
    result = cursor.execute('CREATE DATABASE LOGGING_DATABASE')
    print("result",result)
except:
     cursor.close()
     raise
finally:
    cursor.close()