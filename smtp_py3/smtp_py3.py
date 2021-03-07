from smtplib import SMTP_SSL, SMTP_SSL_PORT
import time
# code derived from https://www.devdungeon.com/content/read-and-send-email-python#toc-4
# remember to turn less secure app on for the working account only
class SMTP_py3( object):

    def __init__(self,site_data,task_id):
        self.site_id = site_data["site"]
        self.node_id = site_data["local_node"]
        self.task_id = task_id
        self.home_account = site_data["home_account"]
        self.password = site_data["home_password"]
        self.target_account = site_data["target_account"]
        
        
        message_headers =[ f"site_id:  {self.site_id}\r\n",f"node_id:  {self.node_id}\r\n" ,f"task_id:  {self.task_id}\r\n"   ]  
        self.message_header = " ".join(message_headers)
        
        self.send_mail("Reboot","Rebooting")
        
    def send_mail(self,subject,message):
    
        try:
            smtp_server = SMTP_SSL('smtp.gmail.com', 465)
            smtp_server.ehlo()
            #smtp_server.set_debuglevel(1)  # Show SMTP server interactions
       
            smtp_server.login(self.home_account,self.password)
        
            from_email = self.home_account  # or simply the email address
            to_emails = [self.target_account]
            body = self.message_header + message
            headers = f"From: {from_email}\r\n"
            headers += f"To: {', '.join(to_emails)}\r\n" 
            headers += f"Subject:{subject}\r\n"
            email_message = headers + "\r\n" + body
            smtp_server.sendmail(from_email, to_emails, email_message)
            smtp_server.close()
        except:
           
           while True:
              print("no email server critical error")
              time.sleep(60*60) # one hour