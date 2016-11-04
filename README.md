# guid
A distribute ID gnerator, it comes from snowflake but make some change.

guid comes from snowflake but made some changes<br>
1、39bits timestamp, pump id for about 10yeas; 12bits sequences, generate 4096 sequence per Millisecond;12bits workerid, can deploy 4096 servers; highest 1bit reserve;<br>
2、workerid in high position in order to make id global increment;<br>
3、in order to make id be more hashed, random sequence from [0,10) where millisecond change,not set sequence to 0.
