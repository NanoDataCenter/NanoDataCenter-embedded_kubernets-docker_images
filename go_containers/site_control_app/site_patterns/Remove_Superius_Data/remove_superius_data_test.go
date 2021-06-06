package superius_data
import "testing"

/*
 * Assumes redis db hash been setup
 * Assumes tree data base has been setup
 * 
 */


func Test_superious_data(t *testing.T) {
    handle := Construct_Superius_Data_Removal("192.168.1.66", 6379 )
    handle.Ping_redis_connections()
    handle.Remove_superius_data()
    
    
    handle.Close_handles()
   
}
