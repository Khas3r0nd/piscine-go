#! /bin/sh 

curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Khas3r0nd\"}}){id}}"}' | jq '.data.user[].id' 

