Version 4
SymbolType CELL
LINE Normal -32 -32 -12 -32
LINE Normal 32 0 20 0
LINE Normal -48 16 -24 16
LINE Normal -48 0 -20 0
LINE Normal -48 -16 -24 -16
LINE Normal -20 24 -12 40
LINE Normal -28 40 -20 24
LINE Normal -28 40 -12 40
LINE Normal 0 -48 0 -28
LINE Normal 0 48 0 28
LINE Normal -12 32 -16 32
LINE Normal -32 32 -24 32
ARC Normal -44 32 20 -32 -12 32 -12 -32
ARC Normal -64 32 -12 -32 -32 32 -32 -32
ARC Normal -72 32 -20 -32 -40 32 -40 -32
TEXT -8 -24 Center 2 +
TEXT -8 24 Center 2 -
WINDOW 0 32 -20 Bottom 2
WINDOW 3 32 28 Top 2
SYMATTR Value TXOR
SYMATTR Prefix X
SYMATTR Description Ternary Balanced TXOR gate
SYMATTR ModelFile /Users/manoelribeiroalmeida/dev/trit/LTspice/TXOR.sub
PIN -48 -16 NONE 8
PINATTR PinName a
PINATTR SpiceOrder 1
PIN -48 0 NONE 8
PINATTR PinName b
PINATTR SpiceOrder 2
PIN 0 -48 NONE 8
PINATTR PinName vcc
PINATTR SpiceOrder 3
PIN 0 48 NONE 8
PINATTR PinName vee
PINATTR SpiceOrder 4
PIN 32 0 NONE 8
PINATTR PinName out
PINATTR SpiceOrder 5
PIN -48 16 NONE 8
PINATTR PinName gnd
PINATTR SpiceOrder 6
