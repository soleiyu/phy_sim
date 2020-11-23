
open(fp, "< res.plt");

$num = 300;
$cnt = 0;
$i = 0;

`rm sres/*.png`;

while ($cnt < $num) {
	$line = <fp>;
	if ($i % 30 == 0) {
		print $cnt;
		open(ofp, "> sres.plt");
		print ofp $line;
		`gnuplot splt.txt`;
		$cstr = sprintf("%05d", $cnt);
		`convert sres.png -negate sres/t$cstr.png`;
		$cnt ++;
	}
	$i ++;
}

`convert -delay 5 -loop 0 sres/t*.png res.gif`;

`eog res.gif`;
