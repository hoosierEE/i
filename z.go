package main

import . "github.com/ktye/wg/module"

func zk() {
	Data(600, "/          48    56 64   72  80   88   96    104   112 120 128 136  144 152  160 168              216    224    232   240    248   256            312\n``x`y`z`k`l`while`in`find`abs`imag`conj`angle`solve`dot`sin`cos`\".(\"`\".)\"`exp`log`plot`a`b`c`d`e`\".qr\"`\".slv\"`\".dot\"`\".rf\"`\".rf1\"`\".rz\"`f`g`h`i`j`m`n\n`\".x\"`\".y\"`\".z\"`\".k\"`\".l\" /320..352\n\n`\".x\":{,/+\"0123456789abcdef\"@(x%16;16/x:0+x)} /`x@ (hex) 320(index.go Atx)\n`\".y\":{x} /328\n`\".z\":{x} /336\n`\".k\":{(.`\".kxy\")[x;1000000]} /344\n\n`\".l\":{t:@x;pad:{(|/#'x)#'x};lst:.`\".l\";k:(.`\".kxy\")[;1]\n kt:{x:$[`T~@x;T x;pad (\"\";\"-\"),$x];(x,'\"|\"),'T y}\n d:{r:!x;x:.x;$[`T~@x;kt[r;x];,'[,'[pad(k'r);\"|\"];k'x]]}\n T:{(,*x),(,(#*x)#\"-\"),1_x:\" \"/:'+pad'$(!x),'.x}\n dd:(\"\";,\"..\")20<#x\n x:$[x~*x;x;(20&#x)#x]\n $[`L~t;k'x;`D~t;d x;`T~t;T x;,k x],dd} /352\n\n`\".kxy\":{t:@x;n:#x;k:(.`\".kxy\")[;y];m:y\n q:{c,(\"\\\\\"/:(0,i)^@[x;i;(qs!\"tnr\\\"\\\\\")x i:&x in qs:\"\\t\\n\\r\\\"\\\\\"]),c:_34}\n s:{$[|/x in \"\\t\\n\\r\"__!31;\"0x\",`x@x;q x]}\n /s:{q x}\n a:{x:$x;$[`c~t;s x;`s~t;\"`\",x;x]}\n d:{r:\"!\",k@.x;n:#!x;x:k@!x;$[(1~n)|(@.x)in`D`T;\"(\",x,\")\";x],r}\n v:{m*:(.`\".kstm\")t;  dd:(\"\";\"..\")m<#x;x:(m&#x)#x\n  x:$[`L~t;k'x;`C~t;x;$x]\n  x:$[`B~t;(*'x),\"b\";`C~t;s x;`S~t;c,(c:\"`\")/:x;`L~t;$[1~n;*x;\"(\",(\";\"/:x),\")\"];\" \"/:x]\n  ((\"\";\",\")(1~n)),x,dd}\n $[n~0;(.`\".kst0\")@t;`T~t;\"+\",d@+x;`D~t;d x;x~*x;a x;v x]} /344\n\n \n`\".kst0\":`B`C`I`S`F`Z`L!(\"0#0b\";c,c:_34;\"!0\";\"0#`\";\"0#0.\";\"0#0a\";\"()\")\n`\".kstm\":`B`C`I`S`F`Z`L!100 100 30 30 20 10 20\n\n`\".dotmv\":{{+/x*y}\\:[x;y]}\n`\".dot\":{f:.`\".dotmv\";f[x;y]}\n\n`\".slv\":{q:$[`i~@*|x;x;(.`\".qr\")@x];s:.`\".qslv\";$[`L~@y;s/:[q;y];s[q;y]]}\n\n`\".qslv\":{H:x 0;r:x 1;n:x 2;m:x 3;j:0;K:!m\n while[j<n;y[K]-:(+/(conj H[j;K])*y K)*H[j;K];K:1_K;j+:1]\n i:n-1;J:!n;y[i]%:r@i\n while[i;j:i_J;i-:1;y[i]:(y[i]-+/H[j;i]*y@j)%r@i]\n n#y}\n\n`\".qr\":{K:!m:#*x;I:!n:#x;j:0;r:n#0a;turn:$[`Z~@*x;{(-x)angle angle y};{x*1. -1@y>0}]\n while[j<n;I:1_I\n  r[j]:turn[s:0. abs/j_x j;xx:x[j;j]]\n  x[j;j]-:r[j]\n  x[j;K]%:%s*(s+abs xx)\n  x[I;K]-:{+/x*y}/:[(conj x[j;K]);x[I;K]]*\\:x[j;K]\n  K:1_K;j+:1];(x;r;n;m)}\n\nej:{(y j),'x_z i j:&(#z)>i:(z x)?y x} /sym t1 t2\navg:{(+/x)%0.+#x}\nvar:{(+/x*x:(x-avg x))%-1+#x}\nstd:{%var x}\n\n`\".rf\": {.5+(x?0)%4294967295.}\n`\".rf1\":{.5+(1.+x?0)%4294967295.}\n`\".rz\": {(%-2*log(.`\".rf1\")x)angle 360*(.`\".rf\")x}\n")
	zn := int32(2217) // should end before 4k
	x := mk(Ct, zn)
	Memorycopy(int32(x), 600, zn)
	dx(Val(x))
}
