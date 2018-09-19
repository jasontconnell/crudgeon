IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.DeleteBusiness') AND type in (N'P'))
begin
    drop procedure dbo.DeleteBusiness
end
GO

create procedure dbo.DeleteBusiness

    @Name varchar(150),
    @Value decimal(18,7),
    @Revenue decimal(18,7),
    @Expenses decimal(18,7)

as
begin

    declare @id int
    select @id = ID from Business where 
        Name = @Name 

    delete from Business where ID = @id

end