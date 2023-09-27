package product

import (
	"PowerX/internal/logic/admin/mediaresource"
	"PowerX/internal/model/product"
	product3 "PowerX/internal/uc/powerx/crm/product"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductCategoryTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductCategoryTreeLogic {
	return &ListProductCategoryTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProductCategoryTreeLogic) ListProductCategoryTree(req *types.ListProductCategoryTreeRequest) (resp *types.ListProductCategoryTreeReply, err error) {
	option := product3.FindProductCategoryOption{
		Names:   req.Names,
		OrderBy: req.OrderBy,
	}

	var pId int64 = 0
	if req.CategoryPId > 0 {
		pId = int64(req.CategoryPId)
	}

	// 获取模型类型的列表
	productCategoryTree := l.svcCtx.PowerX.ProductCategory.ListProductCategoryTree(l.ctx, &option, pId)

	// 转化返回类型的列表
	productCategoryReplyList := TransformProductCategoriesToReplyForMP(productCategoryTree)

	return &types.ListProductCategoryTreeReply{
		ProductCategories: productCategoryReplyList,
	}, nil
}

func TransformProductCategoriesToReplyForMP(productCategoryList []*product.ProductCategory) []*types.ProductCategory {
	var productCategoryReplyList []*types.ProductCategory
	//fmt.Dump(productCategoryList)
	for _, category := range productCategoryList {
		node := &types.ProductCategory{
			Id:          category.Id,
			PId:         category.PId,
			Name:        category.Name,
			Sort:        category.Sort,
			ViceName:    category.ViceName,
			Description: category.Description,
			CreatedAt:   category.CreatedAt.String(),
			CoverImage:  mediaresource.TransformMediaResourceToReply(category.CoverImage),
			ImageAbleInfo: types.ImageAbleInfo{
				Icon:            category.Icon,
				BackgroundColor: category.BackgroundColor,
			},
			Children: nil,
		}
		if len(category.Children) > 0 {
			node.Children = TransformProductCategoriesToReplyForMP(category.Children)

		}

		productCategoryReplyList = append(productCategoryReplyList, node)
	}

	return productCategoryReplyList
}
