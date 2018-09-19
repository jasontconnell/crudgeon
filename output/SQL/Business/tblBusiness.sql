IF  EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'dbo.Business') AND type in (N'U'))
begin
    drop table dbo.Business
end
GO

create table dbo.Business (
    ID int identity(1,1) not null,
    
        Name varchar(150) not null,
        Value decimal(18,7) not null,
        Revenue decimal(18,7) not null,
        Expenses decimal(18,7) not null,
    
        INDEX [Idx_Business] NONCLUSTERED (
            
                Name ASC 
        ),
    
    CONSTRAINT [PK_Business] PRIMARY KEY CLUSTERED 
    (
        [ID] ASC
    )
    WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

