// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"sqlc-rest-api/requests"
	"sqlc-rest-api/responses"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type ProductResolver interface {
	User(ctx context.Context, obj *responses.Product, input *requests.BindUriID) (*responses.User, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Product_user_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *requests.BindUriID
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalOUriID2ᚖsqlcᚑrestᚑapiᚋrequestsᚐBindUriID(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _DeletedProduct_deleted(ctx context.Context, field graphql.CollectedField, obj *responses.DeletedProduct) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_DeletedProduct_deleted(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Deleted, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_DeletedProduct_deleted(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "DeletedProduct",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _DeletedProduct_product_id(ctx context.Context, field graphql.CollectedField, obj *responses.DeletedProduct) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_DeletedProduct_product_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ProductID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNID2int64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_DeletedProduct_product_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "DeletedProduct",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_start_cursor(ctx context.Context, field graphql.CollectedField, obj *responses.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_start_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.StartCursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_start_cursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_end_cursor(ctx context.Context, field graphql.CollectedField, obj *responses.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_end_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EndCursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_end_cursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_has_next_page(ctx context.Context, field graphql.CollectedField, obj *responses.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_has_next_page(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.HasNextPage, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_has_next_page(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_id(ctx context.Context, field graphql.CollectedField, obj *responses.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNID2int64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_name(ctx context.Context, field graphql.CollectedField, obj *responses.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_price(ctx context.Context, field graphql.CollectedField, obj *responses.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_price(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Price, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNInt2int64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_price(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_user_id(ctx context.Context, field graphql.CollectedField, obj *responses.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_user_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNID2int64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_user_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_created_at(ctx context.Context, field graphql.CollectedField, obj *responses.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_created_at(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_created_at(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_user(ctx context.Context, field graphql.CollectedField, obj *responses.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_user(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Product().User(rctx, obj, fc.Args["input"].(*requests.BindUriID))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*responses.User)
	fc.Result = res
	return ec.marshalNUser2ᚖsqlcᚑrestᚑapiᚋresponsesᚐUser(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_user(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_User_id(ctx, field)
			case "name":
				return ec.fieldContext_User_name(ctx, field)
			case "email":
				return ec.fieldContext_User_email(ctx, field)
			case "created_at":
				return ec.fieldContext_User_created_at(ctx, field)
			case "products":
				return ec.fieldContext_User_products(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type User", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Product_user_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _ProductEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *responses.ProductEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ProductEdge_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Cursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ProductEdge_cursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ProductEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ProductEdge_node(ctx context.Context, field graphql.CollectedField, obj *responses.ProductEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ProductEdge_node(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Node, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*responses.Product)
	fc.Result = res
	return ec.marshalNProduct2ᚖsqlcᚑrestᚑapiᚋresponsesᚐProduct(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ProductEdge_node(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ProductEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Product_id(ctx, field)
			case "name":
				return ec.fieldContext_Product_name(ctx, field)
			case "price":
				return ec.fieldContext_Product_price(ctx, field)
			case "user_id":
				return ec.fieldContext_Product_user_id(ctx, field)
			case "created_at":
				return ec.fieldContext_Product_created_at(ctx, field)
			case "user":
				return ec.fieldContext_Product_user(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Product", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Products_edges(ctx context.Context, field graphql.CollectedField, obj *responses.Products) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Products_edges(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Edges, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*responses.ProductEdge)
	fc.Result = res
	return ec.marshalNProductEdge2ᚕᚖsqlcᚑrestᚑapiᚋresponsesᚐProductEdgeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Products_edges(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Products",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "cursor":
				return ec.fieldContext_ProductEdge_cursor(ctx, field)
			case "node":
				return ec.fieldContext_ProductEdge_node(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type ProductEdge", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Products_page_info(ctx context.Context, field graphql.CollectedField, obj *responses.Products) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Products_page_info(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*responses.PageInfo)
	fc.Result = res
	return ec.marshalNPageInfo2ᚖsqlcᚑrestᚑapiᚋresponsesᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Products_page_info(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Products",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "start_cursor":
				return ec.fieldContext_PageInfo_start_cursor(ctx, field)
			case "end_cursor":
				return ec.fieldContext_PageInfo_end_cursor(ctx, field)
			case "has_next_page":
				return ec.fieldContext_PageInfo_has_next_page(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputNewProduct(ctx context.Context, obj interface{}) (requests.CreateProductRequest, error) {
	var it requests.CreateProductRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"user_id", "name", "price"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "user_id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("user_id"))
			it.UserID, err = ec.unmarshalNID2int64(ctx, v)
			if err != nil {
				return it, err
			}
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			it.Name, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "price":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("price"))
			it.Price, err = ec.unmarshalNInt2int64(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdateProduct(ctx context.Context, obj interface{}) (requests.UpdateProductRequest, error) {
	var it requests.UpdateProductRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "name", "price"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			it.ID, err = ec.unmarshalNID2int64(ctx, v)
			if err != nil {
				return it, err
			}
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			it.Name, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "price":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("price"))
			it.Price, err = ec.unmarshalNInt2int64(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var deletedProductImplementors = []string{"DeletedProduct"}

func (ec *executionContext) _DeletedProduct(ctx context.Context, sel ast.SelectionSet, obj *responses.DeletedProduct) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, deletedProductImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("DeletedProduct")
		case "deleted":

			out.Values[i] = ec._DeletedProduct_deleted(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "product_id":

			out.Values[i] = ec._DeletedProduct_product_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var pageInfoImplementors = []string{"PageInfo"}

func (ec *executionContext) _PageInfo(ctx context.Context, sel ast.SelectionSet, obj *responses.PageInfo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, pageInfoImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PageInfo")
		case "start_cursor":

			out.Values[i] = ec._PageInfo_start_cursor(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "end_cursor":

			out.Values[i] = ec._PageInfo_end_cursor(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "has_next_page":

			out.Values[i] = ec._PageInfo_has_next_page(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var productImplementors = []string{"Product"}

func (ec *executionContext) _Product(ctx context.Context, sel ast.SelectionSet, obj *responses.Product) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, productImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Product")
		case "id":

			out.Values[i] = ec._Product_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "name":

			out.Values[i] = ec._Product_name(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "price":

			out.Values[i] = ec._Product_price(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "user_id":

			out.Values[i] = ec._Product_user_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "created_at":

			out.Values[i] = ec._Product_created_at(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "user":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Product_user(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var productEdgeImplementors = []string{"ProductEdge"}

func (ec *executionContext) _ProductEdge(ctx context.Context, sel ast.SelectionSet, obj *responses.ProductEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, productEdgeImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ProductEdge")
		case "cursor":

			out.Values[i] = ec._ProductEdge_cursor(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "node":

			out.Values[i] = ec._ProductEdge_node(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var productsImplementors = []string{"Products"}

func (ec *executionContext) _Products(ctx context.Context, sel ast.SelectionSet, obj *responses.Products) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, productsImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Products")
		case "edges":

			out.Values[i] = ec._Products_edges(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "page_info":

			out.Values[i] = ec._Products_page_info(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNDeletedProduct2sqlcᚑrestᚑapiᚋresponsesᚐDeletedProduct(ctx context.Context, sel ast.SelectionSet, v responses.DeletedProduct) graphql.Marshaler {
	return ec._DeletedProduct(ctx, sel, &v)
}

func (ec *executionContext) marshalNDeletedProduct2ᚖsqlcᚑrestᚑapiᚋresponsesᚐDeletedProduct(ctx context.Context, sel ast.SelectionSet, v *responses.DeletedProduct) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._DeletedProduct(ctx, sel, v)
}

func (ec *executionContext) unmarshalNNewProduct2sqlcᚑrestᚑapiᚋrequestsᚐCreateProductRequest(ctx context.Context, v interface{}) (requests.CreateProductRequest, error) {
	res, err := ec.unmarshalInputNewProduct(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNPageInfo2ᚖsqlcᚑrestᚑapiᚋresponsesᚐPageInfo(ctx context.Context, sel ast.SelectionSet, v *responses.PageInfo) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PageInfo(ctx, sel, v)
}

func (ec *executionContext) marshalNProduct2sqlcᚑrestᚑapiᚋresponsesᚐProduct(ctx context.Context, sel ast.SelectionSet, v responses.Product) graphql.Marshaler {
	return ec._Product(ctx, sel, &v)
}

func (ec *executionContext) marshalNProduct2ᚖsqlcᚑrestᚑapiᚋresponsesᚐProduct(ctx context.Context, sel ast.SelectionSet, v *responses.Product) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Product(ctx, sel, v)
}

func (ec *executionContext) marshalNProductEdge2ᚕᚖsqlcᚑrestᚑapiᚋresponsesᚐProductEdgeᚄ(ctx context.Context, sel ast.SelectionSet, v []*responses.ProductEdge) graphql.Marshaler {
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
			ret[i] = ec.marshalNProductEdge2ᚖsqlcᚑrestᚑapiᚋresponsesᚐProductEdge(ctx, sel, v[i])
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

func (ec *executionContext) marshalNProductEdge2ᚖsqlcᚑrestᚑapiᚋresponsesᚐProductEdge(ctx context.Context, sel ast.SelectionSet, v *responses.ProductEdge) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._ProductEdge(ctx, sel, v)
}

func (ec *executionContext) marshalNProducts2sqlcᚑrestᚑapiᚋresponsesᚐProducts(ctx context.Context, sel ast.SelectionSet, v responses.Products) graphql.Marshaler {
	return ec._Products(ctx, sel, &v)
}

func (ec *executionContext) marshalNProducts2ᚖsqlcᚑrestᚑapiᚋresponsesᚐProducts(ctx context.Context, sel ast.SelectionSet, v *responses.Products) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Products(ctx, sel, v)
}

func (ec *executionContext) unmarshalNUpdateProduct2sqlcᚑrestᚑapiᚋrequestsᚐUpdateProductRequest(ctx context.Context, v interface{}) (requests.UpdateProductRequest, error) {
	res, err := ec.unmarshalInputUpdateProduct(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
