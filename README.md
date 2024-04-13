## KETO 

Notifier for active ethernet connections after a break

### Use Case

1. People who experiences frequent disconnects of ethernet 
2. Students who relies on Internet connections provided by Institutions

### Working 

-> Uses `nmcli` to retrieve the active status of the ethernet interface
-> After a disconnection, If status is healthy again, User will be notified via email
-> Sleep or wait period can be configured depending upon the user need

### Usage

-> Clone the repo and populate an environment file with appropriate values
-> See `smtp.go` to check the required env variables
-> `make run` to compile the program and run it
-> Logs will be stored in`/tmp/Keto.log` 
