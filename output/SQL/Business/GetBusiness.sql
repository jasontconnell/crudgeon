IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.GetBusiness') AND type in (N'P'))
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