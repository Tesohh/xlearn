curl -sS "localhost:8080/api/user/recover"\
    -X POST\
    -b .cookies\
    -c .cookies\
    -H 'Content-Type: application/json'\
    -d '{
        "username": "michele",
        "pin": "1234",
        "new_password": "micheletranquillosaggio"
    }' | jq
