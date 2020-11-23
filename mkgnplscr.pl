
open(ofp, "> splt.txt");

$col = 0x000020;
$cold = 0x030720;
$pnum = 99;
$mnum = 2;
$cstr = sprintf("%06x", $col);

print ofp "set terminal png size 1080, 720
set out \"sres.png\"
set title \"Phy Result\"
	
set xlabel 'x'
set ylabel 'y'
set grid 
	
set xrange [-10:10]
set yrange [-5:5]

plot \"sres.plt\" using 1:2 axis x1y1 with linespoints notitle linewidth 2 lc rgb \"#$cstr\",\\";

	for ($i = 0; $i < $pnum - 1; $i ++) {
		$v1 = 3 + 2 * $i;
		$v2 = 4 + 2 * $i;
		$col = ($col + $cold) % 0x1000000;
		$cstr = sprintf("%06x", $col);
		print ofp "
	\"\" using $v1:$v2 axis x1y1 with points notitle linewidth 2 lc rgb \"#$cstr\",\\";
	}	

	for ($i = 0; $i < $mnum; $i ++) {
		$v1 = $pnum * 2 + 1 + $i * 2;
		$v2 = $pnum * 2 + 2 + $i * 2;
		print ofp "
	\"\" using $v1:$v2 axis x1y1 with points notitle linewidth 3 lc rgb \"#00FFFF\",\\";
	}

	print ofp "\n";
;
