namespace go api
struct AddRequest{
    1:i64 ID
    2:i64 RoomID
    3:i64 ProductID
    4:i64 Price
    5:string time
    6:i64 UserID
}
struct AddResponse{
    1:string message
}
struct DeleteRequest{
    1:i64 ID
}
struct DeleteResponse{
    1:string message
}
struct SelectRequest{
    1:i64 RoomID
}
struct SelectResponse{
    1:i64 Price
    2:string message
}
service gotest{
    AddResponse add(1:AddRequest req)
    DeleteResponse delete(1:DeleteRequest req)
    SelectResponse select(1:SelectRequest req)
}