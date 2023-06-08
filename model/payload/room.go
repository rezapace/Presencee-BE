package payload

type CreateRoomRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	Seat     string `json:"seat" form:"seat" validate:"required"`
}

type CreateRoomResponse struct {
	RoomID uint `json:"room_id"`
}

type UpdateRoomRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	Seat     string `json:"seat" form:"seat" validate:"required"`
}

type UpdateRoomResponse struct {
	RoomID uint `json:"room_id"`
}

type GetRoomResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Seat     string `json:"seat"`
}

type GetRoomsResponse struct {
	Rooms []GetRoomResponse `json:"rooms"`
}
