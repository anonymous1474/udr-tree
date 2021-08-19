implementation of martin kleppmann's highly available tree with move operation

```To run : go run main.go -n <number of replicas> -p <replica ID> -r <number of operations> -o <intervals of operation generation> -s <size of tree>```

Example: 200 total operations, operations generated every 5ms interval(5000 microsecond), tree size 50

1) go run main.go -n 3 - p 0 -r 200 -o 5000 -s 50
2) go run main.go -n 3 - p 1 -r 200 -o 5000 -s 50
3) go run main.go -n 3 - p 2 -r 200 -o 5000 -s 50

after cloning : do ```go get -d``` to get all dependancies

in setup.go change accordingly to setup cloud vms (lines 65-67) as mentioned in comments. Default is for local system run. Comment out (lines 72-74) to stop local run.
