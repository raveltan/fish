all:

data:
	rm -f dataset/fish.csv 
	go run dataset/main.go

