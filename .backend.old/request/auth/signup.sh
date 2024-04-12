curl -sS "localhost:8080/api/user/signup"\
    -X POST\
    -b .cookies\
    -c .cookies\
    -H 'Content-Type: application/json'\
    -d '{"username": "michele9877", "password": "michelepazzofolle"}' | jq
