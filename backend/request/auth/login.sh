curl -sS "localhost:8080/user/login"\
	-X POST\
	-b .cookies -c cookies\
	-H 'Content-Type: application/json'\
	-d "{
		\"tag\": \"$1\",
		\"password\": \"$2\"
	}" | jq
