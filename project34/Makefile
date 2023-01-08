CLIENT_BINARY_NAME=client_group5
MINER_BINARY_NAME=miner_group5

build:
	go build -o ${CLIENT_BINARY_NAME} client.go
	go build -o ${MINER_BINARY_NAME} miner.go

# passphrase-free:
# exec ssh-agent $SHELL
# ssh-add

deploy:
	scp -P 8011 ${CLIENT_BINARY_NAME} osgroup5@122.200.68.26:/home/osgroup5 
	scp -P 8011 ${MINER_BINARY_NAME} osgroup5@122.200.68.26:/home/osgroup5
	ssh osgroup5@122.200.68.26 -p 8011 "chmod +x ${MINER_BINARY_NAME}"
	ssh osgroup5@122.200.68.26 -p 8012 "chmod +x ${MINER_BINARY_NAME}"
	ssh osgroup5@122.200.68.26 -p 8013 "chmod +x ${MINER_BINARY_NAME}"
	ssh osgroup5@122.200.68.26 -p 8014 "chmod +x ${MINER_BINARY_NAME}"
	ssh osgroup5@122.200.68.26 -p 8015 "chmod +x ${MINER_BINARY_NAME}"
	ssh osgroup5@122.200.68.26 -p 8016 "chmod +x ${CLIENT_BINARY_NAME}"

run_miner:
	ssh osgroup5@122.200.68.26 -p 8011 "nohup ./${MINER_BINARY_NAME} -diff 5 -port 6353 -addr 10.1.0.11 > miner1.txt 2>&1 &"
	ssh osgroup5@122.200.68.26 -p 8012 "nohup ./${MINER_BINARY_NAME} -diff 5 -port 6353 -addr 10.1.0.12 > miner2.txt 2>&1 &"
	ssh osgroup5@122.200.68.26 -p 8013 "nohup ./${MINER_BINARY_NAME} -diff 5 -port 6353 -addr 10.1.0.13 > miner3.txt 2>&1 &"
	ssh osgroup5@122.200.68.26 -p 8014 "nohup ./${MINER_BINARY_NAME} -diff 5 -port 6353 -addr 10.1.0.14 > miner4.txt 2>&1 &"
	ssh osgroup5@122.200.68.26 -p 8015 "nohup ./${MINER_BINARY_NAME} -diff 5 -port 6353 -addr 10.1.0.15 > miner5.txt 2>&1 &"

run_client:
	ssh osgroup5@122.200.68.26 -p 8016 "./${CLIENT_BINARY_NAME}"

stop_miner:
    # ssh osgroup5@122.200.68.26 -p 8011 'kill -s KILL -- -1'
	ssh osgroup5@122.200.68.26 -p 8011 "ps -ef | grep ${MINER_BINARY_NAME} | grep -v grep | awk '{print \$$2}' | xargs kill -s SIGKILL 2> /dev/null "
	ssh osgroup5@122.200.68.26 -p 8012 "ps -ef | grep ${MINER_BINARY_NAME} | grep -v grep | awk '{print \$$2}' | xargs kill -s SIGKILL 2> /dev/null"
	ssh osgroup5@122.200.68.26 -p 8013 "ps -ef | grep ${MINER_BINARY_NAME} | grep -v grep | awk '{print \$$2}' | xargs kill -s SIGKILL 2> /dev/null"
	ssh osgroup5@122.200.68.26 -p 8014 "ps -ef | grep ${MINER_BINARY_NAME} | grep -v grep | awk '{print \$$2}' | xargs kill -s SIGKILL 2> /dev/null"
	ssh osgroup5@122.200.68.26 -p 8015 "ps -ef | grep ${MINER_BINARY_NAME} | grep -v grep | awk '{print \$$2}' | xargs kill -s SIGKILL 2> /dev/null"
