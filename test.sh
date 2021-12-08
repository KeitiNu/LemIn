echo --------------------------------------
echo example00 6 turns
go run . example00.txt
echo --------------------------------------
echo example01 8 turns
go run . example01.txt
echo --------------------------------------
echo example02 1 turn
go run . example02.txt
echo --------------------------------------
echo example03 6 turns
go run . example03.txt
echo --------------------------------------
echo example04 6 turns
go run . example04.txt
echo --------------------------------------
echo example05 8 turns
go run . example05.txt
echo --------------------------------------
echo badexample00
go run . badexample00.txt
echo --------------------------------------
echo badexample01
go run . badexample01.txt
echo --------------------------------------
echo example06 less than 1.5 min
time go run . example06.txt
echo --------------------------------------
echo example07 less than 2.5 min
time go run . example07.txt
echo --------------------------------------
echo randomtest
go run . test.txt
echo --------------------------------------

