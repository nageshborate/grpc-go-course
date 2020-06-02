CMD=$1
shift 1
docker exec -it grpcgocourse $CMD -- $@
