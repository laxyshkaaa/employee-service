CREATE TABLE passports (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    number VARCHAR(50) NOT NULL
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    phone  VARCHAR(50) UNIQUE NOT NULL,
    company_id INT not null
);
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    company_id INT not null,
    passport_id INT REFERENCES passports(id),
    department_id INT REFERENCES departments(id)

);

create index idx_employees_company_id ON employees(company_id)
create index idx_employees_passport_id ON employees(passport_id)
create index idx_employees_department_id ON employees(department_id)

INSERT INTO passports (type, number) VALUES
('1234', '112233'),
('2345', '223344'),
('3456', '334455'),
('4567', '445566'),
('5678', '556677');

INSERT INTO departments (name, phone, company_id) VALUES
                                                      ('HR', '123-456-7890', 1),
                                                      ('Engineering', '234-567-8901', 2),
                                                      ('Marketing', '345-678-9012', 1),
                                                      ('Sales', '456-789-0123', 3),
                                                      ('IT', '567-890-1234', 1);

INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) VALUES
                                                                                         ('John', 'Doe', '111-111-1111', 1, 1, 1),
                                                                                         ('Jane', 'Smith', '222-222-2222', 2, 2, 2),
                                                                                         ('Alice', 'Johnson', '333-333-3333', 1, 3, 3),
                                                                                         ('Bob', 'Williams', '444-444-4444', 3, 4, 4),
                                                                                         ('Charlie', 'Brown', '555-555-5555', 2, 5, 2),
                                                                                         ('Emily', 'Davis', '666-666-6666', 1, 1, 1),
                                                                                         ('Frank', 'Miller', '777-777-7777', 3, 2, 3),
                                                                                         ('Grace', 'Wilson', '888-888-8888', 1, 3, 2),
                                                                                         ('Hank', 'Moore', '999-999-9999', 2, 4, 2),
                                                                                         ('Ivy', 'Taylor', '000-000-0000', 3, 5, 1);







