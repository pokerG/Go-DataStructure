package Trie

//Trie is tree struct mapTrieost operations are O(log n), and it allows for finding similar elements.
// type Trie interface {
// 	Add(string, interface{}) error
// 	Get(string) (interface{}, bool)
// 	Search(string) []interface{}
// 	Remove(string) error
// 	Updata(string, interface{}) error
// }

type Trie struct {
	children  []branch
	value     interface{}
	validLeaf bool
}

type branch struct {
	letter byte
	branch *Trie
}

type stackNode struct {
	top  int
	leaf *Trie
}

type traversalStack []stackNode

//NewTrie returns a fully interative implementation of a Trie.
func NewTrie() *Trie {
	return &Trie{}
}

func (this *Trie) getOrAddChildBranch(r byte) *Trie {
	for _, child := range this.children {
		if child.letter == r {
			return child.branch
		}
	}
	branch := branch{r, NewTrie()}
	this.children = append(this.children, branch)
	return branch.branch
}

func (this *Trie) getChild(r byte) int {
	for i, child := range this.children {
		if child.letter == r {
			return i
		}
	}
	return -1
}

func (this *Trie) getChildBrach(r byte) *Trie {
	for _, child := range this.children {
		if child.letter == r {
			return child.branch
		}
	}
	return nil
}

func (this *Trie) removeChild(c *Trie) {
	last := len(this.children) - 1
	for i, child := range this.children {
		if child.branch == c {
			if i < last {
				this.children[i], this.children[last] = this.children[last], this.children[i]
				break
			}
		}
	}
	this.children = this.children[:last]
}

func (this *Trie) removeChildIndex(i int) {
	last := len(this.children) - 1
	if i < last {
		this.children[i], this.children[last] = this.children[last], this.children[i]
	}
	this.children = this.children[:last]
}

//Add an element to the Trie,mapped to the given value
func (this *Trie) Add(key string, val interface{}) error {
	if key == "" {
		return emptyKey()
	}
	for i, l := 0, len(key); i < l; i++ {
		this = this.getOrAddChildBranch(key[i])
	}

	if this.validLeaf {
		return duplicateKey(key)
	}
	this.validLeaf = true
	this.value = val

	return nil
}

//Get a value from the Trie
//Uses a comma ok format
func (this *Trie) Get(key string) (interface{}, bool) {
	if key == "" {
		return nil, false
	}

	for i, l := 0, len(key); i < l && this != nil; i++ {
		this = this.getChildBrach(key[i])
	}
	if this != nil && this.validLeaf {
		return this.value, true
	}

	return nil, false
}

//Search the Trie for all keys starting with the key.
//A full listing of the Trie is possible using this.Search("")
func (this *Trie) Search(key string) []interface{} {
	var inlineStack [16]stackNode
	var results []interface{}

	for i, l := 0, len(key); i < l && this != nil; i++ {
		this = this.getChildBrach(key[i])
	}

	if this == nil {
		return results
	}

	tip := 0
	inlineStack[0] = stackNode{-1, this} //the first item on stack is last character
	stack := traversalStack(inlineStack[0:1])

	for tip >= 0 {

		stack[tip].top++ // Move on to the next sibling of the last leaf we processed

		if stack[tip].top >= len(stack[tip].leaf.children) { // Check to see if we're out of children

			if stack[tip].leaf.validLeaf { // We are, so add ourselves if we're a valid leaf
				results = append(results, stack[tip].leaf.value)
			}
			stack = stack[:tip] // This branch is completely done; remove it from the stack
			tip--
			continue
		}
		next := stack[tip].leaf.children[stack[tip].top].branch

		if len(next.children) > 0 {
			stack = append(stack, stackNode{-1, next})
			tip++
		} else {
			results = append(results, next.value)
		}
	}
	return results
}

//Remove the key from the Trie
//The Trie will compact itself if possible
func (this *Trie) Remove(key string) error {
	var inlineStack [16]*Trie
	stack := inlineStack[0:0]

	// Identify the leaf associated with the key, and add every node we traverse to our stack.
	for i, l := 0, len(key); i < l && this != nil; i++ {
		stack = append(stack, this)
		this = this.getChildBrach(key[i])
	}

	if this == nil || !this.validLeaf {
		return notFound(key)
	}

	stack = append(stack, this)
	tip := len(key)

	stack[tip].value = nil
	stack[tip].validLeaf = false

	for tip > 0 {

		if stack[tip].validLeaf || len(stack[tip].children) > 0 { // If this branch isn't a valid leaf and has no children.
			break
		}

		trim := stack[tip-1] // Find our parent.

		if len(trim.children) == 1 { // Remove us from our parent. (quickly if possible)
			trim.children = nil
		} else {
			trim.removeChild(stack[tip])
		}
		tip--
	}
	return nil
}

//Update the value of an existing element in the Trie
func (this *Trie) Update(key string, val interface{}) error {
	if key == "" {
		return emptyKey()
	}

	for i, l := 0, len(key); i < l && this != nil; i++ {
		this = this.getChildBrach(key[i])
	}

	if this == nil || !this.validLeaf {
		return notFound(key)
	}

	this.value = val

	return nil
}

//error style
type (
	emptyKeyError     bool
	keynotFoundError  string
	duplicateKeyError string
)

func emptyKey() error {
	return emptyKeyError(true)
}

func notFound(key string) error {
	return keynotFoundError(key)
}

func duplicateKey(key string) error {
	return duplicateKeyError(key)
}

func (err emptyKeyError) Error() string {
	return "key empty"
}

func (err keynotFoundError) Error() string {
	return "key not in Trie: '" + string(err) + "'"
}

func (err duplicateKeyError) Error() string {
	return "key already in Trie: '" + string(err) + "'"
}
