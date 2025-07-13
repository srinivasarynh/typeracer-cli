package words

import (
	"math/rand"
	"strings"
	"time"
)

type Generator struct {
	difficulty string
	words      []string
}

func NewGenerator(difficulty string) *Generator {
	return &Generator{
		difficulty: difficulty,
		words:      getWordList(difficulty),
	}
}

func (g *Generator) Generate(count int) []string {
	if count <= 0 {
		count = 10
	}

	rand.Seed(time.Now().UnixNano())
	result := make([]string, count)

	for i := 0; i < count; i++ {
		result[i] = g.words[rand.Intn(len(g.words))]
	}

	return result
}

func getWordList(difficulty string) []string {
	switch strings.ToLower(difficulty) {
	case "easy":
		return []string{
			"the", "and", "for", "are", "but", "not", "you", "all", "can", "had",
			"her", "was", "one", "our", "out", "day", "get", "has", "him", "his",
			"how", "its", "may", "new", "now", "old", "see", "two", "who", "boy",
			"did", "use", "way", "she", "oil", "sit", "set", "run", "eat", "far",
			"sea", "eye", "ask", "try", "own", "say", "too", "any", "end", "why",
		}
	case "hard":
		return []string{
			"algorithm", "architecture", "development", "implementation", "optimization",
			"synchronization", "authentication", "authorization", "configuration", "documentation",
			"infrastructure", "transformation", "initialization", "serialization", "deserialization",
			"polymorphism", "encapsulation", "abstraction", "inheritance", "composition",
			"dependency", "middleware", "framework", "library", "database", "repository",
			"controller", "service", "component", "interface", "protocol", "methodology",
			"paradigm", "refactoring", "debugging", "testing", "deployment", "monitoring",
			"scaling", "performance", "security", "reliability", "availability", "maintainability",
			"extensibility", "compatibility", "interoperability", "transparency", "consistency",
		}
	default: // medium
		return []string{
			"about", "after", "again", "against", "almost", "alone", "along", "already",
			"although", "always", "among", "another", "answer", "around", "because",
			"become", "before", "begin", "being", "believe", "below", "between",
			"build", "business", "change", "choose", "close", "computer", "consider",
			"continue", "create", "decision", "develop", "different", "during",
			"early", "education", "enough", "every", "example", "experience",
			"family", "follow", "friend", "government", "group", "happen", "health",
			"history", "house", "human", "important", "include", "information",
			"interest", "issue", "knowledge", "large", "later", "learn", "level",
			"little", "local", "machine", "manage", "member", "might", "minute",
			"moment", "money", "month", "morning", "mother", "music", "nature",
			"never", "night", "nothing", "number", "often", "order", "other",
			"paper", "parent", "part", "person", "place", "point", "power",
			"present", "problem", "program", "project", "provide", "public",
			"question", "rather", "really", "reason", "remember", "report",
			"result", "right", "school", "science", "second", "several",
			"should", "since", "small", "social", "something", "sometimes",
			"special", "start", "state", "story", "student", "study", "system",
			"table", "think", "third", "though", "three", "through", "today",
			"together", "toward", "under", "until", "water", "where", "which",
			"while", "white", "whole", "without", "woman", "world", "would",
			"write", "young",
		}
	}
}
