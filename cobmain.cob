       IDENTIFICATION DIVISION.
       PROGRAM-ID. datatyp.
       ENVIRONMENT DIVISION.
       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 Arg1 BINARY-SHORT SIGNED.
       01 Arg2 USAGE COMP-2.
       01 Arg3 USAGE COMP-1.
       01 Arg4 PIC X(6).
       01 Arg5 PIC X(8).
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

       CALL "JsonExample" USING BY reference arg1
                            BY reference arg2
                            BY reference arg3
                            BY reference arg4
       END-CALL

       display arg1
       display arg2
       display arg3
       display arg4
      
      
      
      add 1 to Arg1.
       move "Raj" to Arg4
       move "555-905" to Arg5

       CALL "AppendSlice" USING BY reference arg1
                            BY reference arg4
                            BY reference arg5
       END-CALL

       add 2 to Arg1.
       move "Sam" to Arg4
       move "555-955" to Arg5

       CALL "AppendSlice" USING BY reference arg1
                            BY reference arg4
                            BY reference arg5
       END-CALL


       MOVE ZEROES to Arg1.
       move SPACES to Arg4
       move SPACES to Arg5

       CALL "SelectSliceAtIndex" USING BY VALUE 0

       CALL "getSliceAttribute"  USING BY reference arg1
                            BY reference arg4
                            BY reference arg5
       END-CALL

       CALL "PrintSlice"


       display arg1
       display arg4
       display arg5
      
      
       EXIT PROGRAM.
