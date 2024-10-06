# wc-tool
I decided to learn Go by building a mini wc command-line `wc-tool`. Along the way, I picked up a bunch of cool Go skills, like handling files, counting bytes, words, and lines, and parsing command-line flags.

I attempted this as part of the Coding Challenges by John Crickett. [Here](https://codingchallenges.fyi/challenges/challenge-wc) is the link to the challenge.

# Steps to run
1. Navigate to `cmd/wc-tool`
2. Run `./wc-tool -<flag> <filename>` (flag = c, l, w), for counting characters, lines and words. For example, `./wc-tool -c test.txt`
3. Run `./wc-tool <filename>` to fetch all the stats. For example, `./wc-tool test.txt`

I have added some sample text files to run these commands on. But a custom file can also be used by running the below command:

1. For specific stats: `cat <filename> | ./wc-tool -<flag>`
2. For all stats: `cat <filename> | ./wc-tool`