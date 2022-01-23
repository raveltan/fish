all: data

data:
	rm -f dataset/fish.csv 
	go run dataset/main.go

viz-catch:
	live-server dataset/fish_catches.html

viz-cons:
	live-server dataset/fish_consumption.html

viz-reg:
	live-server predict/linear-reg/viz.html

viz-tree:
	live-server predict/dt/viz.html

predict:
	go run predict/main.go
