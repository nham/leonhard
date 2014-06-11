: isdiv3 3 mod 0 = ;
: isdiv5 5 mod 0 = ;
: add2 swap over + swap ;

: eul1 0 999 BEGIN 
    dup isdiv3 if add2 
               else dup isdiv5 if add2 endif 
               endif 
    1 - dup 0= 
UNTIL ;

( the above seems very ugly. what i'd like to say is "if its divisible by 3 or 5, add2, otherwise move on" )
