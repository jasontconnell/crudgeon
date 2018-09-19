IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.SaveBusiness') AND type in (N'P'))
begin
    drop procedure dbo.SaveBusiness
end
GO

create procedure dbo.SaveBusiness

    @Name varchar(150),
    @Value decimal(18,7),
    @Revenue decimal(18,7),
    @Expenses decimal(18,7)

as
begin
    declare @id int
    select @id = ID from Business where 
        Name = @Name 

    if exists (select ID from Business where ID = @id)
    begin

        update Business set
                Name = @Name, 
                Value = @Value, 
                Revenue = @Revenue, 
                Expenses = @Expenses
         where ID = @id
    end
    else
    begin
        insert into Business (
                Name, 
                Value, 
                Revenue, 
                Expenses
        )
        values (
                @Name, 
                @Value, 
                @Revenue, 
                @Expenses
        )
    end
end