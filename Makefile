all:

data:
	rm -f dataset/fish.csv 
	go run dataset/main.go

viz-catch:
	live-server dataset/fish_catches.html

viz-cons:
	live-server dataset/fish_consumption.html

