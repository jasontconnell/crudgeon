IF EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES 
WHERE TABLE_NAME = N'Employee')
begin
    drop table dbo.Employee
end
GO

create table dbo.Employee (
    ID int identity(1,1) not null,
    
        EmployeeID int not null,
        FirstName varchar(150) not null,
        LastName varchar(150) not null,
        Salary   decimal(18,7) not null,
        StartDate datetime not null,
    
        INDEX [Idx_Employee] NONCLUSTERED (
            
                EmployeeID ASC 
        ), CONSTRAINT [PK_Employee] PRIMARY KEY CLUSTERED 
    (
        [ID] ASC
    )
    WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

