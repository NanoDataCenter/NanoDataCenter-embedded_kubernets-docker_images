package eto_support

import (
    
    "fmt"
    "math"
)

type ETO_INPUT struct {
    wind_speed                     float64
    temp_C                         float64
    humidity                       float64
    SolarRadiationWatts_m_squared  float64
    delta_timestamp                float64
}


type ETO_RESULT struct{
    
   input []ETO_INPUT 
    
}

func Construt_Eto_Results(input []ETO_INPUT)ETO_RESULT{
    
    var return_value ETO_RESULT
    return_value.input = input
    return return_value
}    

func (r ETO_RESULT) Calculate_eto(alt,lat float64)float64{
        /*
        Based upon the reference http://www.cimis.water.ca.gov/Content/PDF/CIMIS%20Equation.pdf
        Actual solar radiation used
        albedo of .18 used
        no cloud cover assumed
        */
        
        alt = alt * 0.3048  // convert to meters
        pressure := 101.3-(.0115*alt)+(5.44*math.Pow(10,-7)*alt*alt) // default eq for pressure
        ETod := float64(0.0)
        
        //day_of_year := time.Now().YearDay()
        
        //dr := 1 + .033 * math.Cos(2 * 3.14159 / 365 * float64(day_of_year))
        //delta := .409 * math.Sin(2 * 3.14159 / 365 * float64(day_of_year) - 1.39)
        lat = 3.14159 / 180 * lat

        for _,i :=  range r.input{

            P := pressure
            U2 := i.wind_speed
           

            tc := i.temp_C
            tk := tc + 273.3   //Step 1
            es := .6108 * math.Exp(17.27 * tc / (tk)) // Step 2 vapor pressure Tetens equation
            ea := es * i.humidity / 100.
            VPD := es - ea // vapor pressure deficient // Step 3

            DEL := (4099.*es)/((tc+237.3)*(tc+237.3)) //Step # 4
            G := 0.000646 * P * (1 + 0.000949 * tc)  // Step 6

            W := DEL/(DEL+G)
            
            // For Daytime Conditions (SR>=0.21 MJ/m*m/hr):

            SR := i.SolarRadiationWatts_m_squared
            //
            // Wind Function Step 8
            //
            FU2 := 0.0
            if SR > 10 { 
                FU2 = 0.03 + 0.0576 * U2

            }else{   // For Nighttime Conditions (SR<0.21 MJ/m*m/hr):

                FU2 = 0.125 + 0.0439 * U2
            }
            //SETP 9
            SR = .82 * SR   //# .18 for decidous trees

            NR := SR / (694.5 * (1 - 0.000946 * tc)) // Step 10
            //
            // tk is radiation from ground to sky
            // 273 is temperature of sky
            // ignoring cloulds
            RL := -5.67*math.Pow(10,-8)*math.Pow(273,4) + 5.67*math.Pow(10,-8)*math.Pow(tk,4) 
            
            RL = RL / (694.5 * (1 - 0.000946 * tc))
            ETRL := W * RL

            ETH := NR*W +(1.-W)*VPD*FU2 -ETRL

            ETH = ETH*24.  // equations are per hour we are using %day so multiply by 24 to normalize to hour

            
            
            ETod = ETod + (ETH*i.delta_timestamp)
        }
        fmt.Println(ETod/25.4)
        return ETod / 25.4  // convert  ETO to inches
}











