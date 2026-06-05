

IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.ROUTINES 
WHERE ROUTINE_NAME = N'SaveEmployee')
begin
    drop procedure dbo.SaveEmployee
end
GO

create procedure dbo.SaveEmployee

    @EmployeeID int,
    @FirstName varchar(150),
    @LastName varchar(150),
    @Salary   decimal(18,7),
    @StartDate datetime

as
begin
    declare @id int
    select @id = ID from Employee where 
        EmployeeID = @EmployeeID 
    
    if exists (select ID from Employee where ID = @id)
    begin
        update Employee set
                EmployeeID = @EmployeeID, 
                FirstName = @FirstName, 
                LastName = @LastName, 
                Salary   = @Salary  , 
                StartDate = @StartDate
         where
            ID = @id
    end
    else
    begin
        insert into Employee (
                EmployeeID, 
                FirstName, 
                LastName, 
                Salary  , 
                StartDate
        )
        values (
                @EmployeeID, 
                @FirstName, 
                @LastName, 
                @Salary  , 
                @StartDate
        )
    end
end