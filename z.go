package main

import . "github.com/ktye/wg/module"

func zk() {
	Data(600, "/               64      72    80    88     96     104    112    120                          \n``x`y`z`k`l`a`b`while`\"rf.\"`\"rz.\"`\"uqs.\"`\"uqf.\"`\"gdt.\"`\"lin.\"`\"odo.\"\n\n`\"x.\":{,/+\"0123456789abcdef\"@(x%16;16/x:0+x)} /`x@ (hex)\n`\"t.\":`45         /`t@ token\n`\"p.\":`46         /`p@ parse\n`\"b.\":(`46)[`b;]  /`b@ reinterpret\n`\"c.\":(`46)[`c;]\n`\"i.\":(`46)[`i;]\n`\"s.\":(`46)[`s;]\n`\"f.\":(`46)[`f;]\n`\"z.\":(`46)[`z;]\n\n`\"l.\":{t:@x;pad:{(|/#'x)#'x};k:`kxy 1\n kt:{x:$[`T~@x;T x;pad (\"\";\"-\"),$x];(x,'\"|\"),'T y}\n d:{r:!x;x:.x;$[`T~@x;kt[r;x];,'[,'[pad(k'r);\"|\"];k'x]]}\n T:{$[`L':@'.x;,k x;(,*x),(,(#*x)#\"-\"),1_x:\" \"/:'+pad'$(!x),'.x]}\n dd:(\"\";,\"..\")20<#x\n x:$[x~*x;x;(20&#x)#x]\n $[`L~t;k'x;`D~t;d x;`T~t;T x;,k x],dd} /352\n\n`\"kxy.\":{t:@y;n:#y;k:`kxy x;m:x\n q:{c,(\"\\\\\"/:(0,i)^@[x;i;(qs!\"tnr\\\"\\\\\")x i:&x':qs:\"\\t\\n\\r\\\"\\\\\"]),c:_34}\n s:{$[|/x':\"\\t\\n\\r\"__!31;\"0x\",`x@x;q x]}\n a:{x:$x;$[`c~t;s x;`s~t;\"`\",x;x]}\n d:{r:\"!\",k@.x;n:#!x;x:k@!x;$[(1~n)|(@.x)':`D`T;\"(\",x,\")\";x],r}\n v:{m*:(.`\".kstm\")t;  dd:(\"\";\"..\")m<#x;x:(m&#x)#x\n  x:$[`L~t;k'x;`C~t;x;$x]\n  x:$[`B~t;(*'x),\"b\";`C~t;s x;`S~t;c,(c:\"`\")/:x;`L~t;$[1~n;*x;\"(\",(\";\"/:x),\")\"];\" \"/:x]\n  ((\"\";\",\")(1~n)),x,dd}\n $[n~0;(.`\".kst0\")@t;`T~t;\"+\",d@+y;`D~t;d y;y~*y;a y;v y]} /344\n \n`\"k.\":`kxy 1000000\n\n`\"uqs.\":{x@&~0b~':x:^x} \n`\"uqf.\":{x@&(!#x)=x?x}\n`\"gdt.\":{[t;g](!#t){x g y x}/|.t}\n\n`\".kst0\":`B`C`I`S`F`Z`L!(\"0#0b\";c,c:_34;\"!0\";\"0#`\";\"0#0.\";\"0#0a\";\"()\")\n`\".kstm\":`B`C`I`S`F`Z`L!100 100 30 30 20 10 20\n\n`\"lin.\":{$[`L~@z;(.`\"lin.\")[x;y]'z;[dx:0.+1_-':x;dy:0.+1_-':y;b:(-2+#x)&0|x'z;(y b)+(dy b)*(z-x b)%dx b]]}\n`\"odo.\":{{x/y}'[x;(!*/x)%/:{(*/x)%\\ x}x]}\n\ndot:{dotmv:{{+/x*y}\\:[x;y]};dotmv[x;y]}\n\nsolve:{qslv:{H:x 0;r:x 1;n:x 2;m:x 3;j:0;K:!m\n while[j<n;y[K]-:(+/(conj H[j;K])*y K)*H[j;K];K:1_K;j+:1]\n i:n-1;J:!n;y[i]%:r@i\n while[i;j:i_J;i-:1;y[i]:(y[i]-+/H[j;i]*y@j)%r@i]\n n#y}\n q:$[`i~@*|x;x;qr x];$[`L~@y;qslv/:[q;y];qslv[q;y]]}\n\nqr:{K:!m:#*x;I:!n:#x;j:0;r:n#0a;turn:$[`Z~@*x;{(-x)@angle y};{x*1. -1@y>0}]\n while[j<n;I:1_I\n  r[j]:turn[s:abs/j_x j;xx:x[j;j]]\n  x[j;j]-:r[j]\n  x[j;K]%:%s*(s+abs xx)\n  x[I;K]-:{+/x*y}/:[(conj x[j;K]);x[I;K]]*\\:x[j;K]\n  K:1_K;j+:1];(x;r;n;m)}\n\nany:`30;abs:`32;sin:`44;cos:`39;find:`31;fill:`38;imag:`33;conj:`34;angle:`35;exp:`42;log:`43\n\nej:{(y j),'x_z i j:&~0N=i:(z x)?y x} /sym t1 t2\navg:{(+/x)%0.+#x}\nvar:{(+/x*x:(x-avg x))%-1+#x}\nstd:{%var x}\n\n`\"rf.\": {.5+(x?0)%4294967295.}\n`\"rf1.\":{.5+(1.+x?0)%4294967295.}\n`\"rz.\": {(%-2*log `rf1 x)@360.*`rf x}\n\n`\".html\":{$[`L~@x;\"<div style='display:flex;flex-direction:column'>\",(,/(.`\".html\")'x),\"</div>\";`S~@x;\"<div style='display:flex;flex-direction:row'>\",(,/(.`\".html\")'x),\"</div>\";~`s~@x;\"\";{{\"<\",y,\" id='\",x,\"'></\",y,\">\"}[x;$[y':`i`c`f`z`s`I`B`C`F`Z`S;\"input\";`b~y;\"input type='checkbox'\";`T~y;\"table\";`l~t;\"button\";\"pre\"]]}[$x;@.x]]}\n\n")
	zn := int32(2734) // should end before 4k
	x := mk(Ct, zn)
	Memorycopy(int32(x), 600, zn)
	dx(Val(x))
}
