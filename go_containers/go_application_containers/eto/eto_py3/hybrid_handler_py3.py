
import urllib.request
import time
import datetime
import json

from .calculate_eto_py3 import Calculate_ETO
class Hybrid_Calculator( object ):

   def __init__(self,data,eto_config,eto_sources,rain_sources):
       
       self.calculate_eto = Calculate_ETO()
       self.data = data
       self.pws = data["pws"]
       self.priority = data["priority"]
       self.eto_config = eto_config
       self.eto_sources = eto_sources
       self.rain_sources = rain_sources
       self.main = data["stations"]["main"]
       self.solar = data["stations"]["solar"]
       
       
       
   def compute_previous_day( self): 
       if self.eto_sources.hget("wunder_hybrid:"+self.pws+":Normal") != None:
          print("*********************","am returning hybrid")
          return
       #identify solar
       for i in self.eto_config:
           if i["name"] == self.solar:
               self.solar_calc = i["calculator"]
       for i in self.eto_config:
           if i["name"] == self.main:
               self.main_calc = i["calculator"]
               self.lat = i["lat"]
               self.alt = i["alt"]
       self.main_calc.compute_previous_day(flag=True) 
       main_data = self.main_calc.results_normal       
       self.solar_calc.compute_previous_day(flag=True)
       solar_data =  self.solar_calc.results_normal   
       print("************************************************************************************")
      
       if len(main_data) != len(solar_data):
          raise ValueError("data length not match")
       
       
       for i in range(0,len(main_data)):
           main_data[i]["SolarRadiationWatts/m^2"] = solar_data[i]["SolarRadiationWatts/m^2"]
       eto = self.calculate_eto.__calculate_eto__(main_data,self.alt,self.lat)
      
       self.eto_sources.hset("wunder_hybrid:"+self.pws+":Normal", { "eto":eto,"priority":self.priority,"status":"OK" ,"time":str(datetime.datetime.now())}       ) 
       
       
