curl -sS localhost:8080/api/user/logout\
    -b .cookies\
    -c .cookies | jq
