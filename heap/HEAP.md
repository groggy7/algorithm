**Heap** is basically a complete and full binary tree. 
![a](../images/Pasted%20image%2020240205203907.png)
![a](../images/Pasted%20image%2020240205202559.png)
#### What is full binary tree and complete binary tree?
If array representation of the binary tree doesnt have a blank element from left to the right, it is complete binary tree. Full biinary tree is the tree which has no free space and all nodes have 0 or 2 children.
![](../images/Pasted%20image%2020240205203440.png)
## DELETE OPERATION ON HEAP
Only the element with highest priority can be deleted from a heap, which is always the root node. This is the maximum value in max heap, and minimum value in the min heap. 

1. If we start with the following heap, 50 will be deleted from the heap in first delete operation.![](../images/Pasted%20image%2020240205204309.png)
2. Then the element at the end will be sent to the root as following:  ![](../images/Pasted%20image%2020240205204657.png)
3. After that, starting from the root node, we check nodes and take the highest value node, compare it to the parent node, if its higher than parent node, we change their place: ![](../images/Pasted%20image%2020240205212447.png)
## HEAPIFY
Heapify is the process of creating a heap data structure from a binary tree. It is used to create a Min-Heap or a Max-Heap.
![](../images/Pasted%20image%2020240205220821.png)
Starting from the last node, which is 18, we take the greater one from 25 and 18, so we take 25. Then we compare it to its parent, 15. We see its greater than its parent so we switch them.  ![](../images/Pasted%20image%2020240205221350.png)
Now we go under 20, check if its children is greater than its value or not. As we see its children is greater than its value, we switch 40 with 20.![](../images/Pasted%20image%2020240205221528.png)
Now we reached to top. Its time to put 40 in root. 10 comes to the left node and since 20 is greater than 10 we heapify the binary three as such: ![](../images/Pasted%20image%2020240205221725.png)