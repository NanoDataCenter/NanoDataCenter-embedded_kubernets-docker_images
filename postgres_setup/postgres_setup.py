
from postgres_utilities.postgres_utilities import Postgres_Utilities

pu = Postgres_Utilities("postgres_logging")

print(pu.username,pu.password,pu.database)

pu.close()