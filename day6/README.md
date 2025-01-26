### GO Routines and Channels

Goroutines and channels operate together most of the time, so I wanted to bring both in this example.

### Explanation:
* GO Routines -> These are processes controlled by the Golang engine itself, which simulates a queue of processes without the need to request threads from the OS, in this way, an extremely cheaper method of using resources was developed, where each routine has the cost of 2K of memory and all are controlled by the language engine itself;
* Channels -> Channels are means of exchanging information about routines with other sections, divided into:
	* No buffer -> It is an immediate communication line, where the channel receives the value and is blocked until it is emptied;
	* With buffer -> This line has a fixed size of the number of elements received until it is blocked, after which it needs to be cleared to receive new values;
