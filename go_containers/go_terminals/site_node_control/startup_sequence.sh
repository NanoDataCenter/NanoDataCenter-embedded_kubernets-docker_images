
export site=LACIMA_SITE

export port=6379
export host=192.168.1.66
export graph_db=3.0


export graph_container_image="nanodatacenter/lacima_configuration"  
export graph_container_script='docker run   --log-driver local -ti --rm --network host      --mount type=bind,source=/home/pi/system_config,target=/data/     nanodatacenter/lacima_configuration  /usr/local/bin/bash ./run.bsh'	
export redis_container_name="redis"
export redis_container_image="nanodatacenter/redis"
export redis_start_script='docker run  --log-driver local -d  --network host   --name redis    --mount type=bind,source=/home/pi/mountpoint/redis,target=/data    nanodatacenter/redis /bin/bash  redis_control.bsh'

./site_node_terminal
