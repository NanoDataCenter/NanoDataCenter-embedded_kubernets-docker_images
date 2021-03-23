package smtp

import (
    "fmt"
	"strings"
    //"net"
    //"net/mail"
	//"net/smtp"
    //"crypto/tls"
)

var site_id string
var node_id string
var task_id string
var home_account string
var password string
var target_account string
var message_header string

func Initialization( site_data map[string]interface{} , task_id_input string )   {
     
	site_id = site_data["site"].(string)
    node_id = site_data["local_node"].(string)
    task_id = task_id_input
    home_account = site_data["home_account"].(string)
    password = site_data["home_password"].(string)
    target_account = site_data["target_account"].(string)
        
        
    var message_list = []string{ "site_id: "+ site_id+"\r\n",  "node_id:  " +node_id+"\r\n" , "task_id:  "+task_id+"\r\n"   }
	message_header = strings.Join(message_list,"  ")
    fmt.Println(message_header)
}	



/*



// SSL/TLS Email Example

func main() {

    from := mail.Address{"", "username@example.tld"}
    to   := mail.Address{"", "username@anotherexample.tld"}
    subj := "This is the email subject"
    body := "This is an example body.\n With two lines."

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    servername := "smtp.example.tld:465"

    host, _, _ := net.SplitHostPort(servername)

    auth := smtp.PlainAuth("","username@example.tld", "password", host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }

    // Here is the key, you need to call tls.Dial instead of smtp.Dial
    // for smtp servers running on 465 that require an ssl connection
    // from the very beginning (no starttls)
    conn, err := tls.Dial("tcp", servername, tlsconfig)
    if err != nil {
        log.Panic(err)
    }

    c, err := smtp.NewClient(conn, host)
    if err != nil {
        log.Panic(err)
    }

    // Auth
    if err = c.Auth(auth); err != nil {
        log.Panic(err)
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
    }

    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
    }

    err = w.Close()
    if err != nil {
        log.Panic(err)
    }

    c.Quit()

}*/
