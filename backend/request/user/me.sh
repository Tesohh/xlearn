curl -sS "localhost:8080/user/me"\
	-X GET\
	-b .cookies -c .cookies | jq
