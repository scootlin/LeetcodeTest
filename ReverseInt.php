<?
class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function reverse($x) {
	    $sttrev = $x<0?0-strrev($x):strrev($x);
        $min = -2**31;
        $max = 0-$min-1;
    	if($sttrev < $max && $sttrev > $min){
        	//echo "Reserve  -> $sttrev".PHP_EOL;
        	return (int)$sttrev;
    	}
    	return 0;
    }
}
?>