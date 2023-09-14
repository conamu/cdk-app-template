package enhance

type Output struct {
	Name string `json:"name"`
}

func Enhance(input string) Output {
	return Output{
		Name: input + " is Awesome!",
	}
}
