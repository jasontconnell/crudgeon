IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.GetEmployee') AND type in (N'P'))
begin
    drop procedure dbo.GetEmployee
end
GO

create procedure dbo.GetEmployee

as 
begin
        select
            *
        from
            Employee
        
end