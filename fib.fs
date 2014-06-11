( starting with seeds of 0 fib = 0, 1 fib = 1 )
: fibrec 
    dup 1 > if 1 - dup 1 - recurse swap recurse +
            endif
;


: incfib swap over + ;
: fib
    dup 1 > if 0 1 rot BEGIN rot rot incfib rot 1 - dup 1 = UNTIL 
            drop swap drop
            endif
;
