# gopherchain

This is a very simple usage of a Go list to add hashed blocks and append a hash to each block
that points to the hash of the previous block. You can use it to learn and visualize how a blockchain
links the blocks and how a block can be set up to contain any kind of data we wish.

It's easier to think of a blockchain as blocks stacked on top of each other such as so:



                |-----------------|
                |     index       |
                |     data        |
                |     time        |
                |     SHA256      |
                |-----------------|
                        ^
                       / \
                        |
                        |
                |-----------------|
                |     index       |
                |     data        |
                |     time        |
                |     SHA256      |
                |-----------------|