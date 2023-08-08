package Anonymous
func Anonymous() int{
    anonymous := func (a,b int) int{
        return a+b
    }
    result := anonymous(4,2)
    return result
}
