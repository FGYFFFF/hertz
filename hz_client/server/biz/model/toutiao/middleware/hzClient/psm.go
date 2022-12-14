// Code generated by thriftgo (0.2.3). DO NOT EDIT.

package hzclient

import (
	"context"
	"fmt"
)

type FormReq struct {
	FormValue string `thrift:"FormValue,1" form:"form1" json:"FormValue"`
	FileValue string `thrift:"FileValue,2" form:"FileValue" json:"FileValue" query:"FileValue"`
}

func NewFormReq() *FormReq {
	return &FormReq{}
}

func (p *FormReq) GetFormValue() (v string) {
	return p.FormValue
}

func (p *FormReq) GetFileValue() (v string) {
	return p.FileValue
}

func (p *FormReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FormReq(%+v)", *p)
}

type QueryReq struct {
	QueryValue string `thrift:"QueryValue,1" json:"QueryValue" query:"query1"`
}

func NewQueryReq() *QueryReq {
	return &QueryReq{}
}

func (p *QueryReq) GetQueryValue() (v string) {
	return p.QueryValue
}

func (p *QueryReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QueryReq(%+v)", *p)
}

type PathReq struct {
	PathVaule string `thrift:"PathVaule,1" json:"PathVaule" path:"path1"`
}

func NewPathReq() *PathReq {
	return &PathReq{}
}

func (p *PathReq) GetPathVaule() (v string) {
	return p.PathVaule
}

func (p *PathReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PathReq(%+v)", *p)
}

type BodyReq struct {
	BodyValue  string `thrift:"BodyValue,1" form:"" json:""`
	QueryValue string `thrift:"QueryValue,2" json:"QueryValue" query:"query2"`
}

func NewBodyReq() *BodyReq {
	return &BodyReq{}
}

func (p *BodyReq) GetBodyValue() (v string) {
	return p.BodyValue
}

func (p *BodyReq) GetQueryValue() (v string) {
	return p.QueryValue
}

func (p *BodyReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BodyReq(%+v)", *p)
}

type Resp struct {
	Resp string `thrift:"Resp,1" form:"Resp" json:"Resp" query:"Resp"`
}

func NewResp() *Resp {
	return &Resp{}
}

func (p *Resp) GetResp() (v string) {
	return p.Resp
}

func (p *Resp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Resp(%+v)", *p)
}

type Hertz121 interface {
	FormMethod(ctx context.Context, request *FormReq) (r *Resp, err error)

	QueryMethod(ctx context.Context, request *QueryReq) (r *Resp, err error)

	PathMethod(ctx context.Context, request *PathReq) (r *Resp, err error)

	BodyMethod(ctx context.Context, request *BodyReq) (r *Resp, err error)
}
