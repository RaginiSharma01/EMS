
CREATE TABLE departments (
    dept_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE employees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,

    department_id UUID,
    salary NUMERIC(10,2),
    location VARCHAR(100),

    joining_date TIMESTAMP,

    profile_image TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,


    CONSTRAINT fk_department
        FOREIGN KEY (department_id)
        REFERENCES departments(dept_id)
        ON DELETE SET NULL
);


CREATE TABLE assets (
    asset_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    asset_name VARCHAR(100) NOT NULL,
    asset_type VARCHAR(50),
    asset_price NUMERIC(10,2),

    dept_id UUID,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_asset_department
        FOREIGN KEY (dept_id)
        REFERENCES departments(dept_id)
        ON DELETE SET NULL
);