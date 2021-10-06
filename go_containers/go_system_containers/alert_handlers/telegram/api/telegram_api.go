package telegram



import (
    //"fmt"
    //"time"
	"log"
    "strconv"    
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

var telebot_token string
var monitor_bot_api  *tgbotapi.BotAPI
var send_bot_api     *tgbotapi.BotAPI
var contact_map      map[string]string

func Init(token string, valid_contact_list []string){
 
    contact_map =  make(map[string]string)
    for _,contact := range valid_contact_list {
        contact_map[contact] = contact
    }
    telebot_token = token
    send_bot_api = Setup_bot_api( token )
    //Send_message("my message 4444444444444444444444444444")
    
}


func Setup_bot_api( token_id string )*tgbotapi.BotAPI{
   
    
    bot_api , err := tgbotapi.NewBotAPI(telebot_token)
	if err != nil {
		log.Panic(err)
	}

	bot_api.Debug = false

	log.Printf("Authorized on account %s", bot_api.Self.UserName)
    
    return bot_api
    
    
}



func Send_message(sent_message string)(bool){
 
    for contact,_ := range contact_map {
        
       contact_id, err := strconv.ParseInt(contact, 10, 64)
       if err != nil {
           panic("bad contact non numberic")
       }
       
       msg := tgbotapi.NewMessage(contact_id,sent_message )
       msg.ParseMode = tgbotapi.ModeHTML
       send_bot_api.Send(msg)
       
       
    }
    return true
}






/*

func bot_monitor() {
	bot, err := tgbotapi.NewBotAPI("1914553879:AAHEo-hrJeEFL4p2SlJme52r_N9nhQvvWbc")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

    //msg := tgbotapi.NewMessage(1575166855, "/alert_msg")
    //bot.Send(msg)
	//panic("done")
	
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}








func get_token(site_data map[string]interface{} )
    secrets.Init_file_handler(site_data )
    telebot_token = secrets.Get_Secret("telegram","token")
}
    


Send_message(contact_list,sent_message)
*/
