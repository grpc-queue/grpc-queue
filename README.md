```
└── queuedata
    └── streams
        ├── test1
        │   ├── partition0
        │   │   ├── 0.log
        │   │   ├── 1.log
        │   │   ├── head.position
        │   │   └── partition.info
        │   ├── partition1
        │   │   ├── head.position
        │   │   └── partition.info
        │   └── stream.info
        └── test2
            ├── partition0
            │   ├── head.position
            │   └── partition.info
            ├── partition1
            │   ├── 0.log
            │   ├── head.position
            │   └── partition.info
            ├── partition2
            │   ├── head.position
            │   └── partition.info
            └── stream.info
```

**stream.info:** File to store stream information that we can easly query later on. currently it hold the partitions count only.

example: `2`

**head.position:** file to store the state of the head position(soon to be consumed) of queue.Currently it stores the consumer_group(not used for anything yet), current logfile/byteoffset we are in. (useful for pop operation because we can know where we left off)

example: `main|0.log|6`


**x.log:** file to store the messages of an partition, the first 4 bytes tells how much bytes the message + \n has 

TODO:
* figure out how gestreams is supposed to work
* figure out how and when to perform logfiles cleanup
* figure out if we will support consummer_group id's



