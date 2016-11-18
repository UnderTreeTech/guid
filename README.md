## GUID
A Distribute ID gnerator, it comes from snowflake but made some changes.

GUID comes from snowflake but made some changes:  

* 39bits timestamp, pump id for about 10yeas; 12bits sequences, generate 4096 sequence per Millisecond;12bits workerid, can deploy 4096 servers; highest 1bit reserve;  
* workerid in high position in order to make id global increment;  
* in order to make id be more hashed, random sequence from [0,10) where millisecond change,not set sequence to 0.
