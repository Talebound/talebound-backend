package menus

import (
	"context"
	"github.com/the-medo/talebound-backend/converters"
	"github.com/the-medo/talebound-backend/e"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *ServiceMenus) GetMenuItems(ctx context.Context, req *pb.GetMenuItemsRequest) (*pb.GetMenuItemsResponse, error) {
	violations := validateGetMenuItemsRequest(req)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	menuItemRows, err := server.Store.GetMenuItems(ctx, req.MenuId)
	if err != nil {
		return nil, err
	}

	rsp := &pb.GetMenuItemsResponse{
		MenuItems: make([]*pb.MenuItem, len(menuItemRows)),
	}

	for i, menuItemRow := range menuItemRows {
		rsp.MenuItems[i] = converters.ConvertMenuItem(menuItemRow)
	}

	return rsp, nil
}

func validateGetMenuItemsRequest(req *pb.GetMenuItemsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateMenuId(req.GetMenuId()); err != nil {
		violations = append(violations, e.FieldViolation("menu_id", err))
	}

	return violations
}