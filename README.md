# sieves
GO implementation of parallel sieves of Eratosthenes.

This implementation uses a master / slave approach instead of using `pipelining`. The `pipelining` implementation is more simpler, but there is no control on the number of workers and there is no notion of distribution. 

On the other hand, the master / slave approach, initializes a number of workers where each of them contains a subset of the dataset. This allows for natural distribution where the problem size is partitioned into smaller chucks. 

Channels are used for communication between the master and the slaves. The master tasks is to send values to each worker (slave) so these values are used to filter the worker partition. Then, each worker sends the smaller value in the partition after filtering. The master receives all these local min values and calculates a global min which in turns is used again to filter in the workers. 

When there is no more work, the master receives each partition and forms the final solution, the calculated prime numbers. 


check the C implementation here https://github.com/anicolaspp/Sieve-of-Eratosthenes
