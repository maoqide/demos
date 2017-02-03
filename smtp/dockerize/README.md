```shell
#build
docker build -t maoqide/send-mail .

#run 
docker run --env MAILTO_LIST="mail1@xxx.com|mail2@xxx.com" -e SMTPSERVER=smtp.xxx.com -e MAILHOST=host@xxx.com -e PASSWORD=xxxx maoqide/send-mail
```
