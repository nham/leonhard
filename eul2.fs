: incfib swap over + ;

: even? 2 mod 0 = ;

( starting with seeds 0 1, the fibonacci numbers are
   E O O E O O E O O ...

 becase E + O = O and O + O = E. so we need only add up
 terms 0, 3, 6, 9, ...)
: eul2 
    0 0 1 BEGIN incfib dup even? if rot over + rot rot endif
dup 4000000 > UNTIL 
            drop drop
;
