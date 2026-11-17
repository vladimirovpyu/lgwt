<?php

function scan(string $input): string
{
    echo 'input='.$input . "\n";
    $k = 0;
    $str = '';
    for ($i = 0; $i < strlen($input); $i++) {
        if (is_numeric($input[$i])) {
            $k = $k * 10 + (int) $input[$i];
            echo 'numeric k=' . $k . "\n";
        }
        elseif ($input[$i] == '[') {
            $substr = '';
            for ($j = $i + 1; $j < strlen($input); $j++) {
                if ($input[$j] == ']') {
                    for ($kk = 0; $kk < $k; $kk++) {
                        $str .= $substr;
                    }
                    break;
                } elseif (is_numeric($input[$j])) {
                    $recStr = $input[$j];
                    $count = 0;
                    for ($r = $j + 1; $r < strlen($input); $r++) {
                        $recStr .= $input[$r];
                        if ($input[$r] == '[') {
                            $count ++;
                        } elseif($input[$r] == ']') {
                            $count --;
                            if ($count == 0) {
                                $substr .= scan($recStr);
                                break;
                            }
                        }
                    }
                    $j = $r;
                } else {
                    $substr .= $input[$j];
                }
            }
            $k = 0;
            $i = $j;
        } else {
            $str .= $input[$i];
        }
    }

    return $str;
}

echo scan("3[a]2[bc]") . "\n";

echo scan("3[a2[c]]") . "\n";

echo scan("2[abc]3[cd]ef") . "\n";

echo scan("nnn2[a3[b]с]nnn") . "\n";

