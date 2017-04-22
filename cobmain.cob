       IDENTIFICATION DIVISION.
       PROGRAM-ID. datatyp.
       ENVIRONMENT DIVISION.
       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 Arg1 BINARY-SHORT SIGNED.
       01 Arg2 USAGE COMP-2.
       01 Arg3 USAGE COMP-1.
       01 Arg4 PIC X(6).
       PROCEDURE DIVISION.
       
       add 300 to Arg1.
       add 100.1 to Arg2.
       add 100.2 to Arg3.
       move "hello" to Arg4

       CALL "zeroptr" USING BY reference arg1
                            BY reference arg2
                            BY reference arg3
                            BY reference arg4
       END-CALL

       display arg1
       display arg2
       display arg3
       display arg4
       EXIT PROGRAM.
