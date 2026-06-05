IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.ROUTINES 
WHERE ROUTINE_NAME = N'GetEmployee')
begin
    drop procedure dbo.GetEmployee
end
GO


create procedure dbo.GetEmployee
    @EmployeeID int = null
as 
begin
        select
            *
        from
            Employee
        where 
            (@EmployeeID is null OR EmployeeID = @EmployeeID) 
        
end