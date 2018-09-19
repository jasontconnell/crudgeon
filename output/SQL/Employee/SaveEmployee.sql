IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.SaveEmployee') AND type in (N'P'))
begin
    drop procedure dbo.SaveEmployee
end
GO

create procedure dbo.SaveEmployee

    @FirstName varchar(150),
    @LastName varchar(150),
    @StartDate datetime

as
begin
    declare @id int
    select @id = ID from Employee where 

    if exists (select ID from Employee where ID = @id)
    begin

        update Employee set
                FirstName = @FirstName, 
                LastName = @LastName, 
                StartDate = @StartDate
         where ID = @id
    end
    else
    begin
        insert into Employee (
                FirstName, 
                LastName, 
                StartDate
        )
        values (
                @FirstName, 
                @LastName, 
                @StartDate
        )
    end
end