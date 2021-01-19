  



  def container_exception_log(self,processor_id):
       processor_name = self.processor_names[processor_id]
       temp_list = self.container_control_structure[processor_name]["ERROR_STREAM"].revrange("+","-" , count=20)

       container_exceptions = []
       for j in temp_list:
           i = j["data"]
           i["timestamp"] = j["timestamp"]
           i["datetime"] =  datetime.datetime.fromtimestamp( i["timestamp"]).strftime('%Y-%m-%d %H:%M:%S')

           temp = i["error_output"]
           if len(temp) > 0:
               temp = i["error_output"]
               if len(temp) > 0:
                   temp = [temp]
                   #temp = temp.split("\n")
                   i["error_output"] = temp
                   container_exceptions.append(i)
       
       return self.render_template(self.path_dest+"/docker_exception_log",                                 
                                  log_data = container_exceptions,
                                  processor_id = processor_id,
                                  processors = self.processor_names ) 



{% extends "base_template" %}

{% block application_javascript %}
  <script type="text/javascript" >
       False = false
       True = true
       None = null
       container_id = {{container_id}}                         
                               
      

      </script>
      <script type="text/javascript">

function change_container(event,ui)
{
  current_page = window.location.href
 
 
  current_page = current_page.slice(0,-2)
  
  current_page = current_page+"/"+$("#container_select")[0].selectedIndex
  
  window.location.href = current_page
}


 $(document).ready(
 function()
 {
   
   
   $("#container_select").val( {{ container_id|int  }});
   $("#container_select").bind('change',change_container)   
   

 }
)
 
</script>
  
{% endblock %}

{% block application %}
<div class="container">
<center>
<h4>Select Container</h4>
</center>

<div id="select_tag">
<center>
<select id="container_select">
  {% for item in containers %}
  
  <option value="{{loop.index0}}">{{item}}</option>
  {% endfor %}
  
</select>
</center>
</div>
<div style="margin-top:20px"></div>
<h4>Exception Log </h4>

{% for item in log_data %}
<div style="margin-top:10px"></div>
<h5>{{item.script}}</h5>
<ul>  
<li>script: {{item.script}} </li>
<li>crc: {{item.crc}} </li>
<li>timestamp: {{item.timestamp}} </li>
<li>date time:  {{item.datetime}} </li>
<li>exception stack:</li>
<ul>
{% for line in item.error_output %}
 <li>{{line}}</li>
{% endfor %}
</ul>
</ul>
{% endfor %}
</div>

{% endblock %}                                  