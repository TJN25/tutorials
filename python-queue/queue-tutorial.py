import sys
from queue import Queue
import queue

def main() -> None:
    q = Queue(maxsize = 3)
    print(q.qsize()) 
    try: 
        for i in range(5):
            q.put(i, block=True)
    except queue.Full:
        print('Queue is full')
    print("\nFull: ", q.full()) 
    print("\nElements dequeued from the queue")
    print(q.get())
    print(q.get())
    print(q.get())
    print("\nEmpty: ", q.empty())
    q.put(1)
    print("\nEmpty: ", q.empty()) 
    print("Full: ", q.full())

if __name__ == '__main__':
    sys.exit(main())
