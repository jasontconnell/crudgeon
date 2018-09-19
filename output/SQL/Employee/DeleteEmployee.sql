IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.DeleteEmployee') AND type in (N'P'))
begin
    drop procedure dbo.DeleteEmployee
end
GO

create procedure dbo.DeleteEmployee

    @EmployeeID int,
    @FirstName varchar(150),
    @LastName varchar(150),
    @StartDate datetime

as
begin

    declare @id int
    select @id = ID from Employee where 
        EmployeeID = @EmployeeID 

    delete from Employee where ID = @id

end