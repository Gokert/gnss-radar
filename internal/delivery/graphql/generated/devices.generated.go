// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNGroupingType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingType(ctx context.Context, v interface{}) (model.GroupingType, error) {
	var res model.GroupingType
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNGroupingType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingType(ctx context.Context, sel ast.SelectionSet, v model.GroupingType) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNSignalType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalType(ctx context.Context, v interface{}) (model.SignalType, error) {
	var res model.SignalType
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNSignalType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalType(ctx context.Context, sel ast.SelectionSet, v model.SignalType) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOGroupingType2ᚕgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingTypeᚄ(ctx context.Context, v interface{}) ([]model.GroupingType, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]model.GroupingType, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNGroupingType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingType(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOGroupingType2ᚕgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingTypeᚄ(ctx context.Context, sel ast.SelectionSet, v []model.GroupingType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNGroupingType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalOGroupingType2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingType(ctx context.Context, v interface{}) (*model.GroupingType, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(model.GroupingType)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOGroupingType2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGroupingType(ctx context.Context, sel ast.SelectionSet, v *model.GroupingType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

func (ec *executionContext) unmarshalOSignalType2ᚕgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalTypeᚄ(ctx context.Context, v interface{}) ([]model.SignalType, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]model.SignalType, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNSignalType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalType(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOSignalType2ᚕgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalTypeᚄ(ctx context.Context, sel ast.SelectionSet, v []model.SignalType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNSignalType2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalOSignalType2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalType(ctx context.Context, v interface{}) (*model.SignalType, error) {
	if v == nil {
		return nil, nil
	}
	var res = new(model.SignalType)
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOSignalType2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐSignalType(ctx context.Context, sel ast.SelectionSet, v *model.SignalType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return v
}

// endregion ***************************** type.gotpl *****************************
