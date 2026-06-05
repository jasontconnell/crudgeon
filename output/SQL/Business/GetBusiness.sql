IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.ROUTINES 
WHERE ROUTINE_NAME = N'GetBusiness')
begin
    drop procedure dbo.GetBusiness
end
GO


create procedure dbo.GetBusiness
    @Name varchar(150) = null
as 
begin
        select
            *
        from
            Business
        where 
            (@Name is null OR Name = @Name) 
        
end