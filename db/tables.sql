-- departments table
CREATE TABLE departments (
    dept_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- employees table
CREATE TABLE employees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    department_id UUID REFERENCES departments(dept_id) ON DELETE CASCADE,
    salary NUMERIC(10,2),
    location VARCHAR(100),
    joining_date DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- assets table
CREATE TABLE assets (
    asset_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_name VARCHAR(100) NOT NULL,
    asset_type VARCHAR(100),
    asset_price NUMERIC(10,2),
    dept_id UUID REFERENCES departments(dept_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- employee_assets table
CREATE TABLE employee_assets (
    emp_asset_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    emp_id UUID REFERENCES employees(id) ON DELETE CASCADE,
    asset_id UUID REFERENCES assets(asset_id) ON DELETE CASCADE
);

-- salary_category table
CREATE TABLE salary_category (
    cat_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cat_name VARCHAR(100),
    min_sal NUMERIC(10,2),
    max_sal NUMERIC(10,2)
);