curl -sS "localhost:8080/api/user/login"\
    -X POST\
    -c .cookies\
    -b .cookies\
    -H 'Content-Type: application/json'\
    -d '{"username": "michele", "password": "micheletranquillosaggio"}' | jq
