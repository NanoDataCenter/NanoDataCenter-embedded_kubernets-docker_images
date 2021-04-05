#
#
#  push all nano_data_center images to github
#
#

from docker_control.docker_interface_py3 import Docker_Interface


def push_image(name,tag):
   name_list = name.split("/")
   
   if (len(name_list)==2) and (name_list[0] =="nanodatacenter"):
      print("pushing image",name)
      print(docker_control.push(name))
   else:
      print("invalid",name,tag)

docker_control = Docker_Interface()
images = docker_control.images_raw()
for i in images:
   name = i[0]
   tag = i[1]
   if tag == "latest":
      push_image(name,tag)
   else:
      print("invalid name",name,tag)

