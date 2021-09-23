package msg_pack_utils

import "github.com/vmihailenco/msgpack/v5"


func Pack_string( input string) string{
    
    b, err := msgpack.Marshal(&input)
    if err != nil {
        panic(err)
    }
    return string(b)
}


func Unpack_string( input string) (string,bool){
    
    item  := ""
    state := true
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        state = false
    } 
    
    return item,state
    
}



func Pack_bool( input bool) string{
    
    b, err := msgpack.Marshal(&input)
    if err != nil {
        panic(err)
    }
    return string(b)
}


func Unpack_bool( input string) (bool,bool){
    
    item  := false
    state := true
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        state = false
    } 
    
    return item,state
    
}

func Pack_int64( input int64) string{
    
    b, err := msgpack.Marshal(&input)
    if err != nil {
        panic(err)
    }
    return string(b)
}


func Unpack_int64( input string) (int64,bool){
    
    item  := int64(0)
    state := true
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        item = 0
        state = false
    } 
    
    return item,state
    
}

func Unpack_interface( input string )( interface{}, bool){
    var item interface{}
    state := true
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        item = 0
        state = false
    } 
    
    return item,state
    
}    
    
    
