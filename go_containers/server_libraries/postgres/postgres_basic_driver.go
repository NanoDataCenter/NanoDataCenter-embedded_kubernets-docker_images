package pg_drv

import (
    
    "fmt"   
    //"strings"
    //"strconv"
    //"time"
	"context"
	"github.com/jackc/pgx/v4"   
	
)


type Postgres_Basic_Driver struct {
    
     conn          *pgx.Conn  
    
    
}


func (v *Postgres_Basic_Driver) connect( connection_url string ) bool {
    
   v.conn = nil 
   conn, err := pgx.Connect(context.Background(), connection_url )
   if err != nil {
      return false
  }
  v.conn = conn
  return true
    
}




/*
 * low level drivers
 * 
 */
 
func (v  Postgres_Basic_Driver )Exec( script string )bool {
    
    
  _, err := v.conn.Exec(context.Background(),script) 
  if err != nil {
      fmt.Println("err",err,script)
      return false
  }
  return true
    
}

 

 
func (v Postgres_Basic_Driver) Close() {
    
    v.conn.Close(context.Background())
    v.conn = nil
}

func (v Postgres_Basic_Driver)Ping()bool{
    
    if v.conn.Ping(context.Background()) != nil {
        return false
    }
    return true
    
}
