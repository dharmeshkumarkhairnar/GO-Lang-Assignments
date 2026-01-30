Problem Statement:

 Your task is to create a worker pool with N workers. Jobs should be sent through a shared channel. Each worker should pick a job, wait for some time to simulate work, and print which worker processed which job. The program should make sure that only N jobs are processed at the same time and should stop after all jobs are finished.