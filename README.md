Experiment to see if I can port the pong chip8 source to Golang, but still use
my existing emulator to run it... 

Not sure how yet, maybe create go functions for used chip8 opcodes, use THEM to 
write the code, but they return uint16 opcodes that the emulator can use, just 
like it was reading it from 0x200 onwards.

This project would have to mock the memory at location 0x200 (where chip 8 loads
game code)

The hex code for pong

```
00000000: 6a02 6b0c 6c3f 6d0c a2ea dab6 dcd6 6e00  j.k.l?m.......n.
00000010: 22d4 6603 6802 6060 f015 f007 3000 121a  ".f.h.``....0...
00000020: c717 7708 69ff a2f0 d671 a2ea dab6 dcd6  ..w.i....q......
00000030: 6001 e0a1 7bfe 6004 e0a1 7b02 601f 8b02  `...{.`...{.`...
00000040: dab6 600c e0a1 7dfe 600d e0a1 7d02 601f  ..`...}.`...}.`.
00000050: 8d02 dcd6 a2f0 d671 8684 8794 603f 8602  .......q....`?..
00000060: 611f 8712 4602 1278 463f 1282 471f 69ff  a...F..xF?..G.i.
00000070: 4700 6901 d671 122a 6802 6301 8070 80b5  G.i..q.*h.c..p..
00000080: 128a 68fe 630a 8070 80d5 3f01 12a2 6102  ..h.c..p..?...a.
00000090: 8015 3f01 12ba 8015 3f01 12c8 8015 3f01  ..?.....?.....?.
000000a0: 12c2 6020 f018 22d4 8e34 22d4 663e 3301  ..` .."..4".f>3.
000000b0: 6603 68fe 3301 6802 1216 79ff 49fe 69ff  f.h.3.h...y.I.i.
000000c0: 12c8 7901 4902 6901 6004 f018 7601 4640  ..y.I.i.`...v.F@
000000d0: 76fe 126c a2f2 fe33 f265 f129 6414 6500  v..l...3.e.)d.e.
000000e0: d455 7415 f229 d455 00ee 8080 8080 8080  .Ut..).U........
000000f0: 8000 0000 0000                           ......
```

test_rom.ch8

prints:
```
0123456789
ABCDEF

CUOTOS
```

```
00000000: 6600 6700 a000 d675 7605 a005 d675 7605  f.g....uv....uv.
00000010: a00a d675 7605 a00f d675 7605 a014 d675  ...uv....uv....u
00000020: 7605 a019 d675 7605 a01e d675 7605 a023  v....uv....uv..#
00000030: d675 7605 a028 d675 7605 a02d d675 6600  .uv..(.uv..-.uf.
00000040: 7706 a032 d675 7605 a037 d675 7605 a03c  w..2.uv..7.uv..<
00000050: d675 7605 a041 d675 7605 a03c d675 7605  .uv..A.uv..<.uv.
00000060: a04b d675 770c 6600 a03c d675 7605 6090  .K.uw.f..<.uv.`.
00000070: 6190 6290 6390 64f0 a600 f555 d675 7605  a.b.c.d....U.uv.
00000080: a000 d675 7605 6070 6120 6220 6320 6420  ...uv.`pa b c d 
00000090: a600 f555 d675 7605 a000 d675 7605 a019  ...U.uv....uv...
000000a0: d675  
```

`