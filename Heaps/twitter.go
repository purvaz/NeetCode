package Heaps

import "container/heap"

type Tweet struct {
	time       int
	tweetId    int
	followeeId int
	index      int
}

type Twitter struct {
	time      int
	tweetMap  map[int][]*Tweet
	followMap map[int]map[int]bool
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TConstructor() Twitter {
	return Twitter{tweetMap: make(map[int][]*Tweet), followMap: make(map[int]map[int]bool)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, exists := this.tweetMap[userId]; !exists {
		this.tweetMap[userId] = make([]*Tweet, 0)
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], &Tweet{time: this.time, tweetId: tweetId})
	this.time -= 1
}

func (this *Twitter) GetNewsFeed(userId int) []int {

	result := make([]int, 0)
	minHeap := TweetHeap{}
	// add the user id to the followMap (user follows themselves)
	if _, exists := this.followMap[userId]; !exists {
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = true

	for followeeId, _ := range this.followMap[userId] {
		if _, exists := this.tweetMap[followeeId]; exists {
			// get the latest tweet
			index := len(this.tweetMap[followeeId]) - 1
			tweet := this.tweetMap[followeeId][index]
			// we need followeeId and index because we need the same followeeId's next tweet
			heap.Push(&minHeap, &Tweet{
				time:       tweet.time,
				tweetId:    tweet.tweetId,
				followeeId: followeeId,
				index:      index - 1,
			})
		}
	}

	for len(minHeap) > 0 && len(result) < 10 {
		tweet := heap.Pop(&minHeap).(*Tweet)
		result = append(result, tweet.tweetId)
		if tweet.index >= 0 {
			// get the next tweet from the same followeeId and add in the heap
			nextTweet := this.tweetMap[tweet.followeeId][tweet.index]
			heap.Push(&minHeap, &Tweet{
				time:       nextTweet.time,
				tweetId:    nextTweet.tweetId,
				followeeId: tweet.followeeId,
				index:      tweet.index - 1,
			})
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := this.followMap[followerId]; !exists {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, exists := this.followMap[followerId]; exists {
		delete(this.followMap[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
