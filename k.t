~1 0 2 /010b
~```a`b /1100b
~" ab" /000b
1. /1.
.1 /0.1
1.1 /1.1
1.+2. /3.
1 2+3 /4 5
1+2 3 /3 4
1 2+3 4 /4 6
-(1 2 3) /-1 -2 -3
- 0101b /0 -1 0 -1
1a /1a
."1+2" /3
`2 /+
`1+(+) /-
.(1;2;`66) /3
1 /1
1b /1b
101b /101b
`x /`x
``x /``x
`"a" /`a
"" /""
"3" /"3"
1+2 /3
1-2 /-1
1 2 3+4 /5 6 7
1 2 3+4a /5a 6a 7a
1-2 3 /-1 -2
"aBc"<"abc" /010b
1a=1 2 /10b
-|1+2 /-3
10 20%2. /5. 10.
101b&110b /100b
101b|110b /111b
3#!5 /0 1 2
-3#!5 /2 3 4
6#!3 /0 1 2 0 0 0
-6#!3 /0 0 0 0 1 2
"ab"#"abc" /"ab"
3_!5 /3 4
-3_!5 /0 1
6_!3 /!0
-6_!3 /!0
1_("ab";"cd") /,"cd"
"ab"_"abc" /,"c"
2 5^"alphabeta" /("pha";"beta")
3^!8 /(0 1;2 3;4 5)
"ABC"^"abcCdeAgh" /("abc";"de";"gh")
"b"\:"abc" /(,"a";,"c")
"x"\:"abxdexfg" /("ab";"de";"fg")
"xd"\:"abxdexfg" /("ab";"exfg")
"xy"\:"abcdx" /,"abcdx"
"x"/:("ab";,"c";"ef") /"abxcxef"
"xy"/:("ab";,"c";"ef") /"abxycxyef"
1 2 2 1?2 /1
0001001b?1b /3
1 2 2 1?1 2 3 /0 1 4
("abc";"de")?"de" /1
("abc";"de")?("de";"gh") /1 2
3 in 0 1 2 /0b
3 in !5 /1b
"a"in"abc" /1b
"ad"in"abc" /10b
in 000b /0b
in 010b /1b
"ab" find "aaabcabca" /2 5
"aa" find "aaaaaaa" /0 2 4
"ab" find "xab" /,1
3=3 /1b
$1b /"1b"
3 4 5=4 /010b
1 2 3 /1 2 3
+/1 2 3 /6
3+/1 2 3 /9
-'!5 /0 -1 -2 -3 -4
-':3 2 4 0 /0 -1 2 -4
2-':3 2 4 0 /1 -1 2 -4
=':3 2 2 3 /1010b
2~':3 2 2 3 /0010b
1(*-)/:(1;2 3) /0 -1
-/:[1;2 3] /-1 -2
(3+2) /5
(1+2)*3 /9
() /()
(1+2;2) /(3;2)
(;1) /(;1)
(1;;) /(1;;)
(2 1)0 /2
1 2 3[1] /2
1 2 3[2 1][0] /3
+- /+-
(*-)1 2 /-1
1+ /1+
(1+)2 /3
1+- /1+-
1+/- /1+/-
-1+/- /-1+/-
x:1 /
x:1;x /1
x*x:2 /4
x::1;x /1
x+:x:1;x /2
(1+x;x:3) /(4;3)
x[1]+:*x:1 2;x /1 3
(-+)[3;5] /-8
1;2 /2
1;;;2 /2
;;2 /2
{x+y} /{x+y}
.{x+y} /((`y;.;`x;.;+);`x`y;"{x+y}";2)
{x+y}[3;4] /7
{z*x+y}[3;4;5] /35
$23 /"23"
$(+':) /"+':"
`i$"31" /31
`F$"1 2 3" /1. 2. 3.
`k@23 /"23"
1>2 /0b
1<2. /1b
1<2a /1b
`x=`y`x`z /010b
"alphabetagamma"="m" /00000000000110b
+/"alpha"="a" /2
+/"abc"="rbx" /1
1101001b~1101001b /1b
*|!100000 /99999
$[1;2;3+4] /2
$[1>2;2;3+4] /7
n:2;x:1;while[n>0;n-:1;x+:1] /3
while[1-1;2+2] /