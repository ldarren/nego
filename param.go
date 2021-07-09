package nego

type Param struct {
	Key string
	Value string
}

const MatchedPathKey = "$matchedPath"

type Params []Param

func (me Params) Get(name string) string {
	for _, p := range me {
		if p.Key == name {
			return p.Value
		}
	}
	return ""
}

func (me Params) GetMatchedPath() string {
	return me.Get(MatchedPathKey)
}

func (me Params) SetMatchedPath(path string) Params {
	return append(me, Param{Key: MatchedPathKey, Value: path})
}

type paramsKey struct{}

var ParamsKey = paramsKey{}
