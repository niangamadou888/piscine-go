curl -s https://learn.zone01dakar.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"amaniang\"}}){id}}"}' | cut -d ':' -f4 | cut -d '}' -f1
