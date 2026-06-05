

IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.ROUTINES 
WHERE ROUTINE_NAME = N'DeleteBusiness')
begin
    drop procedure dbo.DeleteBusiness
end
GO

create procedure dbo.DeleteBusiness
    @Name varchar(150)
as
begin
        delete from Business where 
            Name = @Name 
end