echo "$1"
echo "$2"
curl -sS "localhost:8080/user/signup"\
	-X POST\
	-b .cookies -c .cookies\
	-H 'Content-Type: application/json'\
	-d "{
		\"tag\": \"$1\",
		\"password\": \"$2\"
	}" | jq
