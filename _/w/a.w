add:I:II{x+y}
loop:I:I{x/r:i;r}
jota:I:I{r:0;x/r:r+i;r}
iota:I:I{x/r+:i;r}
stoC:I:II{x::y;C x}
stoI:I:II{x::y;I x}
stoJ:J:IJ{x::y;J x}
stoF:F:IF{x::y;F x}
getC:I:I{C x}
getI:I:I{I x}
getJ:J:I{J x}
getF:F:I{F x}
iff:I:I{(x>3)?x+:1;x}
ret:I:I{(x>3)?(:0-x);x}
cond3:I:III{$[x;y;z]}
cond5:I:IIIII{$[x;y;z;x3;x4]}
whl1:I:I{1/(x+:1;?x>5);x}
whl2:I:I{(x<3)?/x+:1;x}
st:I:I{I?255j&1130366807310592j>>J?8*x}
cal1:I:I{1 add x}
cal2:I:II{x add y}
