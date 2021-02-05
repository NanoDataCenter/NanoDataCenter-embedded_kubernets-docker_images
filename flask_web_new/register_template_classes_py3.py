

from templates.irrigation.control_by_schedule_py3  import Control_By_Schedule
from templates.irrigation.control_by_controller_py3 import Control_By_Controller
from templates.irrigation.control_by_valve_group_py3 import Control_By_Valve_Group
from templates.irrigation.set_irrigation_parameters_py3  import Irrigation_Parameters
from templates.irrigation.control_irrigation_queue_py3 import Control_Irrigation_Queue
from templates.irrigation.past_actions_py3 import Past_Actions
from templates.irrigation.irrigation_stream_py3 import Irrigation_Stream


from templates.container_control.container_exception_status_py3 import Container_Exception_Status
from templates.container_control.container_exception_log_py3 import Container_Exception_Log
from templates.container_control.container_processor_logs_py3 import Container_Cpu_Loading
from templates.container_control.container_processor_logs_py3 import Container_Vsz_Loading
from templates.container_control.container_processor_logs_py3 import Container_Rss_Loading
from templates.container_control.start_and_stop_managed_container_processes_py3 import Start_and_Stop_Managed_Container_Processes

from templates.node_control.node_exception_log_py3 import Node_Exception_Log
from templates.node_control.node_exception_status_py3 import Node_Exception_Status
from templates.node_control.node_control_start_stop_processes_py3 import Node_Processor_Control
from templates.node_control.manage_containers_py3 import Manage_Containers


from templates.common.table_manager.table_manager_py3 import Table_Manager

from templates.common.site_manager_py3 import Site_Manager

from templates.redis.redis_streams_py3 import  *
from templates.processor_performance.process_performance_py3 import *


class Register_Template_Classes( object):
   def __init__(self,parent_self):
       self.parent_self = parent_self
       
       class_map = {}
       class_map["irrigation/control_by_schedule"]  = Control_By_Schedule
       class_map["irrigation/control_by_controller"]  = Control_By_Controller
       class_map["irrigation/control_by_valve_group"]  = Control_By_Valve_Group
       class_map["irrigation/set_irrigation_parameters"]  = Irrigation_Parameters
       class_map["irrigation/control_irrigation_queue"]  = Control_Irrigation_Queue
       class_map["irrigation/past_actions"] = Past_Actions
       class_map["irrigation/stream_manager"] = Irrigation_Stream
       
       class_map["common/table_manager"]  = Table_Manager
       class_map["common/redis_stream_manager"]  = Redis_Stream_Manager
       class_map["common/site_manager"]  = Site_Manager
       class_map["redis/key_stream"] = Redis_Key_Stream
       class_map["redis/call_stream"] = Redis_Call_Stream
       class_map["redis/command_time"] = Redis_Command_Time_Stream
       class_map["redis/client_stream"] = Redis_Client_Stream
       class_map["redis/server_time"] = Redis_Server_Time_Stream
       class_map["redis/memory_stream"] = Redis_Memory_Stream
       
 
       class_map["processor/free_cpu"] = Processor_Free_CPU
       class_map["processor/ram"] = Processor_Free_Ram
       class_map["processor/disk_space"] = Processor_Free_Disk
       class_map["processor/temperature"] = Processor_Temperature 
       class_map["processor/cpu_core"] = Processor_Core
       class_map["processor/swap_space"] = Processor_Swap_Space
       class_map["processor/io_space"] =Processor_Io_Space
       class_map["processor/block_dev"] = Processor_Block_Dev
       class_map["processor/context_switches"] = Processor_Context_Switches
       class_map["processor/run_queue"] =Processor_Run_Queue
       class_map["processor/edev"] = Processor_Edev

       
       
       class_map["manage_containers/view_exception_status"] = Container_Exception_Status
       class_map["manage_containers/view_exception_log"]  = Container_Exception_Log
       class_map["manage_containers/cpu_loading"]  = Container_Cpu_Loading
       class_map["manage_containers/vsz"]  = Container_Vsz_Loading
       class_map["manage_containers/rss"]  = Container_Rss_Loading
       class_map["manage_containers/start_and_stop_processes"] = Start_and_Stop_Managed_Container_Processes
       
       class_map["node_control/view_exception_log"] = Node_Exception_Log
       class_map["node_control/view_exception_status"] = Node_Exception_Status
       class_map["node_control/process_control"] = Node_Processor_Control
       class_map["node_control/manage_containers"] = Manage_Containers
       self.parent_self.class_map = class_map
      
      
 