"""
class Wunder_Personal( object ):

   def __init__(self,data,eto_sources,rain_sources):
    
     self.access_key = data["access_key"]
     self.pws = data["pws"]
     self.alt  = data["alt"]
     self.lat  = data["lat"]
     self.priority = data["priority"]
     self.calculate_eto = Calculate_ETO()
     self.eto_sources = eto_sources
     self.rain_sources = rain_sources
     
   #https://api.weather.com/v2/pws/observations/hourly/7day?stationId=KCAMURRI101&format=json&units=e&apiKey=dc0b888f054d45a88b888f054db5a83b  
     
   def compute_previous_day( self):
       
       if self.eto_sources.hget("wunder:"+self.pws+":Normal") != None:
          print("*********************","am returning wunder")
          return
       dt = datetime.datetime.now() + datetime.timedelta(days=-1)
       year = str(dt.year).zfill(4)
       month = str(dt.month).zfill(2)
       day = str(dt.day).zfill(2)
       
       url = 'https://api.weather.com/v2/pws/observations/hourly/7day?stationId='+self.pws+'&format=json&units=e&apiKey='+self.access_key 
       print("-------------url----------------",url)
       response = self.__load_web_page__(url)
       
       if response[0] == True:
          
          return True, self.__parse_data__(response[1])
       else:
          return False, None


   def __load_web_page__( self, url ):
       
       f = urlopen(url)
       json_data = f.read()
       f.close()
       data = json.loads(json_data.decode())
       return True, data
       #'stationID', 'winddirAvg', 'qcStatus', 'obsTimeUtc', 'obsTimeLocal', 'solarRadiationHigh', 'epoch', 'tz', 'uvHigh', 'imperial', 'lat', 'humidityLow', 'humidityAvg', 'lon', 'humidityHigh'])

   def __parse_data__(self,observations):
       data = observations["observations"]
       #print("data",data)
      
       valid_data = []
       self.determing_observation_window()
       for i in data:
          
          if self.match_time(i) == True:
            valid_data.append(i)

           
       
       
       
       results_normal = []
       results_max  = []
        
       for i in valid_data:
           #print(i['imperial'].keys())
           delta_timestamp = 1./24.
      
           
           results_normal.append({
                            "delta_timestamp": 1./24,
                            "TC":self.__convert_to_C__(i['imperial']['tempAvg']),
                            "HUM":float(i['humidityAvg']),
                            "wind_speed":float(i['imperial'][ 'windspeedAvg'])*0.44704,
                            "SolarRadiationWatts/m^2":float(i['solarRadiationHigh']) })  #i["wgusti"] )
           results_max.append({
                            "delta_timestamp": 1./24,
                            "TC":self.__convert_to_C__(float(i['imperial']['tempAvg'])+2),
                            "HUM":float(i['humidityAvg'])*.95,
                            "wind_speed":float(i['imperial']['windspeedAvg'])*1.1*0.44704,
                            "SolarRadiationWatts/m^2":float(i['solarRadiationHigh'])*1.15 })  #i["wgusti"] )
                         

       
       
       self.eto_sources.hset("wunder:"+self.pws+":Normal", { "eto":self.calculate_eto.__calculate_eto__(results_normal,self.alt,self.lat),
                                                            "priority":self.priority,"status":"OK" ,"time":str(datetime.datetime.now())}       ) 
                                                            ### These sources are for information only                                                          
       
       self.eto_sources.hset("wunder:"+self.pws+":Max",   { "eto":self.calculate_eto.__calculate_eto__(results_max,self.alt,self.lat),
                                                            "priority":100,"status":"OK" ,"time":str(datetime.datetime.now())}       )                                                       
                                                            
       self.rain_sources.hset("wunder:"+self.pws ,{"rain":valid_data[-1]['imperial']['precipTotal'],"priority":self.priority,"time":str(datetime.datetime.now())})


   def __convert_to_C__(self, deg_f):
        deg_f = float(deg_f)
        return ((deg_f - 32) * 5.0) / 9.0     
    

   def determing_observation_window(self):
       current_time = time.time()
       reference_time = current_time-24*3600
       ref_datetime = datetime.datetime.fromtimestamp(reference_time)
       self.ref_month = ref_datetime.month
       self.ref_day  = ref_datetime.day
       #print(self.ref_month,self.ref_day)
 
   def match_time(self,input):
       
       temp = input["obsTimeLocal"].split(" ")
       date_list = temp[0].split("-")
       
       if (int(date_list[1]) == self.ref_month) and (int(date_list[2]) == self.ref_day):
         return True
       else:
         return False


class Messo_ETO(object):
    def __init__(self, access_data,eto_dict, rain_dict):
        self.calculate_eto = Calculate_ETO()
        self.eto_dict = eto_dict
        self.messo_data = access_data
        self.alt = access_data["altitude"]
        self.lat = access_data["latitude"]
        self.priority = access_data["priority"]
        self.app_key = self.messo_data["access_key"]
        self.url = self.messo_data["url"]
        self.station = self.messo_data["station"]
        self.token = "&token=" + self.app_key

    def compute_previous_day(self):
        if self.eto_dict.hget("messo:"+self.station+":normal_eto" ) != None:
           print("****************** messo eto returning")
           return
        ts = time.time()
        date_1 = datetime.datetime.fromtimestamp(
            ts - 1 * ONE_DAY).strftime('%Y%m%d')
        date_2 = datetime.datetime.fromtimestamp(
            ts - 0 * ONE_DAY).strftime('%Y%m%d')
        start_time = "&start=" + date_1 + "0800"
        end_time = "&end=" + date_2 + "0900"

        url = self.url + "stid=" + self.station + self.token + start_time + end_time + \
            "&vars=relative_humidity,air_temp,solar_radiation,peak_wind_speed,wind_speed&obtimezone=local"

        req = urllib.request.Request(url)
        response = urllib.request.urlopen(req)
        temp = response.read()
        data = json.loads(temp.decode())

        station = data["STATION"]

        # print data.keys()
        # print data["UNITS"]
        station = station[0]
        station_data = station["OBSERVATIONS"]

        keys = station_data.keys()
        # print "keys",keys
        return_value_normal = []
        return_value_gust = []
        
        for i in range(0, 24):
            temp = {}
            temp["delta_timestamp"] = 1./24.
            temp["wind_speed"] = station_data["wind_speed_set_1"][i]
            temp["HUM"] = station_data["relative_humidity_set_1"][i]
            temp["SolarRadiationWatts/m^2"] = station_data["solar_radiation_set_1"][i]
            temp["TC"] = station_data["air_temp_set_1"][i]
            return_value_normal.append(temp)
            temp = {}
            temp["delta_timestamp"] = 1./24.
            temp["wind_speed"] = station_data["peak_wind_speed_set_1"][i]
            temp["HUM"] = station_data["relative_humidity_set_1"][i]
            temp["SolarRadiationWatts/m^2"] = station_data["solar_radiation_set_1"][i]
            temp["TC"] = station_data["air_temp_set_1"][i]
            return_value_gust.append(temp)
            
        
        print("messo calculation")
        date_string = str(datetime.datetime.now())
        self.eto_dict.hset("messo:"+self.station+":normal_eto",
                           { "eto":self.calculate_eto.__calculate_eto__( results  =  return_value_normal, alt = self.alt,lat = self.lat ), 
                           "priority":self.priority,"status":"OK","time": date_string  })
        self.eto_dict.hset("messo:"+self.station+":gust_eto",
                           { "eto":self.calculate_eto.__calculate_eto__( return_value_gust, self.alt,self.lat ), "priority":100,"status":"OK","time": date_string  })

"""