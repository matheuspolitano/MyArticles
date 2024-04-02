## Elevate Your Go Skills: The Essential Guide to Leveraging Cache Lines

Golang is renowned for its speed, designed from the ground up to offer blistering performance. But ever wonder why some libraries outpace others in execution the same tasks? They were engineered to harness the full potential of the processor and memory.

Quick note before we begin: This is my inagural article on Mechanical Sympathy in Golang. From now on, . i plain to publish weekly pieces exploring different concepts related to CPU cache, memory, and processor operations. We'll explore deep, not just in theory but also "Getting our hands dirty and showing it works in practice". Consider following me and sharing this post to stay updated on this fascinating journey


## What's cacheline 

By google "Cache line is the smallest portion of data the can be mapped into a CPU cache". Let's explore that setence, first off whats CPU cache? 

its a mechanism to store data most used to avoid the cpu travel to the main memory(RAM) every time that it necessary get the value. CPU cache is divide in three different levels L1, L2 and L3. And the L1 its closest level from CPU and L3 its farthest, meaning that get that in L1 is much fast then L3. 

The speed acceses
- L1: about 1ns
- L2: about 4 times slower than L1
- L3: about 10 times slower than L1
- main memory (or RAM) between 50 abd 100 times slower then L1


![Alt text for the image](https://raw.githubusercontent.com/matheuspolitano/MyArticles/master/docs/Article-4/static/cache1.png)

![Alt text for the image](https://raw.githubusercontent.com/matheuspolitano/MyArticles/master/docs/Article-4/static/cache2.png)















## References

- [](https://www.cybercomputing.co.uk/Languages/Hardware/cacheCPU.html)
- [Intel Xeon Phi Processor High Performance Programming](https://www.sciencedirect.com/book/9780128091944/intel-xeon-phi-processor-high-performance-programming)


