package electionday
import "fmt"
// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
	var count *int
    count = &initialVotes
    return count
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    c := counter
    if counter == nil {
        return 0
    } else  {
    return *c
    }
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var c ElectionResult 
    c.Name = candidateName
    c.Votes = votes
    return &c
	
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	var s string
    x := result
    s = fmt.Sprintf("%s (%d)", x.Name, x.Votes)
    return s
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for i,_ := range(results) {
        if i == candidate {
            results[i] = results[i] - 1
        }
    }
}
