import smtplib  
from email.mime.text import MIMEText  
mailto_list=["qmao@linkernetworks.com"]
mail_host="smtp.xxx.xxx.com"	#smtp server
mail_user="mail@xxx.com"	#username
mail_pass="xxxxxx"		#password
  
def send_mail(to_list, subject, content):
  
    msg = MIMEText(content,_subtype='plain','utf-8')  
    msg['Subject'] = subject
    msg['From'] = mail_user
    msg['To'] = ";".join(to_list)  
    try:  
        server = smtplib.SMTP()  
        server.connect(mail_host)
        server.login(mail_user,mail_pass)  
        server.sendmail(mail_user, to_list, msg.as_string())  
        server.close()
        return True
    except Exception, e:  
        print str(e)  
        return False

if __name__ == '__main__':
    if send_mail(mailto_list,"hello","hello world"):
        print "sending mail succeeded"
    else:
        print "sending mail failed"
