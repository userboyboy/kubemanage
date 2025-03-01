package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/wonderivan/logger"
)

var NameSpace namespace

type namespace struct{}

func NameSpaceRegister(router *gin.RouterGroup) {
	router.PUT("/create", NameSpace.CreateNameSpace)
	router.DELETE("/del", NameSpace.DeleteNameSpace)
	router.GET("/list", NameSpace.GetNameSpaceList)
	router.GET("/detail", NameSpace.GetNameSpaceDetail)
}

// CreateNameSpace 创建namespace
// ListPage godoc
// @Summary      创建namespace
// @Description  创建namespace
// @Tags         NameSpace
// @ID           /api/k8s/namespace/create
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/namespace/create [put]
func (n *namespace) CreateNameSpace(ctx *gin.Context) {
	params := &kubernetes.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	if err := kube.NameSpace.CreateNameSpace(params.Name); err != nil {
		logger.Error("创建命名空间失败:", err)
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// DeleteNameSpace 删除namespace
// ListPage godoc
// @Summary      删除namespace
// @Description  删除namespace
// @Tags         NameSpace
// @ID           /api/k8s/namespace/del
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/namespace/del [delete]
func (n *namespace) DeleteNameSpace(ctx *gin.Context) {
	params := &kubernetes.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.NameSpace.DeleteNameSpace(params.Name); err != nil {
		logger.Error("删除NameSpace失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// GetNameSpaceList 获取NameSpace列表
// ListPage godoc
// @Summary      获取NameSpace列表
// @Description  获取NameSpace列表
// @Tags         NameSpace
// @ID           /api/k8s/namespace/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.NameSpaceResp}"
// @Router       /api/k8s/namespace/list [get]
func (n *namespace) GetNameSpaceList(ctx *gin.Context) {
	params := &kubernetes.NameSpaceListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.NameSpace.GetNameSpaces(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取Namespace列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetNameSpaceDetail 获取NameSpace详情
// ListPage godoc
// @Summary      获取NameSpace详情
// @Description  获取NameSpace详情
// @Tags         NameSpace
// @ID           /api/k8s/namespace/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":data }"
// @Router       /api/k8s/namespace/detail [get]
func (n *namespace) GetNameSpaceDetail(ctx *gin.Context) {
	params := &kubernetes.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.NameSpace.GetNameSpacesDetail(params.Name)
	if err != nil {
		logger.Error("获取Namespace详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
