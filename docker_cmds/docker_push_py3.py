import os
from docker_interface_py3 import Docker_Interface
from pprint import pprint
docker_ctrl = Docker_Interface()

docker_images = docker_ctrl.images()
docker_list = []
for i in docker_images:
   name = i.tags[0]
   name_list = name.split("/")
   if (len(name_list) == 2) and (name_list[0] =="nanodatacenter"):
      os.system("docker push "+name)
  
   