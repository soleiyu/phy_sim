defualt:
	go run main.go > res.plt
	gnuplot plot.txt
	convert res.png -negate res.png
	eog res.png
movie:
	perl mkgnplscr.pl
	perl split.pl
