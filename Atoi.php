<?
class Solution {

    /**
     * @param String $str
     * @return Integer
     */
    function myAtoi($str) {
    	$str = preg_replace('/[A-Za-z].*/', '', $str);
        $min = -2**31;
        $max = 0-$min-1;
        if($str > $max)
            return (int)$max;
        if($str < $min)
            return (int)$min;
        return (int)$str;
    }
}
?>