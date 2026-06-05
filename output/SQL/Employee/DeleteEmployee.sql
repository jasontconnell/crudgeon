

IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.ROUTINES 
WHERE ROUTINE_NAME = N'DeleteEmployee')
begin
    drop procedure dbo.DeleteEmployee
end
GO

create procedure dbo.DeleteEmployee
    @EmployeeID int
as
begin
        delete from Employee where 
            EmployeeID = @EmployeeID 
end