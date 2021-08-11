IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.GetEmployee') AND type in (N'P'))
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